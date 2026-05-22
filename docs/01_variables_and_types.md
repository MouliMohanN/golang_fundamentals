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

`const` values are set at compile time and cannot change.

```go
const gravity = 9.81
const appName = "GoFundamentals"
```

`iota` auto-increments within a `const` block — ideal for enums and bit flags.

```go
type Weekday int
const (
    Sunday  Weekday = iota  // 0
    Monday                  // 1
    Saturday                // 6
)

const (
    KB = 1 << (10 * (iota + 1))  // 1024
    MB                            // 1048576
    GB                            // 1073741824
)
```

---

## Printing Types at Runtime

Use `%T` with `fmt.Printf`.

```go
fmt.Printf("%T\n", 42)          // int
fmt.Printf("%T\n", 3.14)        // float64
fmt.Printf("%T\n", []int{1,2})  // []int
```

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
