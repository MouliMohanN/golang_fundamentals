package main

import (
	"fmt"
	"strconv"
)

func main() {
	// -------------------------------------------------------------------------
	// 1. Implicit declaration — Go infers type from the assigned value
	// -------------------------------------------------------------------------
	name := "Mouli"
	age := 25
	pi := 3.14159
	isActive := true

	fmt.Printf("%-12s %-10v type: %T\n", "name", name, name)
	fmt.Printf("%-12s %-10v type: %T\n", "age", age, age)
	fmt.Printf("%-12s %-10v type: %T\n", "pi", pi, pi)
	fmt.Printf("%-12s %-10v type: %T\n", "isActive", isActive, isActive)

	// -------------------------------------------------------------------------
	// 2. Explicit declaration — use var when type matters or value comes later
	// -------------------------------------------------------------------------
	var score int = 100
	var label string
	label = "gopher"

	var a, b, c int = 1, 2, 3 // multiple same-type vars in one line

	fmt.Println("\n-- explicit var --")
	fmt.Println(score, label, a, b, c)

	// -------------------------------------------------------------------------
	// 3. Zero values — every type has a default zero value
	// -------------------------------------------------------------------------
	var zInt int
	var zFloat float64
	var zBool bool
	var zString string
	var zPointer *int

	fmt.Println("\n-- zero values --")
	fmt.Printf("int=%d, float64=%f, bool=%t, string=%q, *int=%v\n",
		zInt, zFloat, zBool, zString, zPointer)

	// -------------------------------------------------------------------------
	// 4. Multiple assignment and swap
	// -------------------------------------------------------------------------
	x, y := 10, 20
	x, y = y, x // swap without temp variable
	fmt.Println("\n-- swap --")
	fmt.Println("x=", x, "y=", y)

	// -------------------------------------------------------------------------
	// 5. Blank identifier — discard unwanted values
	// -------------------------------------------------------------------------
	quotient, _ := divide(10, 3) // ignore remainder
	fmt.Println("\n-- blank identifier --")
	fmt.Println("quotient:", quotient)

	// -------------------------------------------------------------------------
	// 6. Constants and iota
	// -------------------------------------------------------------------------
	const gravity = 9.81
	const appName = "GoFundamentals"

	fmt.Println("\n-- constants --")
	fmt.Println(appName, "gravity:", gravity)
	fmt.Println("Sunday=", Sunday, "Monday=", Monday, "Saturday=", Saturday)
	fmt.Println("KB=", KB, "MB=", MB, "GB=", GB)

	// -------------------------------------------------------------------------
	// 7. Printing the type at runtime using %T
	// -------------------------------------------------------------------------
	values := []any{42, 3.14, "hello", true, []int{1, 2}, map[string]int{"a": 1}}
	fmt.Println("\n-- runtime types --")
	for _, v := range values {
		fmt.Printf("value=%-20v type=%T\n", v, v)
	}

	// -------------------------------------------------------------------------
	// 8. Numeric type casting (always explicit in Go)
	// -------------------------------------------------------------------------
	var i int = 100
	f := float64(i)           // int → float64
	back := int(f * 1.9)      // float64 → int  (truncates toward zero)
	var i32 int32 = int32(i)  // int → int32
	var u uint = uint(i)      // int → uint

	fmt.Println("\n-- numeric casting --")
	fmt.Printf("int=%d  float64=%v  int(f*1.9)=%d  int32=%d  uint=%d\n",
		i, f, back, i32, u)

	// -------------------------------------------------------------------------
	// 9. String conversions
	// -------------------------------------------------------------------------
	str := "hello, 世界"

	// string ↔ []byte  (raw UTF-8 bytes)
	bytes := []byte(str)
	fromBytes := string(bytes)

	// string ↔ []rune  (Unicode code points — correct for multibyte chars)
	runes := []rune(str)
	fromRunes := string(runes)

	// rune (int32) → single character string
	ch := string(rune(72)) // code point 72 = 'H'

	fmt.Println("\n-- string conversions --")
	fmt.Printf("[]byte  len=%d  bytes=%v\n", len(bytes), bytes)
	fmt.Printf("[]rune  len=%d  runes=%v\n", len(runes), runes)
	fmt.Printf("fromBytes=%q  fromRunes=%q\n", fromBytes, fromRunes)
	fmt.Printf("rune(72)=%q\n", ch)

	// -------------------------------------------------------------------------
	// 10. strconv — convert between strings and primitive types
	// -------------------------------------------------------------------------
	fmt.Println("\n-- strconv --")

	// int ↔ string
	numStr := strconv.Itoa(42)        // int → string "42"
	num, _ := strconv.Atoi("123")     // string → int

	// float ↔ string
	fStr := strconv.FormatFloat(3.14159, 'f', 2, 64) // "3.14"
	fParsed, _ := strconv.ParseFloat("2.718", 64)

	// bool ↔ string
	bStr := strconv.FormatBool(true)
	bParsed, _ := strconv.ParseBool("true")

	fmt.Printf("Itoa=%q  Atoi=%d\n", numStr, num)
	fmt.Printf("FormatFloat=%q  ParseFloat=%v\n", fStr, fParsed)
	fmt.Printf("FormatBool=%q  ParseBool=%v\n", bStr, bParsed)
}

// -------------------------------------------------------------------------
// Package-level constants using iota
// -------------------------------------------------------------------------
type Weekday int

const (
	Sunday Weekday = iota // 0
	Monday                // 1
	Tuesday               // 2
	Wednesday             // 3
	Thursday              // 4
	Friday                // 5
	Saturday              // 6
)

const (
	KB = 1 << (10 * (iota + 1)) // 1024
	MB                           // 1048576
	GB                           // 1073741824
)

func divide(a, b int) (int, int) {
	return a / b, a % b
}
