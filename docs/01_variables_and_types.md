# Variables & Types

## Implicit Declaration (`:=`)

Go infers the type from the right-hand value. Only usable inside functions.

```go
name     := "Mouli"   // string
age      := 25        // int
pi       := 3.14159   // float64
isActive := true      // bool
```

Use `var` instead when: you need an explicit type, declaring at package level, or declaring without initializing.

```go
var score int = 100
var label string       // zero value ""
var a, b, c int = 1, 2, 3
```

---

## Zero Values

Every type has a default when declared without a value.

| Type | Zero value |
|---|---|
| `int`, `float64` | `0`, `0.0` |
| `bool` | `false` |
| `string` | `""` |
| pointer, slice, map, channel, func | `nil` |

---

## Multiple Assignment & Swap

```go
x, y := 10, 20
x, y = y, x   // swap — no temp variable needed
```

---

## Blank Identifier

Use `_` to discard values you don't need.

```go
quotient, _ := divide(10, 3)   // remainder discarded
```

---

## Constants & `iota`

`const` values are set at **compile time** and can never be changed. The type is inferred from the value, just like `:=`.

```go
const gravity = 9.81        // inferred float64
const appName = "GoFundamentals"  // inferred string
```

---

### Enums with `iota`

`type Weekday int` creates a **new named type** based on `int`. `Weekday` and `int` are now distinct types — the compiler prevents you from mixing them accidentally. This is Go's enum pattern.

```go
type Weekday int
```

Inside a `const (...)` block, `iota` starts at `0` and increments by `1` for each line. You only write `= iota` once — every line below **implicitly repeats the last expression** with the new `iota` value.

```go
const (
    Sunday    Weekday = iota  // iota=0 → Sunday=0
    Monday                    // iota=1 → Monday=1
    Tuesday                   // iota=2 → Tuesday=2
    Wednesday                 // iota=3 → Wednesday=3
    Thursday                  // iota=4 → Thursday=4
    Friday                    // iota=5 → Friday=5
    Saturday                  // iota=6 → Saturday=6
)
```

---

### Bit flags with `iota`

`<<` is the **left-shift operator** — it multiplies by powers of 2. `1 << 10` means 1 × 2¹⁰ = 1024.

`iota` **resets to `0`** at the start of every new `const (...)` block.

```go
const (
    KB = 1 << (10 * (iota + 1))  // iota=0 → 10*(0+1)=10  → 1<<10 = 1,024
    MB                            // iota=1 → 10*(1+1)=20  → 1<<20 = 1,048,576
    GB                            // iota=2 → 10*(2+1)=30  → 1<<30 = 1,073,741,824
)
```

Why `iota + 1`? Because without it, `iota=0` would give `1 << 0 = 1` (not a useful byte unit). Adding 1 shifts the sequence to start at KB.

---

### Key rules

| Rule | Detail |
|---|---|
| `iota` starts at `0` | Resets in every new `const (...)` block |
| Implicit repetition | Lines without `=` repeat the previous expression with the new `iota` |
| Named types for enums | `type X int` gives type safety — `int` and `X` are not interchangeable |
| Compile-time only | Constants can't be assigned at runtime or hold values from functions |

---

## Printing Types at Runtime

Use `%T` with `fmt.Printf`.

```go
fmt.Printf("%T\n", 42)          // int
fmt.Printf("%T\n", 3.14)        // float64
fmt.Printf("%T\n", []int{1,2})  // []int
```

To print types across a mixed collection, combine `[]any`, `range`, and `%T`:

```go
values := []any{42, 3.14, "hello", true, []int{1, 2}, map[string]int{"a": 1}}

for _, v := range values {
    fmt.Printf("value=%-20v type=%T\n", v, v)
}
```

**Line by line:**

`[]any{...}` — a slice where each element can hold any type. `any` is an alias for `interface{}`, Go's way of saying "type is unknown at compile time." This lets you mix `int`, `float64`, `string`, `bool`, slices, and maps in one collection.

`for _, v := range values` — `range` yields two values per iteration: the index and the element. `_` discards the index. `v` holds the current element.

`fmt.Printf("value=%-20v type=%T\n", v, v)` — `v` is passed twice: `%-20v` prints the value left-aligned in a 20-character field (keeps columns tidy), `%T` prints its runtime type.

**Output:**
```
value=42                   type=int
value=3.14                 type=float64
value=hello                type=string
value=true                 type=bool
value=[1 2]                type=[]int
value=map[a:1]             type=map[string]int
```

> `any` / `interface{}` trades compile-time type safety for flexibility. Use it only when the type genuinely isn't known ahead of time — prefer concrete types everywhere else.

---

## Numeric Type Casting

Go never casts implicitly — always explicit.

```go
var i   int    = 100
var f   float64 = float64(i)    // int → float64
var back int   = int(f * 1.9)   // float64 → int  (truncates, never rounds)
var i32 int32  = int32(i)
var u   uint   = uint(i)
```

> **Truncation vs rounding:** `int(2.9)` gives `2`, not `3`.

---

## String Conversions

### `string` ↔ `[]byte`
Operates on raw UTF-8 bytes. `len([]byte(s))` counts bytes, not characters.

```go
b    := []byte("hello")   // [104 101 108 108 111]
back := string(b)         // "hello"
```

### `string` ↔ `[]rune`
Operates on Unicode code points. Correct for multibyte characters.

```go
r    := []rune("世界")    // [19990 30028]  — 2 runes, not 6 bytes
back := string(r)         // "世界"
```

### `rune` → `string`
Gives the character at that Unicode code point — **not** the digit string.

```go
string(rune(72))  // "H"   (code point 72)
string(rune(42))  // "*"   (not "42"!)
```

To get `"42"` from `42`, use `strconv`.

---

## `strconv` — String ↔ Primitive

```go
// int ↔ string
strconv.Itoa(42)         // "42"
strconv.Atoi("123")      // 123, nil

// float ↔ string
strconv.FormatFloat(3.14159, 'f', 2, 64)  // "3.14"
strconv.ParseFloat("2.718", 64)           // 2.718, nil

// bool ↔ string
strconv.FormatBool(true)      // "true"
strconv.ParseBool("true")     // true, nil
```

> `Atoi` and `ParseFloat` return `(value, error)` — always check the error in production code.
