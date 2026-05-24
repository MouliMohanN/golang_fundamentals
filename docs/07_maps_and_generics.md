# Maps & Generics

---

## 1. Map Declaration & Initialisation

A map is an **unordered collection of key-value pairs** — like a dictionary or hash table.

```
map[KeyType]ValueType
```

Keys must be **comparable** (support `==`) — strings, ints, bools, and structs work. Slices, maps, and functions cannot be keys.

```go
// make — preferred when populating later
m := make(map[string]int)
m["cat"] = 4

// literal — preferred when values are known upfront
legs := map[string]int{
    "cat":    4,
    "bird":   2,
    "spider": 8,
}

// nil map — reading is safe (returns zero value), writing panics
var m3 map[string]int
fmt.Println(m3["cat"]) // 0 — no panic
m3["cat"] = 4          // panic: assignment to entry in nil map
```

---

## 2. CRUD

```go
legs := map[string]int{"cat": 4, "bird": 2}

// Create
legs["spider"] = 8

// Read
fmt.Println(legs["cat"])  // 4

// Update — same syntax as create
legs["cat"] = 5

// Delete
delete(legs, "bird")

// Length
fmt.Println(len(legs))  // 2
```

---

## 3. Comma-ok Idiom — Key Existence

Reading a missing key returns the **zero value**, not an error. Use the two-value form to tell the difference between "key is missing" and "value happens to be zero".

```go
scores := map[string]int{
    "alice": 95,
    "bob":   0,   // explicitly zero — not missing
}

// Single value — can't distinguish missing from zero
v := scores["carol"]  // 0 — carol doesn't exist, but looks the same as bob

// Comma-ok — definitive answer
score, ok := scores["bob"]
// score=0, ok=true  → bob exists, value is 0

score, ok = scores["carol"]
// score=0, ok=false → carol does not exist
```

---

## 4. Iterating a Map

Map iteration order is **random every run** — by design. Never rely on order.

```go
for key, value := range m { }  // both
for key := range m { }         // keys only
```

To iterate in sorted order, collect keys, sort, then iterate:

```go
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
slices.Sort(keys)
for _, k := range keys {
    fmt.Println(k, m[k])
}
```

---

## 5. Map of Slices

Map values can be any type, including slices. A common pattern is grouping items by a key.

```go
byLegs := map[int][]string{}

for _, a := range animals {
    byLegs[a.legs] = append(byLegs[a.legs], a.name)
}

// byLegs[4] → ["cat", "dog"]
// byLegs[2] → ["bird", "eagle"]
```

`append` on a nil slice works fine — no need to initialise each slice manually.

---

## 6. Nested Maps

```go
speeds := map[string]map[string]int{
    "land":  {"cheetah": 120, "lion": 80},
    "water": {"sailfish": 110},
}
```

Always guard each level with comma-ok before reading deep:

```go
if land, ok := speeds["land"]; ok {
    if speed, ok := land["cheetah"]; ok {
        fmt.Println(speed)
    }
}
```

---

## 7. Map as a Set

Go has no built-in set. Use `map[T]struct{}` — `struct{}` takes **zero bytes** in memory, so you only pay for the keys.

```go
seen := map[string]struct{}{}

for _, w := range words {
    seen[w] = struct{}{}  // struct{}{} is the zero-size empty value
}

// Check membership
_, exists := seen["cat"]

// Remove
delete(seen, "cat")
```

---

## 8. Generics — Type Parameters

Introduced in Go 1.18. Write one function that works for multiple types without duplicating code.

**Without generics — one function per type:**
```go
func sumInts(nums []int) int       { ... }
func sumFloats(nums []float64) float64 { ... }
```

**With generics — one function for all numeric types:**
```go
type Number interface {
    ~int | ~float32 | ~float64  // ~ means "this type or any type whose underlying type is this"
}

func sum[T Number](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}

sum([]int{1, 2, 3})         // T inferred as int
sum([]float64{1.1, 2.2})    // T inferred as float64
```

**Syntax breakdown:**
```
func sum[T Number](nums []T) T
         ────────  ────────  ─
         type      param     return
         param     uses T    type T
         T with
         constraint
         Number
```

### The `~` (tilde) operator

`~int` means "int **or any named type whose underlying type is int**":

```go
type MyInt int        // underlying type is int
sum([]MyInt{1, 2, 3}) // works because ~int covers MyInt

// Without ~, only bare int would work — MyInt would be rejected
```

---

## 9. Built-in Generic Constraints

| Constraint | Meaning |
|---|---|
| `any` | Any type — alias for `interface{}` |
| `comparable` | Any type supporting `==` and `!=` |
| `cmp.Ordered` | Any type supporting `<`, `>`, `<=`, `>=` |

```go
// comparable — works for any type that supports ==
func contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item { return true }
    }
    return false
}
contains([]string{"cat", "dog"}, "dog")  // true
contains([]int{1, 2, 3}, 5)             // false

// any — T can be anything
func Map[T any, U any](slice []T, f func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = f(v)
    }
    return result
}
Map([]int{1, 2, 3}, func(n int) int { return n * 2 }) // [2 4 6]

func Filter[T any](slice []T, predicate func(T) bool) []T { ... }
Filter([]int{1,2,3,4,5,6}, func(n int) bool { return n%2 == 0 }) // [2 4 6]

// cmp.Ordered — any type with < > <= >=
func Min[T cmp.Ordered](a, b T) T { if a < b { return a }; return b }
Min(3, 7)       // 3
Min(3.14, 2.71) // 2.71
```

---

## 10. Generic Types

Structs can also have type parameters.

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T)       { s.items = append(s.items, item) }
func (s *Stack[T]) Pop() (T, bool)    { ... }
func (s *Stack[T]) Peek() (T, bool)   { ... }
func (s *Stack[T]) Len() int          { return len(s.items) }
```

```go
var intStack Stack[int]     // type must be explicit here — no inference on types
intStack.Push(10)
intStack.Push(20)
top, _ := intStack.Pop()    // 20

var strStack Stack[string]
strStack.Push("eagle")
```

**Pair — two different type parameters:**

```go
type Pair[A, B any] struct {
    First  A
    Second B
}

p := NewPair("eagle", 150)  // Pair[string, int]
fmt.Println(p.First, p.Second)  // eagle 150
```

---

## Quick Reference

| Operation | Syntax |
|---|---|
| Declare map | `m := map[K]V{}` or `make(map[K]V)` |
| Set key | `m[k] = v` |
| Get key | `v := m[k]` |
| Key exists? | `v, ok := m[k]` |
| Delete key | `delete(m, k)` |
| Length | `len(m)` |
| Iterate | `for k, v := range m` |
| Set (map) | `map[T]struct{}{}` |
| Generic function | `func f[T Constraint](x T) T` |
| Generic type | `type S[T any] struct { ... }` |
| Tilde constraint | `~int` — int or named type with int underlying |
| comparable | `==` / `!=` supported |
| cmp.Ordered | `<` `>` `<=` `>=` supported |
