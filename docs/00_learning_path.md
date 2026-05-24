# Go Learning Path

A structured guide to learning Go from scratch. Each topic links to its doc and has a matching runnable example in `src/`.

**Progress:** `01` ✓ `02` ✓ `03` ✓ `04` ✓ `05` ✓ `06` ✓ `07` ✓

---

## How to run any example

```bash
go run src/<folder>/main.go
```

---

## Modules

### 01 — Variables & Types
**Doc:** `docs/01_variables_and_types.md` | **Code:** `src/01_variables_and_types/main.go`

- Implicit (`:=`) vs explicit (`var`) declaration
- Zero values
- Multiple assignment and swap
- Blank identifier `_`
- Constants and `iota` (enums, bit flags)
- Printing types at runtime with `%T`
- Numeric type casting
- String conversions (`[]byte`, `[]rune`, `rune → char`)
- `strconv` — converting between strings and primitives

---

### 02 — fmt & Printing
**Doc:** `docs/02_fmt_printing.md` | **Code:** `src/02_fmt_printing/main.go`

- `Print` vs `Println` vs `Printf`
- General verbs: `%v`, `%+v`, `%#v`, `%T`, `%%`
- Boolean: `%t`
- Integer verbs: `%d`, `%b`, `%o`, `%O`, `%x`, `%X`, `%c`, `%U`, `%q`
- Float verbs: `%f`, `%e`, `%E`, `%g`, `%G` — compact explained
- `%c` vs `%q` — raw char vs quoted/escaped
- String verbs: `%s`, `%q`, `%x`
- Width, precision, padding, flags
- `Sprintf`, `Sprint`, `Sprintln`
- `Fprintf`, `Fprint`, `Fprintln` — writing to any `io.Writer`
- `Errorf` with `%w` wrapping
- `Sscanf`, `Sscan`, `Scanf`

---

### 03 — Loops
**Doc:** `docs/03_loops.md` | **Code:** `src/03_loops/main.go`

- C-style `for` (init; condition; post)
- While-style `for` (condition only)
- Infinite loop (`for { }`)
- `continue` and `break`
- `range` over slice, string, map
- String range gotcha — byte offset vs character index
- Nested loops
- Labeled `break` and `continue`

---

### 04 — Functions & Error Handling ✓
**Doc:** `docs/04_functions_and_errors.md` | **Code:** `src/04_functions_and_errors/main.go`

- Function declaration and calling
- Multiple return values
- Named return values
- Variadic functions (`...`)
- Functions as first-class values
- Higher-order functions
- Closures
- `defer` — single and multiple (LIFO order)
- `panic` and `recover`
- The `error` interface
- `errors.New` and `fmt.Errorf`
- Custom error types
- `errors.Is` and `errors.As`
- Wrapping errors with `%w`

---

### 05 — Pointers ✓
**Doc:** `docs/05_pointers.md` | **Code:** `src/05_pointers/main.go`

- What is a pointer (`&`, `*`)
- Dereferencing
- Pass by value vs pass by pointer
- Pointer to struct
- `new()` vs `&T{}`
- Nil pointers and safety
- When to use pointers (mutation, large structs, optional values)
- No pointer arithmetic in Go

---

### 06 — Structs, Interfaces & Types ✓
**Doc:** `docs/06_structs_interfaces_types.md` | **Code:** `src/06_structs_interfaces_types/main.go`

- Struct definition and initialization
- Value vs pointer receivers on methods
- Embedded structs (composition over inheritance)
- Anonymous structs
- Interface definition and implicit implementation
- Polymorphism via interfaces
- Type assertion (`x.(T)`)
- Type switch
- The `Stringer` interface (`fmt.Stringer`)
- Custom named types

---

### 07 — Maps & Generics ✓
**Doc:** `docs/07_maps_and_generics.md` | **Code:** `src/07_maps_and_generics/main.go`

- Map declaration and initialization
- CRUD — create, read, update, delete keys
- Checking key existence (comma-ok idiom)
- Iterating maps (random order)
- Maps of slices, nested maps
- Generic functions with type parameters
- Generic constraints (`comparable`, `any`, custom)
- Generic types (generic structs)

---

## Suggested order for a beginner

```
01 → 02 → 03 → 04 → 05 → 06 → 07
```

Pointers (05) will make more sense after functions (04) since Go passes arguments by value — understanding that gap is what motivates pointers. Interfaces (06) build directly on structs and methods, so do those together.
