package main

import "fmt"

func main() {
	// --- Implicit variable declaration using := ---
	name := "Mouli"       // string inferred
	age := 25             // int inferred
	pi := 3.14            // float64 inferred
	isActive := true      // bool inferred

	// --- Print type using %T verb ---
	fmt.Printf("name=%v,    type: %T\n", name, name)
	fmt.Printf("age=%v,     type: %T\n", age, age)
	fmt.Printf("pi=%v,   type: %T\n", pi, pi)
	fmt.Printf("isActive=%v, type: %T\n", isActive, isActive)

	fmt.Println("---")

	// --- Printf vs Println ---
	// Println adds a newline automatically and spaces between arguments
	fmt.Println("Hello,", "World!", age)

	// Printf requires explicit format verbs; no automatic newline
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

	// Printf is useful for formatted numeric output
	fmt.Printf("Pi to 4 decimal places: %.4f\n", pi)

	fmt.Println("---")

	// --- Type casting (explicit conversion) ---
	var x int = 42
	var y float64 = float64(x) // int → float64
	var z int = int(y * 1.5)   // float64 → int (truncates decimal)

	fmt.Printf("x (int)=%d, y (float64)=%v, z (int)=%d\n", x, y, z)

	// string ↔ byte slice conversion
	str := "hello"
	bytes := []byte(str)           // string → []byte
	backToStr := string(bytes)     // []byte → string

	fmt.Printf("str=%q, bytes=%v, backToStr=%q\n", str, bytes, backToStr)

	// int → string via rune (gives the Unicode character, not the digit)
	char := string(rune(72)) // 72 is 'H' in ASCII
	fmt.Printf("rune(72) as string: %q\n", char)
}
