# Go Basics

## Implicit Variable Declaration

Use `:=` to declare and initialize a variable without stating its type. Go infers the type from the value.

```go
name     := "Mouli"  // string
age      := 25       // int
pi       := 3.14     // float64
isActive := true     // bool
```

> Use `var x int = 5` when you need an explicit type, or when declaring without initializing.

---

## Printing the Type of a Variable

Use `fmt.Printf` with the `%T` verb to print a variable's type at runtime.

```go
fmt.Printf("%T\n", age)  // int
fmt.Printf("%T\n", pi)   // float64
```

---

## Printf vs Println

| Feature | `fmt.Println` | `fmt.Printf` |
|---|---|---|
| Newline | Added automatically | Must add `\n` manually |
| Spacing | Adds spaces between args | No automatic spacing |
| Formatting | None | Full format verbs (`%s`, `%d`, `%f`, etc.) |

```go
fmt.Println("Hello,", "World!", 25)         // Hello, World! 25
fmt.Printf("Hello, %s! Age: %d\n", "Mouli", 25) // Hello, Mouli! Age: 25
fmt.Printf("Pi: %.4f\n", 3.14159)          // Pi: 3.1416
```

**Common format verbs:**
- `%v` — default format
- `%T` — type of the value
- `%d` — integer
- `%f` — float (`%.2f` for 2 decimal places)
- `%s` — string
- `%q` — quoted string
- `%t` — boolean

---

## Type Casting

Go requires **explicit** type conversion — there is no implicit casting.

```go
var x int     = 42
var y float64 = float64(x)   // int → float64
var z int     = int(y * 1.5) // float64 → int (truncates, does not round)
```

### String ↔ Byte Slice

```go
str   := "hello"
b     := []byte(str)   // string → []byte  →  [104 101 108 108 111]
back  := string(b)     // []byte → string  →  "hello"
```

### Int → String via Rune

```go
char := string(rune(72)) // → "H"  (Unicode code point 72)
```

> Do **not** use `string(42)` expecting `"42"` — it gives the character at that code point.  
> Use `fmt.Sprintf("%d", 42)` or `strconv.Itoa(42)` to convert a number to its digit string.
