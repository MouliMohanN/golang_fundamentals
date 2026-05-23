# Functions & Error Handling

---

## 1. Basic Function

```go
func add(a, b int) int {
    return a + b
}
```

- `func` keyword, then name, params in `()`, return type after params
- Consecutive params of the same type can be grouped: `(a, b int)` instead of `(a int, b int)`

---

## 2. Multiple Return Values

Go functions can return more than one value. The most common pattern is `(value, error)`.

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

result, err := divide(10, 3)
if err != nil {
    fmt.Println("error:", err)
}
```

> Always handle the error — ignoring it with `_` is a code smell unless you're certain it can't happen.

---

## 3. Named Return Values

Names act as pre-declared variables. A bare `return` (naked return) returns them automatically.

```go
func minMax(nums []int) (min, max int) {
    min, max = nums[0], nums[0]
    for _, n := range nums[1:] {
        if n < min { min = n }
        if n > max { max = n }
    }
    return // returns min and max
}
```

> Use named returns when the names add clarity to the signature. Avoid naked returns in long functions — they hurt readability.

---

## 4. Variadic Functions

Accept any number of arguments of the same type using `...`. Inside the function they arrive as a slice.

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum(1, 2, 3)       // pass individual values
sum(nums...)       // spread a slice with ...
sum()              // zero args is valid — nums is an empty slice
```

---

## 5. Functions as First-Class Values

Functions are values in Go — assign them to variables, pass them as arguments, return them from other functions.

```go
// Pass a function as an argument
func apply(a, b int, op func(int, int) int) int {
    return op(a, b)
}
apply(3, 4, add)  // pass the add function

// Assign a function to a variable (anonymous function)
multiply := func(a, b int) int { return a * b }
apply(3, 4, multiply)
```

```go
// Return a function from a function (factory pattern)
func makeMultiplier(factor int) func(int) int {
    return func(n int) int {
        return n * factor
    }
}

double := makeMultiplier(2)
double(5)  // 10
```

---

## 6. Closures

A closure is a function that **captures variables from its surrounding scope**. The captured variable lives as long as the closure does — it is not a copy.

```go
func makeCounter() func() int {
    count := 0
    return func() int {
        count++    // count is shared between all calls to this closure
        return count
    }
}

counter := makeCounter()
counter()  // 1
counter()  // 2
counter()  // 3

other := makeCounter()  // independent — has its own count
other()    // 1
```

> Each call to `makeCounter` creates a fresh `count`. The returned closure holds a reference to that specific `count`, not a copy.

### How closure variables are garbage collected

When a variable is captured by a closure, Go automatically moves it from the **stack to the heap**. The garbage collector tracks it by reference — as long as the closure is reachable, the captured variable stays alive.

```go
counter := makeCounter()  // counter holds a reference → count stays alive on heap
counter()                 // count = 1
counter()                 // count = 2

counter = nil             // no more references to the closure
                          // GC can now collect both the closure and count
```

Each closure from the same factory gets its **own** variable on the heap — they never share:

```go
a := makeCounter()  // a's count lives at its own heap address
b := makeCounter()  // b's count lives at a separate heap address
a()                 // a's count = 1
b()                 // b's count = 1  (completely independent)
```

---

## 7. `defer`

`defer` schedules a function call to run **when the surrounding function returns**, regardless of how it returns (normal, error, panic).

```go
func readFile(name string) {
    fmt.Println("opening")
    defer fmt.Println("closing")  // guaranteed to run
    // ... rest of function
}
```

**Multiple defers run in LIFO order** (last deferred = first to run):

```go
defer fmt.Println("1")  // runs 3rd
defer fmt.Println("2")  // runs 2nd
defer fmt.Println("3")  // runs 1st
```

**Common use cases:**
- Close a file/DB connection after opening
- Unlock a mutex after locking
- Log when a function exits

---

## 8. `panic` and `recover`

`panic` stops normal execution and unwinds the stack, running all deferred functions along the way.

`recover` can catch a panic — but **only inside a deferred function**.

```go
func safeDiv(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered from panic: %v", r)
        }
    }()
    result = a / b  // panics if b == 0
    return
}

safeDiv(10, 2)  // result=5, err=nil
safeDiv(10, 0)  // result=0, err="recovered from panic: runtime error: integer divide by zero"
```

### Step-by-step execution when `b == 0`

```
1. defer func(){...}()
   └─ registers the anonymous function to run on exit. Does NOT run yet.

2. result = a / b
   └─ b is 0 → integer divide by zero → PANIC triggered
   └─ normal execution stops immediately here
   └─ the `return` line below is never reached

3. Go unwinds the stack, running all deferred functions

4. The deferred anonymous function runs:
   └─ r := recover()
         └─ catches the panic value ("runtime error: integer divide by zero")
         └─ stops the panic from crashing the program
         └─ r is now that error string (not nil)

5. r != nil → true

6. err = fmt.Errorf("recovered from panic: %v", r)
   └─ err is a named return value — assigning it here changes what
      the function returns, but does NOT trigger an immediate return
   └─ the deferred function simply finishes

7. safeDiv exits: returns result=0 (zero value, never assigned), err=<the error>
```

### Named return values do NOT cause immediate returns

Setting a named return variable only changes what will be returned when the function eventually exits — it does not trigger a return by itself. The function (or deferred function) must finish normally.

```go
func safeDiv(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered: %v", r)  // sets err, keeps running
            // deferred func finishes here → safeDiv exits with result=0, err=<error>
        }
    }()
    result = a / b  // panic stops execution HERE
    return          // never reached when panic occurs
}
```

> Use `panic` for truly unrecoverable states (programmer errors, impossible conditions). For expected failures, return an `error` instead.

---

## 9. The `error` Interface

`error` is a built-in interface with a single method:

```go
type error interface {
    Error() string
}
```

Any type with an `Error() string` method satisfies it. `nil` means no error.

**`errors.New`** — simplest way, for static messages:
```go
var ErrNotFound = errors.New("not found")
```

**`fmt.Errorf`** — for dynamic, formatted messages:
```go
return fmt.Errorf("findAnimal(%q): animal does not exist", name)
```

> Declare sentinel errors (`var ErrX = errors.New(...)`) at package level so callers can compare against them.

---

## 10. Custom Error Types

Implement the `error` interface on a struct to attach structured data to an error.

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %q: %s", e.Field, e.Message)
}
```

This lets callers extract the `Field` and `Message` programmatically — not just read the string.

### Why `&ValidationError{...}` and not `ValidationError{...}`

The `Error()` method is defined on `*ValidationError` (pointer receiver):

```go
func (e *ValidationError) Error() string { ... }
```

Only `*ValidationError` satisfies the `error` interface — `ValidationError` (value) does not, because the method is attached to the pointer. So you must return a pointer:

```go
return &ValidationError{Field: "age", Message: "cannot be negative"}  // ✓ pointer
return ValidationError{Field: "age", Message: "cannot be negative"}   // ✗ compile error
```

---

## 11. `errors.Is` and `errors.As`

### `errors.Is` — check if an error matches a sentinel value

Works even when the error is **wrapped** (see section 12).

```go
if errors.Is(err, ErrNotFound) {
    // handle not found
}
```

### `errors.As` — extract a specific error type from the chain

```go
var valErr *ValidationError
if errors.As(err, &valErr) {
    fmt.Println(valErr.Field)    // structured access
    fmt.Println(valErr.Message)
}
```

### Why `&valErr` in `errors.As`

`errors.As` needs to **write into** `valErr` — it needs to set it to the unwrapped error value. To modify a variable from inside a function in Go, you must pass its address:

```go
var valErr *ValidationError   // valErr is *ValidationError
errors.As(err, &valErr)       // &valErr is **ValidationError — lets errors.As set it
```

- If you passed `valErr` (the pointer itself), `errors.As` would get a copy — any assignment inside would be lost when it returns
- Passing `&valErr` gives `errors.As` the address of the pointer, so it can reach back and set the original variable

> Never use `==` to compare errors — it breaks as soon as errors are wrapped. Always use `errors.Is` / `errors.As`.

---

## 12. Error Wrapping with `%w`

`%w` in `fmt.Errorf` wraps an error, adding context while preserving the original for inspection.

```go
func loadConfig(path string) error {
    _, err := findAnimal(path)
    if err != nil {
        return fmt.Errorf("loadConfig(%q): %w", path, err)  // wrap
    }
    return nil
}
```

The caller sees the full message but can still unwrap it:

```go
err := loadConfig("phoenix")
// err.Error() → `loadConfig("phoenix"): findAnimal: not found`

errors.Is(err, ErrNotFound)  // true — unwraps through the chain
```

**Wrapping chain:**
```
loadConfig error
  └─ fmt.Errorf("loadConfig: %w", ...)
       └─ findAnimal error
            └─ fmt.Errorf("findAnimal: %w", ErrNotFound)
                 └─ ErrNotFound
```

---

## Quick Reference

| Concept | Syntax |
|---|---|
| Basic function | `func name(a, b int) int` |
| Multiple returns | `func f() (int, error)` |
| Named returns | `func f() (x, y int)` + naked `return` |
| Variadic | `func f(nums ...int)` |
| Spread slice | `f(slice...)` |
| First-class function | `op := func(a, b int) int { ... }` |
| Higher-order function | `func apply(op func(int,int) int)` |
| Factory function | `func make() func() int { return func() int {...} }` |
| defer | `defer cleanup()` — runs on function exit, LIFO |
| panic | `panic("message")` — unrecoverable error |
| recover | `r := recover()` — inside a deferred func only |
| Simple error | `errors.New("msg")` |
| Formatted error | `fmt.Errorf("context: %w", err)` |
| Custom error type | struct + `Error() string` method |
| Check sentinel | `errors.Is(err, ErrX)` |
| Extract type | `errors.As(err, &target)` |
