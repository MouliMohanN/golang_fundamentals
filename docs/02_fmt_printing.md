# fmt Package — Printing & Formatting

## Print Family

| Function | Newline | Spaces between args | Format string |
|---|---|---|---|
| `fmt.Print` | No | Only between non-string args | No |
| `fmt.Println` | Yes (auto) | Yes (always) | No |
| `fmt.Printf` | No | No | Yes |

```go
fmt.Print("no newline ")
fmt.Println("auto newline", "spaced")
fmt.Printf("hello %s, age %d\n", "Mouli", 25)
```

---

## General Verbs

| Verb | Meaning |
|---|---|
| `%v` | Default format |
| `%+v` | Struct with field names |
| `%#v` | Go syntax representation |
| `%T` | Type of the value |
| `%%` | Literal `%` |

```go
type Point struct{ X, Y int }
p := Point{3, 7}
fmt.Printf("%v\n", p)   // {3 7}
fmt.Printf("%+v\n", p)  // {X:3 Y:7}
fmt.Printf("%#v\n", p)  // main.Point{X:3, Y:7}
fmt.Printf("%T\n", p)   // main.Point
```

---

## Boolean

```go
fmt.Printf("%t\n", true)   // true
fmt.Printf("%t\n", false)  // false
```

---

## Integer Verbs

| Verb | Meaning | Example (n=255) |
|---|---|---|
| `%d` | Decimal | `255` |
| `%b` | Binary | `11111111` |
| `%o` | Octal | `377` |
| `%O` | Octal with `0o` prefix | `0o377` |
| `%x` | Hex lowercase | `ff` |
| `%X` | Hex uppercase | `FF` |
| `%c` | Unicode character | `ÿ` |
| `%U` | Unicode code point | `U+00FF` |
| `%q` | Quoted character | `'ÿ'` |

```go
fmt.Printf("%b\n", 255)  // 11111111
fmt.Printf("%x\n", 255)  // ff
fmt.Printf("%c\n", 72)   // H
fmt.Printf("%U\n", '世') // U+4E16
```

---

## Float Verbs

| Verb | Meaning | Example (f=123456.789) |
|---|---|---|
| `%f` | Decimal notation | `123456.789000` |
| `%F` | Same, uppercase `INF`/`NAN` | `123456.789000` |
| `%e` | Scientific, lowercase | `1.234568e+05` |
| `%E` | Scientific, uppercase | `1.234568E+05` |
| `%g` | Compact (`%e` for large, `%f` otherwise) | `123456.789` |
| `%G` | Same as `%g`, uppercase `E` | `123456.789` |

```go
fmt.Printf("%e\n", 123456.789)  // 1.234568e+05
fmt.Printf("%.2f\n", 3.14159)   // 3.14
fmt.Printf("%g\n", 0.00001)     // 1e-05
```

---

## String & Byte Verbs

| Verb | Meaning |
|---|---|
| `%s` | Plain string |
| `%q` | Double-quoted, escaped string |
| `%x` | Hex encoding of each byte, lowercase |
| `%X` | Hex encoding of each byte, uppercase |

```go
fmt.Printf("%q\n", "hello\nworld")  // "hello\nworld"
fmt.Printf("%x\n", "hi")            // 6869
```

---

## Pointer

```go
val := 42
fmt.Printf("%p\n", &val)  // 0xc000014090
```

---

## Width, Precision & Padding

| Format | Meaning |
|---|---|
| `%10d` | Right-align in field of width 10 |
| `%-10d` | Left-align in field of width 10 |
| `%010d` | Zero-pad to width 10 |
| `%+d` | Always show sign (`+42`, `-42`) |
| `% d` | Space for positive (aligns with negatives) |
| `%.2f` | 2 decimal places |
| `%10.2f` | Width 10, 2 decimal places |
| `%.3s` | Truncate string to 3 characters |
| `%*d` | Dynamic width (pass as argument) |
| `%.*f` | Dynamic precision (pass as argument) |

```go
fmt.Printf("[%10d]\n", 42)     // [        42]
fmt.Printf("[%-10d]\n", 42)    // [42        ]
fmt.Printf("[%010d]\n", 42)    // [0000000042]
fmt.Printf("[%.2f]\n", 3.14159) // [3.14]
fmt.Printf("[%*d]\n", 8, 42)   // [      42]   — dynamic width
```

---

## Sprintf — Format to String

Returns a formatted string without printing it.

```go
greeting := fmt.Sprintf("Hello, %s!", "Mouli")
hex      := fmt.Sprintf("0x%X", 255)        // "0xFF"
label    := fmt.Sprintf("user_%05d", 42)    // "user_00042"
```

### Sprint / Sprintln

```go
s1 := fmt.Sprint("a", "b", "c")    // "abc"
s2 := fmt.Sprintln("a", "b", "c")  // "a b c\n"
```

---

## Fprintf — Format to any `io.Writer`

Writes to stdout, stderr, files, buffers — anything implementing `io.Writer`.

```go
fmt.Fprintf(os.Stdout, "hello %s\n", "world")
fmt.Fprintf(os.Stderr, "error: %v\n", err)

var buf bytes.Buffer
fmt.Fprintf(&buf, "%d + %d = %d", 1, 2, 3)
fmt.Println(buf.String())  // "1 + 2 = 3"

var sb strings.Builder
fmt.Fprintf(&sb, "item%d ", i)
```

---

## Errorf — Create Formatted Errors

```go
err := fmt.Errorf("user %d not found", 42)

// %w wraps an error for errors.Is / errors.As
base    := fmt.Errorf("connection timeout")
wrapped := fmt.Errorf("fetchUser failed: %w", base)
```

---

## Sscanf / Sscan — Parse from String

```go
var name string
var age  int

fmt.Sscanf("Alice 30", "%s %d", &name, &age)
// name = "Alice", age = 30

var x, y int
fmt.Sscan("10 20", &x, &y)
// x = 10, y = 20
```

For interactive programs, `fmt.Scanf` reads from stdin:

```go
fmt.Scanf("%s %d", &name, &age)
```

---

## Quick Reference

```
%v  %+v  %#v  %T  %%           general
%t                               bool
%d  %b  %o  %O  %x  %X  %c  %U  %q   integer
%f  %F  %e  %E  %g  %G         float
%s  %q  %x  %X                  string
%p                               pointer

width:  %10d  %-10d  %010d  %*d
prec:   %.2f  %10.2f  %.*f
flags:  %+d  % d  %.3s
```
