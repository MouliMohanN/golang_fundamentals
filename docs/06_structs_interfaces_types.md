# Structs, Interfaces & Types

---

## 1. Struct Definition

A struct groups related fields under one named type — Go's primary way to model data.

```go
type Animal struct {
    Name   string
    Sound  string
    Legs   int
    IsTame bool
}
```

---

## 2. Struct Initialisation

```go
// Named fields — preferred, order doesn't matter
cat := Animal{Name: "Cat", Sound: "meow", Legs: 4, IsTame: true}

// Positional — fragile, avoid for structs with more than 2 fields
dog := Animal{"Dog", "woof", 4, true}

// Zero value — all fields get their type's zero value
var unknown Animal  // Name="", Sound="", Legs=0, IsTame=false
```

> Always use named fields. Positional init breaks silently if fields are reordered.

---

## 3. Methods — Value vs Pointer Receiver

A method is a function attached to a type.

```go
// Value receiver — gets a copy, cannot mutate the original
func (a Animal) Describe() string {
    return fmt.Sprintf("%s says %q", a.Name, a.Sound)
}

// Pointer receiver — gets the original, can mutate it
func (a *Animal) Rename(name string) {
    a.Name = name
}
```

```go
eagle := Animal{Name: "Eagle", Legs: 2}
eagle.Describe()       // read — value receiver fine
eagle.Rename("Hawk")   // mutate — pointer receiver required
```

| Receiver | Gets | Can mutate? | Use when |
|---|---|---|---|
| `(a Animal)` | Copy | No | Read-only operations |
| `(a *Animal)` | Original | Yes | Mutations, large structs |

> If any method uses a pointer receiver, make all methods use pointer receivers for consistency.

---

## 4. Embedded Structs — Composition over Inheritance

Go has no inheritance. Instead, embed one struct inside another. The outer type **promotes** the inner type's fields and methods — you access them directly.

```go
type Location struct {
    Habitat string
    Region  string
}

func (l Location) Where() string {
    return fmt.Sprintf("%s in %s", l.Habitat, l.Region)
}

type WildAnimal struct {
    Animal              // embedded — Animal's fields and methods promoted
    Location            // embedded — Location's fields and methods promoted
    EndangeredLevel int // plain field — belongs only to WildAnimal, not promoted from anywhere
}
```

`EndangeredLevel int` is just a regular field on `WildAnimal`. It's not embedded or promoted — it's data that only makes sense for wild animals, not for `Animal` or `Location` in general.

```go
wolf := WildAnimal{
    Animal:          Animal{Name: "Wolf", Legs: 4},
    Location:        Location{Habitat: "forest", Region: "North America"},
    EndangeredLevel: 2,
}
```

### How promotion works — Go does it automatically

Promotion is a **language feature, not something you write**. The moment you embed a struct, all its exported fields and methods are accessible directly on the outer type. No code needed to enable it.

When Go sees `wolf.Name`, it automatically resolves it to `wolf.Animal.Name` for you.

### Both access styles work

```go
wolf.Name         // promoted — Go resolves to wolf.Animal.Name automatically
wolf.Animal.Name  // explicit path — always works
```

Note: the embedded field name is `Animal` (capital A — same as the type name). Go uses the type name as the field name when you don't give it one explicitly. So `wolf.animal.Name` (lowercase) would **not compile**.

### When promotion breaks — name conflicts

If two embedded structs have a field or method with the same name, Go refuses to promote either and forces you to use the explicit path:

```go
type A struct{ Name string }
type B struct{ Name string }

type C struct{ A; B }

c := C{}
c.Name    // compile error — ambiguous: both A.Name and B.Name exist
c.A.Name  // fine — explicit path resolves the ambiguity
c.B.Name  // fine
```

**Composition vs inheritance:** embedding means "has a", not "is a". A `WildAnimal` has an `Animal`, it doesn't inherit from one.

---

## 5. Anonymous Structs

A struct without a named type — useful for one-off data grouping or table-driven tests.

```go
point := struct {
    X, Y int
}{X: 3, Y: 7}

// Common in tests
tests := []struct {
    input    string
    expected int
}{
    {"cat", 3},
    {"eagle", 5},
}
```

---

## 6. Interfaces — Define Behaviour

An interface is a set of method signatures. Any type that implements all those methods **satisfies** the interface — no explicit declaration needed. This is called **implicit / structural typing**.

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

```go
type Circle struct{ Radius float64 }
func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

type Rectangle struct{ Width, Height float64 }
func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }
```

`Circle` and `Rectangle` both satisfy `Shape` without ever writing `implements Shape`.

---

## 7. Polymorphism

One function that works for any type satisfying the interface:

```go
func printShapeInfo(s Shape) {
    fmt.Printf("%T → area=%.2f  perimeter=%.2f\n", s, s.Area(), s.Perimeter())
}

shapes := []Shape{Circle{5}, Rectangle{4, 6}, Triangle{3, 4, 5}}
for _, s := range shapes {
    printShapeInfo(s) // works for all three — same function call
}
```

---

## 8. Composed Interfaces

Interfaces can embed other interfaces to combine behaviour requirements.

```go
type Mover   interface { Move() string }
type Speaker interface { Speak() string }

type LiveAnimal interface {
    Mover   // must implement Move()
    Speaker // must implement Speak()
}

func describeAnimal(a LiveAnimal) {
    fmt.Println(a.Move(), "|", a.Speak())
}
```

Any type implementing both `Move()` and `Speak()` satisfies `LiveAnimal`.

### Embedded Structs vs Composed Interfaces

These look similar but solve completely different problems:

| | Embedded Structs | Composed Interfaces |
|---|---|---|
| Works with | Concrete types | Abstract contracts |
| Contains data? | Yes — fields live in memory | No — method signatures only |
| Purpose | Reuse fields and behaviour | Combine behaviour requirements |
| What gets promoted | Fields AND methods | Nothing — just a rule |

**Embedded struct** = **"has a"** — `WildAnimal` physically contains an `Animal`. Its fields live in memory inside `WildAnimal`.

```go
type WildAnimal struct {
    Animal    // I HAVE an Animal — Name, Legs etc. stored in memory inside me
    Location  // I HAVE a Location — Habitat, Region stored in memory inside me
}
wolf.Name    // real field stored in memory
```

**Composed interface** = **"can do"** — `LiveAnimal` just says "you must implement both `Mover` and `Speaker`". No data, no memory, purely a contract.

```go
type LiveAnimal interface {
    Mover    // you must be able to Move()
    Speaker  // you must be able to Speak()
}
// LiveAnimal has zero fields — it's a behaviour contract only
```

---

## 9. The `Stringer` Interface

`fmt.Stringer` is a standard library interface:

```go
type Stringer interface {
    String() string
}
```

If your type implements `String() string`, `fmt.Println`, `%v`, and `%s` all use it automatically.

```go
func (a Animal) String() string {
    return fmt.Sprintf("Animal(%s, legs=%d)", a.Name, a.Legs)
}

cat := Animal{Name: "Cat", Legs: 4}
fmt.Println(cat)   // Animal(Cat, legs=4)  — no explicit call needed
```

---

## 10. Type Assertion

Extract the concrete type from an interface value. Always use the **safe (comma-ok) form**.

```go
var s Shape = Circle{Radius: 3}

// Safe form — ok is false if assertion fails, no panic
c, ok := s.(Circle)
fmt.Println(ok, c)  // true, {3}

r, ok := s.(Rectangle)
fmt.Println(ok, r)  // false, {0 0}

// Panicking form — only when you are 100% certain
c2 := s.(Circle)  // panics if s is not a Circle
```

---

## 11. Type Switch

Cleanly handle multiple possible concrete types behind an interface.

```go
func describe(s Shape) {
    switch v := s.(type) {
    case Circle:
        fmt.Printf("Circle radius=%.2f\n", v.Radius)
    case Rectangle:
        fmt.Printf("Rectangle %.2f×%.2f\n", v.Width, v.Height)
    default:
        fmt.Printf("Unknown: %T\n", v)
    }
}
```

`v` inside each `case` is already the concrete type — no extra assertion needed.

---

## 12. Custom Named Types

`type X Y` creates a **new distinct type** based on `Y`. They are not interchangeable even though the underlying type is the same.

```go
type Celsius    float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}
```

```go
var temp Celsius = 100
// var f Fahrenheit = temp  // compile error — different types
var f Fahrenheit = temp.ToFahrenheit()  // must convert explicitly
```

**Named types vs type aliases:**

| | `type X Y` | `type X = Y` |
|---|---|---|
| New type? | Yes — distinct | No — just another name |
| Methods allowed? | Yes | Only if Y already has them |
| Interchangeable with Y? | No — must cast | Yes |
| Use when | Domain modelling (Celsius, Direction) | Compatibility / readability aliases |

```go
// Enum pattern using named type + iota
type Direction int
const (
    North Direction = iota
    East
    South
    West
)

func (d Direction) String() string {
    return [...]string{"North", "East", "South", "West"}[d]
}

fmt.Println(North)  // "North" — uses String() automatically
```

---

## Quick Reference

| Concept | Syntax |
|---|---|
| Define struct | `type T struct { Field Type }` |
| Named init | `T{Field: value}` |
| Value receiver | `func (t T) Method()` |
| Pointer receiver | `func (t *T) Method()` |
| Embed struct | `type Outer struct { Inner }` |
| Anonymous struct | `struct{ X int }{X: 1}` |
| Define interface | `type I interface { Method() T }` |
| Satisfy interface | Implement all methods — no keyword needed |
| Compose interfaces | `type I interface { A; B }` |
| Type assertion (safe) | `v, ok := x.(T)` |
| Type assertion (panic) | `v := x.(T)` |
| Type switch | `switch v := x.(type) { case T: }` |
| Named type | `type X Y` |
| Type alias | `type X = Y` |
| Stringer | `func (t T) String() string` |
