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

### `%c` vs `%q`

| | `%c` | `%q` |
|---|---|---|
| Output | Raw character, no quotes | Character wrapped in single quotes, escaped |
| Unprintable chars | Invisible / breaks layout | Always human-readable (`'\t'`, `'\n'`) |
| Use when | Outputting the character itself | Debugging or inspecting what a character is |

```go
fmt.Printf("%c", 72)   // H        — raw character
fmt.Printf("%q", 72)   // 'H'      — quoted character

fmt.Printf("%c", 9)    // 	        — raw tab (invisible)
fmt.Printf("%q", 9)    // '\t'     — escaped, readable

fmt.Printf("%c", 10)   //          — raw newline (breaks the line)
fmt.Printf("%q", 10)   // '\n'     — escaped, readable
```

> Prefer `%q` when debugging — it never produces invisible or layout-breaking output.

---

## Float Verbs

| Verb | Meaning | Example (f=123456.789) |
|---|---|---|
| `%f` | Decimal notation | `123456.789000` |
| `%F` | Same, uppercase `INF`/`NAN` | `123456.789000` |
| `%e` | Scientific, lowercase `e` | `1.234568e+05` |
| `%E` | Scientific, uppercase `E` | `1.234568E+05` |
| `%g` | Compact — auto-picks `%e` or `%f`, lowercase | `123456.789` |
| `%G` | Compact — auto-picks `%E` or `%f`, uppercase | `123456.789` |

```go
fmt.Printf("%e\n", 123456.789)  // 1.234568e+05
fmt.Printf("%.2f\n", 3.14159)   // 3.14
fmt.Printf("%g\n", 0.00001)     // 1e-05
```

### What "compact" means (`%g` / `%G`)

`%g` doesn't force decimal or scientific — it **picks whichever is shorter** for that number.

Go's rule: use `%e` if the exponent is **< -4** or **>= precision (default 6)**, otherwise use `%f`.

```go
// Normal number — %g stays decimal, cleaner than forced scientific
fmt.Printf("%f\n", 123.456)   // 123.456000   — trailing zeros, noisy
fmt.Printf("%e\n", 123.456)   // 1.234560e+02 — overkill for a simple number
fmt.Printf("%g\n", 123.456)   // 123.456       — clean, no noise

// Very small number — %g switches to scientific, easier to read
fmt.Printf("%f\n", 0.000001)  // 0.000001      — easy to miscount the zeros
fmt.Printf("%e\n", 0.000001)  // 1.000000e-06  — clear but always forces scientific
fmt.Printf("%g\n", 0.000001)  // 1e-06          — shortest, still clear

// Very large number — %g switches to scientific automatically
fmt.Printf("%f\n", 99999999.9)  // 99999999.900000 — verbose
fmt.Printf("%g\n", 99999999.9)  // 9.9999999e+07   — switches only when needed
```

> Use `%g` when you don't know ahead of time whether a number will be tiny or huge — it adapts. Use `%f` when you always want decimal (e.g. currency). Use `%e` when you always want scientific (e.g. physics output).

---

### Lowercase vs Uppercase (`%e` vs `%E`, `%g` vs `%G`)

The numbers are identical — only the exponent letter differs.

```go
fmt.Printf("%e\n", 123456.789)  // 1.234568e+05
fmt.Printf("%E\n", 123456.789)  // 1.234568E+05
fmt.Printf("%g\n", 0.00001)     // 1e-05
fmt.Printf("%G\n", 0.00001)     // 1E-05
```

| Case | When to use |
|---|---|
| Lowercase (`%e`, `%g`) | Logs, terminals, JSON, general programming output — the default |
| Uppercase (`%E`, `%G`) | Scientific papers, financial reports, Excel exports, APIs that expect uppercase `E` notation |

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
