package main

import (
	"fmt"
	"math"
)

// =============================================================================
// 1. Struct definition
//    A struct groups related fields under one named type
// =============================================================================

type Animal struct {
	Name   string
	Sound  string
	Legs   int
	IsTame bool
}

// =============================================================================
// 2. Struct initialisation — three styles
// =============================================================================

func structInit() {
	// Named fields (preferred — order doesn't matter, clear intent)
	cat := Animal{Name: "Cat", Sound: "meow", Legs: 4, IsTame: true}

	// Positional (fragile — breaks if fields are reordered, avoid for structs >2 fields)
	dog := Animal{"Dog", "woof", 4, true}

	// Zero value — all fields get their zero values
	var unknown Animal

	fmt.Printf("named:      %+v\n", cat)
	fmt.Printf("positional: %+v\n", dog)
	fmt.Printf("zero:       %+v\n", unknown)
}

// =============================================================================
// 3. Methods — value receiver vs pointer receiver
//    Value receiver  → works on a copy, cannot mutate the original
//    Pointer receiver → works on the original, can mutate it
// =============================================================================

func (a Animal) Describe() string {
	return fmt.Sprintf("%s says %q and has %d legs", a.Name, a.Sound, a.Legs)
}

func (a *Animal) Tame() {
	a.IsTame = true // mutates the original
}

func (a *Animal) Rename(name string) {
	a.Name = name
}

func methods() {
	eagle := Animal{Name: "Eagle", Sound: "screech", Legs: 2, IsTame: false}
	fmt.Println(eagle.Describe())

	eagle.Tame()
	eagle.Rename("Golden Eagle")
	fmt.Printf("after Tame+Rename: %+v\n", eagle)
}

// =============================================================================
// 4. Embedded structs — composition over inheritance
//    Go has no inheritance. You embed one struct inside another to reuse fields
//    and methods. The outer type "promotes" the inner type's fields and methods.
// =============================================================================

type Location struct {
	Habitat string
	Region  string
}

func (l Location) Where() string {
	return fmt.Sprintf("%s in %s", l.Habitat, l.Region)
}

type WildAnimal struct {
	Animal              // embedded — fields and methods promoted to WildAnimal
	Location            // embedded
	EndangeredLevel int
}

func embeddedStructs() {
	wolf := WildAnimal{
		Animal:          Animal{Name: "Wolf", Sound: "howl", Legs: 4, IsTame: false},
		Location:        Location{Habitat: "forest", Region: "North America"},
		EndangeredLevel: 2,
	}

	// Promoted fields — access directly as if they belong to WildAnimal
	fmt.Println(wolf.Name)    // Animal.Name promoted
	fmt.Println(wolf.Habitat) // Location.Habitat promoted

	// Promoted methods
	fmt.Println(wolf.Describe()) // Animal.Describe() promoted
	fmt.Println(wolf.Where())    // Location.Where() promoted

	// Explicit path still works
	fmt.Println(wolf.Animal.Name)
}

// =============================================================================
// 5. Anonymous structs — one-off structs without a named type
//    Useful for grouping data locally without polluting the package namespace
// =============================================================================

func anonymousStructs() {
	// Single value
	point := struct {
		X, Y int
	}{X: 3, Y: 7}
	fmt.Printf("point: %+v\n", point)

	// Slice of anonymous structs — common in table-driven tests
	tests := []struct {
		input    string
		expected int
	}{
		{"cat", 3},
		{"eagle", 5},
		{"shark", 5},
	}
	for _, tt := range tests {
		fmt.Printf("len(%q) = %d, expected %d, pass=%v\n",
			tt.input, len(tt.input), tt.expected, len(tt.input) == tt.expected)
	}
}

// =============================================================================
// 6. Interfaces — define behaviour, not data
//    An interface is a set of method signatures.
//    Any type that implements all the methods satisfies the interface —
//    no explicit declaration needed (implicit / structural typing)
// =============================================================================

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle satisfies Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

// Rectangle satisfies Shape
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

// Triangle satisfies Shape
type Triangle struct {
	A, B, C float64 // side lengths
}

func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2 // semi-perimeter (Heron's formula)
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}
func (t Triangle) Perimeter() float64 { return t.A + t.B + t.C }

// =============================================================================
// 7. Polymorphism — one function works for any type that satisfies the interface
// =============================================================================

func printShapeInfo(s Shape) {
	fmt.Printf("%T → area=%.2f  perimeter=%.2f\n", s, s.Area(), s.Perimeter())
}

func interfaces() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
		Triangle{A: 3, B: 4, C: 5},
	}
	for _, s := range shapes {
		printShapeInfo(s)
	}
}

// =============================================================================
// 8. Interface with multiple behaviours — compose interfaces
// =============================================================================

type Mover interface {
	Move() string
}

type Speaker interface {
	Speak() string
}

// Composed interface — must implement both
type LiveAnimal interface {
	Mover
	Speaker
}

type Dog struct{ Name string }

func (d Dog) Move() string  { return d.Name + " runs" }
func (d Dog) Speak() string { return d.Name + " barks" }

type Fish struct{ Name string }

func (f Fish) Move() string  { return f.Name + " swims" }
func (f Fish) Speak() string { return f.Name + " is silent" }

func describeAnimal(a LiveAnimal) {
	fmt.Printf("%s  |  %s\n", a.Move(), a.Speak())
}

func composedInterfaces() {
	describeAnimal(Dog{Name: "Rex"})
	describeAnimal(Fish{Name: "Nemo"})
}

// =============================================================================
// 9. The Stringer interface — fmt.Stringer
//    If a type implements String() string, fmt uses it automatically
// =============================================================================

func (a Animal) String() string {
	tame := "wild"
	if a.IsTame {
		tame = "tame"
	}
	return fmt.Sprintf("Animal(%s, %s, legs=%d, %s)", a.Name, a.Sound, a.Legs, tame)
}

func stringer() {
	cat := Animal{Name: "Cat", Sound: "meow", Legs: 4, IsTame: true}
	fmt.Println(cat)        // automatically uses String()
	fmt.Printf("%v\n", cat) // same
	fmt.Printf("%s\n", cat) // same
}

// =============================================================================
// 10. Type assertion — extract the concrete type from an interface
//     Two forms: panicking (x.(T)) and safe comma-ok (x.(T), ok)
// =============================================================================

func typeAssertion() {
	var s Shape = Circle{Radius: 3}

	// Safe form — always use this
	c, ok := s.(Circle)
	fmt.Printf("is Circle: %v  value: %+v\n", ok, c)

	r, ok := s.(Rectangle)
	fmt.Printf("is Rectangle: %v  value: %+v\n", ok, r)

	// Panicking form — only use when you are 100% certain of the type
	// c2 := s.(Circle)  // panics if s is not a Circle
}

// =============================================================================
// 11. Type switch — cleanly handle multiple possible concrete types
// =============================================================================

func describe(s Shape) {
	switch v := s.(type) {
	case Circle:
		fmt.Printf("Circle with radius %.2f\n", v.Radius)
	case Rectangle:
		fmt.Printf("Rectangle %.2f × %.2f\n", v.Width, v.Height)
	case Triangle:
		fmt.Printf("Triangle sides %.2f, %.2f, %.2f\n", v.A, v.B, v.C)
	default:
		fmt.Printf("Unknown shape: %T\n", v)
	}
}

func typeSwitch() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 3, Height: 4},
		Triangle{A: 5, B: 12, C: 13},
	}
	for _, s := range shapes {
		describe(s)
	}
}

// =============================================================================
// 12. Custom named types
//     type X Y creates a new distinct type based on Y
//     Different from type alias (type X = Y) which is just another name
// =============================================================================

type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", float64(c))
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", float64(f))
}

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

func customTypes() {
	boiling := Celsius(100)
	fmt.Printf("%s = %s\n", boiling, boiling.ToFahrenheit())

	body := Fahrenheit(98.6)
	fmt.Printf("%s = %s\n", body, body.ToCelsius())

	dir := North
	fmt.Println("heading:", dir) // uses Direction.String()
}

// =============================================================================
// main
// =============================================================================

func main() {
	fmt.Println("=== 1 & 2. Struct Init ===")
	structInit()

	fmt.Println("\n=== 3. Methods ===")
	methods()

	fmt.Println("\n=== 4. Embedded Structs ===")
	embeddedStructs()

	fmt.Println("\n=== 5. Anonymous Structs ===")
	anonymousStructs()

	fmt.Println("\n=== 6 & 7. Interfaces & Polymorphism ===")
	interfaces()

	fmt.Println("\n=== 8. Composed Interfaces ===")
	composedInterfaces()

	fmt.Println("\n=== 9. Stringer ===")
	stringer()

	fmt.Println("\n=== 10. Type Assertion ===")
	typeAssertion()

	fmt.Println("\n=== 11. Type Switch ===")
	typeSwitch()

	fmt.Println("\n=== 12. Custom Named Types ===")
	customTypes()
}
