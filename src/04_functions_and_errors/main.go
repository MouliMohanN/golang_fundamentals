package main

import (
	"errors"
	"fmt"
	"math"
)

// =============================================================================
// 1. Basic function — name, params, return type
// =============================================================================

func add(a, b int) int {
	return a + b
}

// When consecutive params share a type, you can group them: (a, b int)
func greet(name string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += "Hello, " + name + "! "
	}
	return result
}

// =============================================================================
// 2. Multiple return values — idiomatic in Go, especially for (value, error)
// =============================================================================

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %.2f by zero", a)
	}
	return a / b, nil
}

// =============================================================================
// 3. Named return values — names act as pre-declared variables
// =============================================================================

func minMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return // "naked return" — returns the named variables min and max
}

// =============================================================================
// 4. Variadic functions — accept any number of arguments of the same type
// =============================================================================

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Spread a slice into a variadic function with ...
func spreadExample() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("sum via spread:", sum(nums...))
}

// =============================================================================
// 5. Functions as first-class values — assign, pass, return
// =============================================================================

func apply(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func makeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// =============================================================================
// 6. Closures — a function that captures variables from its outer scope
// =============================================================================

func makeCounter() func() int {
	count := 0          // captured by the closure below
	return func() int {
		count++         // count lives as long as the closure does
		return count
	}
}

// =============================================================================
// 7. defer — schedules a call to run when the surrounding function returns
//    Multiple defers run in LIFO order (last in, first out)
// =============================================================================

func deferExample() {
	fmt.Println("start")
	defer fmt.Println("deferred 1 — runs last")
	defer fmt.Println("deferred 2 — runs second")
	defer fmt.Println("deferred 3 — runs first")
	fmt.Println("end")
}

func readFile(name string) error {
	fmt.Println("opening", name)
	// defer guarantees cleanup even if the function returns early or panics
	defer fmt.Println("closing", name)

	if name == "" {
		return fmt.Errorf("filename cannot be empty")
	}
	fmt.Println("reading", name)
	return nil
}

// =============================================================================
// 8. panic and recover
//    panic — stops normal execution, unwinds the stack running defers
//    recover — catches a panic inside a deferred function
// =============================================================================

func safeDiv(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()
	result = a / b // panics if b == 0 (integer division by zero)
	return
}

// =============================================================================
// 9. The error interface
//    type error interface { Error() string }
//    Any type that has an Error() string method satisfies the error interface
// =============================================================================

// errors.New — simplest way to create an error
var ErrNotFound = errors.New("not found")

// fmt.Errorf — create a formatted error message
func findAnimal(name string) (string, error) {
	animals := map[string]string{
		"cat":   "mammal",
		"eagle": "bird",
		"shark": "fish",
	}
	kind, ok := animals[name]
	if !ok {
		return "", fmt.Errorf("findAnimal: %w", ErrNotFound)
	}
	return kind, nil
}

// =============================================================================
// 10. Custom error types — struct that implements the error interface
// =============================================================================

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %q: %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "cannot be negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "unrealistically large"}
	}
	return nil
}

// =============================================================================
// 11. errors.Is — checks if an error (or any in its chain) matches a target
//     errors.As — extracts a specific error type from the chain
// =============================================================================

func errorsIsAs() {
	_, err := findAnimal("dragon")

	// errors.Is — works even when the error is wrapped with %w
	if errors.Is(err, ErrNotFound) {
		fmt.Println("errors.Is: dragon not found (sentinel matched)")
	}

	err2 := validateAge(-5)

	// errors.As — unwraps and type-asserts in one step
	var valErr *ValidationError
	if errors.As(err2, &valErr) {
		fmt.Printf("errors.As: field=%q message=%q\n", valErr.Field, valErr.Message)
	}
}

// =============================================================================
// 12. Error wrapping chain — %w lets callers inspect the original cause
// =============================================================================

func loadConfig(path string) error {
	_, err := findAnimal(path)
	if err != nil {
		return fmt.Errorf("loadConfig(%q): %w", path, err)
	}
	return nil
}

// =============================================================================
// main
// =============================================================================

func main() {
	// --- basic ---
	fmt.Println("=== Basic ===")
	fmt.Println(add(3, 4))
	fmt.Println(greet("Mouli", 2))

	// --- multiple returns ---
	fmt.Println("\n=== Multiple Returns ===")
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", result)
	}
	_, err = divide(5, 0)
	fmt.Println("divide by zero:", err)

	// --- named returns ---
	fmt.Println("\n=== Named Returns ===")
	min, max := minMax([]int{3, 1, 7, 2, 9, 4})
	fmt.Printf("min=%d max=%d\n", min, max)

	// --- variadic ---
	fmt.Println("\n=== Variadic ===")
	fmt.Println("sum(1,2,3):", sum(1, 2, 3))
	fmt.Println("sum():", sum())
	spreadExample()

	// --- first-class functions ---
	fmt.Println("\n=== First-Class Functions ===")
	multiply := func(a, b int) int { return a * b }
	fmt.Println("apply add:", apply(3, 4, add))
	fmt.Println("apply multiply:", apply(3, 4, multiply))

	double := makeMultiplier(2)
	triple := makeMultiplier(3)
	fmt.Println("double(5):", double(5))
	fmt.Println("triple(5):", triple(5))

	// --- closures ---
	fmt.Println("\n=== Closures ===")
	counter := makeCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
	other := makeCounter() // independent counter, its own count variable
	fmt.Println(other())   // 1

	// --- defer ---
	fmt.Println("\n=== Defer ===")
	deferExample()
	fmt.Println()
	readFile("config.yaml")
	fmt.Println()
	readFile("")

	// --- panic / recover ---
	fmt.Println("\n=== Panic / Recover ===")
	res, err := safeDiv(10, 2)
	fmt.Printf("safeDiv(10,2): result=%d err=%v\n", res, err)
	res, err = safeDiv(10, 0)
	fmt.Printf("safeDiv(10,0): result=%d err=%v\n", res, err)

	// --- errors ---
	fmt.Println("\n=== Errors ===")
	kind, err := findAnimal("eagle")
	fmt.Printf("eagle: kind=%s err=%v\n", kind, err)
	_, err = findAnimal("dragon")
	fmt.Println("dragon:", err)

	// --- custom error type ---
	fmt.Println("\n=== Custom Error Type ===")
	fmt.Println(validateAge(25))
	fmt.Println(validateAge(-1))
	fmt.Println(validateAge(200))

	// --- errors.Is / errors.As ---
	fmt.Println("\n=== errors.Is / errors.As ===")
	errorsIsAs()

	// --- error wrapping chain ---
	fmt.Println("\n=== Error Wrapping Chain ===")
	err = loadConfig("phoenix")
	fmt.Println(err)
	fmt.Println("Is ErrNotFound?", errors.Is(err, ErrNotFound))

	// --- math.Sqrt as a function value ---
	fmt.Println("\n=== Function Value (stdlib) ===")
	ops := []func(float64) float64{math.Sqrt, math.Log, math.Sin}
	labels := []string{"Sqrt", "Log", "Sin"}
	for i, op := range ops {
		fmt.Printf("%-6s(2) = %.4f\n", labels[i], op(2))
	}
}
