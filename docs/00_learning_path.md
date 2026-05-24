# Go Learning Path

A structured guide to learning Go from scratch. Each topic links to its doc and has a matching runnable example in `src/`.

---

## How to run any example

```bash
go run src/<folder>/main.go
```

---

## Where you stand

> Background: 12 years frontend experience, switching to backend full time.

| Stage | Topics | Readiness |
|---|---|---|
| Fundamentals | 01 – 07 (done) | 30% |
| + Concurrency | 08 Goroutines & Channels | 45% |
| + Context | 09 context package | 55% |
| + HTTP | 10 net/http | 65% |
| + JSON | 11 encoding/json | 70% |
| + Testing | 12 Testing | 78% |
| + Slices | 13 Slices in depth | 82% |
| + Structure | 14 Packages & Project Layout | 87% |
| + Database | 15 database/sql | 95% |

The last 5% comes from real-world production experience — logging, metrics, tracing, deployment, and patterns you only learn by building actual services.

**Your FE background accelerates:** you already understand HTTP request/response cycles, JSON, async thinking, and TypeScript interfaces — all of which map directly to Go concepts in modules 09–11.

**The biggest unlock is module 08** — goroutines and channels are Go's killer feature and what separates performant backend services from everything else. Go there next.

---

## Progress: `01` ✓ `02` ✓ `03` ✓ `04` ✓ `05` ✓ `06` ✓ `07` ✓ `08` ○ `09` ○ `10` ○ `11` ○ `12` ○ `13` ○ `14` ○ `15` ○

---

## Phase 1 — Language Fundamentals (30%)

### 01 — Variables & Types ✓
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

### 02 — fmt & Printing ✓
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

### 03 — Loops ✓
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
- Closures + garbage collection of captured variables
- `defer` — single and multiple (LIFO order)
- `panic` and `recover` — step-by-step execution
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
- Pointer to struct + auto-dereference
- `new()` vs `&T{}`
- Nil pointers and safety
- Pointer as optional value
- When to use pointers (mutation, large structs, optional values)
- Pointer to pointer
- No pointer arithmetic in Go

---

### 06 — Structs, Interfaces & Types ✓
**Doc:** `docs/06_structs_interfaces_types.md` | **Code:** `src/06_structs_interfaces_types/main.go`

- Struct definition and initialization
- Value vs pointer receivers on methods
- Embedded structs — promotion, name conflicts, explicit paths
- Embedded structs vs composed interfaces
- Anonymous structs
- Interface definition and implicit implementation
- Polymorphism via interfaces
- Composed interfaces
- Type assertion — safe vs panicking form
- Type switch
- The `Stringer` interface (`fmt.Stringer`)
- Custom named types — `type X Y` vs `type X = Y`

---

### 07 — Maps & Generics ✓
**Doc:** `docs/07_maps_and_generics.md` | **Code:** `src/07_maps_and_generics/main.go`

- Map declaration and initialization
- CRUD — create, read, update, delete keys
- Checking key existence (comma-ok idiom)
- Iterating maps (random order) + sorted iteration pattern
- Maps of slices, nested maps
- Map as a set (`map[T]struct{}`)
- Generic functions with type parameters
- The `~` tilde operator
- Generic constraints (`comparable`, `any`, `cmp.Ordered`, custom)
- Generic types (Stack, Pair)

---

## Phase 2 — Backend Essentials (30% → 95%)

### 08 — Goroutines & Channels _(coming next)_
**Doc:** `docs/08_goroutines_and_channels.md` | **Code:** `src/08_goroutines_and_channels/main.go`

- Goroutines — lightweight threads (`go func()`)
- Channels — typed communication between goroutines
- Buffered vs unbuffered channels
- `select` — wait on multiple channels
- `sync.WaitGroup` — wait for goroutines to finish
- `sync.Mutex` — protect shared state
- Common patterns: fan-out, fan-in, worker pool
- Race conditions and the race detector (`go run -race`)

---

### 09 — Context _(coming)_
**Doc:** `docs/09_context.md` | **Code:** `src/09_context/main.go`

- What `context.Context` is and why every BE function takes one
- `context.Background()` and `context.TODO()`
- Cancellation — `context.WithCancel`
- Timeouts — `context.WithTimeout`, `context.WithDeadline`
- Request-scoped values — `context.WithValue`
- Propagating context through call chains

---

### 10 — net/http _(coming)_
**Doc:** `docs/10_net_http.md` | **Code:** `src/10_net_http/main.go`

- HTTP server with `http.ListenAndServe`
- Handlers and `http.HandlerFunc`
- Request: method, URL, headers, body
- Response: status codes, headers, writing body
- Middleware pattern
- HTTP client — making outbound requests
- Reading and writing JSON over HTTP

---

### 11 — encoding/json _(coming)_
**Doc:** `docs/11_encoding_json.md` | **Code:** `src/11_encoding_json/main.go`

- `json.Marshal` — struct to JSON
- `json.Unmarshal` — JSON to struct
- Struct tags (`json:"name"`, `omitempty`, `-`)
- Streaming — `json.Encoder`, `json.Decoder`
- Handling unknown fields with `map[string]any`
- Custom marshalling with `json.Marshaler` / `json.Unmarshaler`

---

### 12 — Testing _(coming)_
**Doc:** `docs/12_testing.md` | **Code:** `src/12_testing/`

- `testing.T` — writing and running tests
- Table-driven tests — Go's idiomatic test pattern
- `t.Run` — subtests
- Benchmarks with `testing.B`
- Test helpers and `t.Helper()`
- `testify` — assertions library
- Mocking with interfaces

---

### 13 — Slices in Depth _(coming)_
**Doc:** `docs/13_slices.md` | **Code:** `src/13_slices/main.go`

- Slice internals — pointer, length, capacity
- `append` and when it allocates
- `copy`
- Slice of slices (2D)
- Common gotcha — slice sharing underlying array
- `slices` package (Go 1.21+)

---

### 14 — Packages & Project Layout _(coming)_
**Doc:** `docs/14_packages_and_layout.md` | **Code:** `src/14_packages/`

- `go.mod` and modules
- Package naming conventions
- Exported vs unexported identifiers
- Internal packages
- Standard project layout for a REST API
- Dependency management with `go get`, `go mod tidy`

---

### 15 — database/sql _(coming)_
**Doc:** `docs/15_database_sql.md` | **Code:** `src/15_database_sql/main.go`

- Connecting to a database
- Querying — `db.Query`, `db.QueryRow`
- Scanning results into structs
- Exec — insert, update, delete
- Transactions
- Prepared statements
- `sqlx` — ergonomic wrapper over `database/sql`
