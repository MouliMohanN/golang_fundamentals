package main

import "fmt"

// =============================================================================
// 1. What is a pointer
//    A pointer stores the memory address of another variable
//    &variable  → gives the address (reference)
//    *pointer   → gives the value at that address (dereference)
// =============================================================================

func basicPointers() {
	x := 42
	p := &x // p holds the memory address of x

	fmt.Printf("x      = %d\n", x)
	fmt.Printf("&x     = %p  (address of x)\n", &x)
	fmt.Printf("p      = %p  (p holds the same address)\n", p)
	fmt.Printf("*p     = %d  (dereference: value at that address)\n", *p)

	*p = 100 // change x through the pointer
	fmt.Printf("x after *p = 100 → x = %d\n", x)
}

// =============================================================================
// 2. Pass by value vs pass by pointer
//    Go passes everything by VALUE — the function gets a copy
//    To mutate the original, pass a pointer
// =============================================================================

func doubleByValue(n int) {
	n *= 2 // only modifies the local copy
}

func doubleByPointer(n *int) {
	*n *= 2 // dereferences and modifies the original
}

func passByValueVsPointer() {
	a := 10
	doubleByValue(a)
	fmt.Println("after doubleByValue:", a) // still 10

	b := 10
	doubleByPointer(&b)
	fmt.Println("after doubleByPointer:", b) // 20
}

// =============================================================================
// 3. Pointer to struct
//    Go auto-dereferences struct pointers — p.Field is the same as (*p).Field
// =============================================================================

type Animal struct {
	Name string
	Legs int
}

func birthday(a *Animal) {
	a.Legs++ // same as (*a).Legs++ — Go handles the dereference automatically
}

func pointerToStruct() {
	cat := Animal{Name: "Cat", Legs: 4}
	fmt.Printf("before: %+v\n", cat)
	birthday(&cat)
	fmt.Printf("after:  %+v\n", cat)
}

// =============================================================================
// 4. new() vs &T{}
//    Both allocate memory and return a pointer
//    new(T)  → allocates zero-valued T, returns *T
//    &T{}    → allocates T with field initialisation, returns *T  (preferred)
// =============================================================================

func newVsLiteral() {
	p1 := new(Animal)                          // *Animal, zero values: Name="", Legs=0
	p2 := &Animal{Name: "Eagle", Legs: 2}      // *Animal, initialised

	fmt.Printf("new(Animal)       = %+v\n", *p1)
	fmt.Printf("&Animal{...}      = %+v\n", *p2)
}

// =============================================================================
// 5. Nil pointers
//    A pointer's zero value is nil — dereferencing nil panics
// =============================================================================

func nilPointers() {
	var p *Animal // nil — points to nothing

	fmt.Println("p is nil:", p == nil)

	// Always guard before dereferencing
	if p != nil {
		fmt.Println(p.Name)
	} else {
		fmt.Println("cannot dereference nil pointer — guarded safely")
	}
}

// =============================================================================
// 6. Pointer as optional value
//    A *T can represent "value or nothing" — nil means absent
// =============================================================================

type Config struct {
	Timeout *int // nil = "not set", non-nil = explicitly configured
}

func printTimeout(c Config) {
	if c.Timeout == nil {
		fmt.Println("Timeout: not set (using default)")
	} else {
		fmt.Printf("Timeout: %d seconds\n", *c.Timeout)
	}
}

func optionalValue() {
	t := 30
	c1 := Config{}           // Timeout not set
	c2 := Config{Timeout: &t} // Timeout explicitly set

	printTimeout(c1)
	printTimeout(c2)
}

// =============================================================================
// 7. When to use pointers
//    - You need to mutate the original variable
//    - The struct is large and copying is expensive
//    - You need to represent an optional / absent value (nil)
//    - Methods that modify state use pointer receivers
// =============================================================================

type Counter struct {
	value int
}

// Pointer receiver — modifies the actual Counter
func (c *Counter) Increment() {
	c.value++
}

// Value receiver — works on a copy, original unchanged
func (c Counter) ValueDouble() int {
	return c.value * 2
}

func whenToUsePointers() {
	c := Counter{}
	c.Increment()
	c.Increment()
	fmt.Println("after 2 increments:", c.value)          // 2
	fmt.Println("ValueDouble (no mutation):", c.ValueDouble()) // 4
	fmt.Println("value unchanged:", c.value)               // still 2
}

// =============================================================================
// 8. Pointers to pointers
//    Uncommon, but valid — *T, **T, etc.
// =============================================================================

func pointerToPointer() {
	x := 42
	p := &x   // *int
	pp := &p  // **int — pointer to the pointer

	fmt.Printf("x=%d  *p=%d  **pp=%d\n", x, *p, **pp)

	**pp = 99 // change x through two levels of indirection
	fmt.Printf("x after **pp=99 → x=%d\n", x)
}

// =============================================================================
// 9. No pointer arithmetic
//    Unlike C, Go does not allow p++ or p+1 on pointers
//    This is intentional — it prevents a whole class of memory bugs
// =============================================================================

func noPointerArithmetic() {
	nums := []int{10, 20, 30}
	// In C you could do: ptr++  to advance to next element
	// In Go, use index or range instead
	for i := range nums {
		fmt.Printf("nums[%d] = %d  addr = %p\n", i, nums[i], &nums[i])
	}
}

// =============================================================================
// main
// =============================================================================

func main() {
	fmt.Println("=== 1. Basic Pointers ===")
	basicPointers()

	fmt.Println("\n=== 2. Pass by Value vs Pointer ===")
	passByValueVsPointer()

	fmt.Println("\n=== 3. Pointer to Struct ===")
	pointerToStruct()

	fmt.Println("\n=== 4. new() vs &T{} ===")
	newVsLiteral()

	fmt.Println("\n=== 5. Nil Pointers ===")
	nilPointers()

	fmt.Println("\n=== 6. Pointer as Optional Value ===")
	optionalValue()

	fmt.Println("\n=== 7. When to Use Pointers ===")
	whenToUsePointers()

	fmt.Println("\n=== 8. Pointer to Pointer ===")
	pointerToPointer()

	fmt.Println("\n=== 9. No Pointer Arithmetic ===")
	noPointerArithmetic()
}
