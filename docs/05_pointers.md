# Pointers

A pointer stores the **memory address** of another variable instead of the value itself.

---

## 1. The Two Pointer Operators

| Operator | Name | What it does |
|---|---|---|
| `&x` | address-of | gives the memory address of `x` |
| `*p` | dereference | gives the value stored at address `p` |

```go
x := 42
p := &x       // p holds the address of x

fmt.Println(x)   // 42        — the value
fmt.Println(&x)  // 0xc000014090 — the address
fmt.Println(p)   // 0xc000014090 — same address
fmt.Println(*p)  // 42        — dereference: value at that address

*p = 100         // change x through the pointer
fmt.Println(x)   // 100
```

Memory picture:
```
 x  │ 42  │ address: 0xc000014090
    └─────┘
      ↑
 p  holds  0xc000014090
```

---

## 2. Pass by Value vs Pass by Pointer

Go passes **everything by value** — every function call gets a copy. To mutate the original, pass a pointer.

```go
func doubleByValue(n int) {
    n *= 2   // modifies only the local copy — caller sees no change
}

func doubleByPointer(n *int) {
    *n *= 2  // dereferences and modifies the original
}

a := 10
doubleByValue(a)    // a is still 10
doubleByPointer(&a) // a is now 20
```

**Why this matters:** without understanding pass-by-value, you'll write functions that appear to work but silently do nothing to the original data.

---

## 3. Pointer to Struct

Go **auto-dereferences** struct pointers — `p.Field` is exactly the same as `(*p).Field`. You never need to write the explicit dereference.

```go
type Animal struct {
    Name string
    Legs int
}

func birthday(a *Animal) {
    a.Legs++  // Go auto-dereferences — no need for (*a).Legs++
}

cat := Animal{Name: "Cat", Legs: 4}
birthday(&cat)
fmt.Println(cat.Legs)  // 5
```

---

## 4. `new()` vs `&T{}`

Both allocate memory and return a pointer. Prefer `&T{}` — it lets you initialise fields.

```go
p1 := new(Animal)                      // *Animal — zero values (Name="", Legs=0)
p2 := &Animal{Name: "Eagle", Legs: 2}  // *Animal — initialised fields
```

| | `new(T)` | `&T{}` |
|---|---|---|
| Returns | `*T` | `*T` |
| Initial value | Zero value | Whatever you set |
| Field init | Not possible | Supported |
| Prefer | Rarely | Almost always |

---

## 5. Nil Pointers

A pointer's zero value is `nil` — it points to nothing. Dereferencing `nil` causes a **panic**.

```go
var p *Animal   // nil

fmt.Println(p == nil)  // true

p.Name  // PANIC: runtime error: invalid memory address or nil pointer dereference
```

Always guard before dereferencing an uncertain pointer:

```go
if p != nil {
    fmt.Println(p.Name)
}
```

---

## 6. Pointer as Optional Value

A `*T` can represent "value or nothing" — `nil` means the value is absent. This is Go's idiomatic way to express optional fields.

```go
type Config struct {
    Timeout *int   // nil = not configured, non-nil = explicitly set
}

t := 30
c1 := Config{}             // Timeout is nil
c2 := Config{Timeout: &t}  // Timeout is set to 30

if c2.Timeout == nil {
    // use default
} else {
    fmt.Println(*c2.Timeout)  // 30
}
```

---

## 7. When to Use Pointers

| Situation | Use pointer? |
|---|---|
| Need to mutate the original variable | Yes |
| Large struct (avoid expensive copy) | Yes |
| Optional / absent value | Yes (`*T`, nil = absent) |
| Method that modifies struct state | Yes (pointer receiver) |
| Small value, read-only | No — pass by value is cheaper |
| Primitive types (`int`, `bool`, `string`) | Rarely |

### Value receiver vs pointer receiver on methods

```go
type Counter struct{ value int }

func (c *Counter) Increment() { c.value++ }     // pointer receiver — mutates
func (c Counter) ValueDouble() int { return c.value * 2 }  // value receiver — read only

c := Counter{}
c.Increment()         // c.value = 1
c.ValueDouble()       // returns 2, c.value still 1
```

> Be consistent: if any method on a type uses a pointer receiver, all methods should use a pointer receiver.

---

## 8. Pointer to Pointer

Valid but uncommon. Each `*` adds one level of indirection.

```go
x  := 42
p  := &x   // *int   — points to x
pp := &p   // **int  — points to the pointer p

fmt.Println(**pp)  // 42  — two levels of dereference
**pp = 99          // changes x through two levels
fmt.Println(x)     // 99
```

---

## 9. No Pointer Arithmetic

Go deliberately forbids pointer arithmetic (`p++`, `p + 1`). This eliminates a whole class of memory bugs common in C/C++. Use slice indexing or `range` instead.

```go
nums := []int{10, 20, 30}

// C-style pointer arithmetic — NOT allowed in Go
// *(ptr++) — illegal

// Go way
for i, v := range nums {
    fmt.Println(i, v)
}
```

---

## Quick Reference

| Expression | Meaning |
|---|---|
| `&x` | Address of `x` — produces `*T` |
| `*p` | Value at address `p` — dereference |
| `var p *T` | Nil pointer (zero value) |
| `p == nil` | Check if pointer is unset |
| `new(T)` | Allocate zero-valued `T`, return `*T` |
| `&T{...}` | Allocate initialised `T`, return `*T` |
| `p.Field` | Auto-dereference (same as `(*p).Field`) |
| `func (r *T) M()` | Pointer receiver — can mutate `T` |
| `func (r T) M()` | Value receiver — read only copy |
