package main

import (
	"cmp"
	"fmt"
	"slices"
)

// =============================================================================
// 1. Map declaration and initialisation
//    A map is an unordered collection of key-value pairs
//    map[KeyType]ValueType
//    Keys must be comparable (==) — strings, ints, bools, structs (no slices/maps)
// =============================================================================

func mapInit() {
	// Using make — preferred when you know the map will be populated later
	m1 := make(map[string]int)
	m1["cat"] = 4
	m1["bird"] = 2

	// Map literal — preferred when you know the values upfront
	m2 := map[string]int{
		"cat":    4,
		"bird":   2,
		"spider": 8,
	}

	// Zero value of a map is nil — reading is safe, writing panics
	var m3 map[string]int
	fmt.Println("nil map read:", m3["cat"]) // returns zero value, no panic
	// m3["cat"] = 4                        // panic: assignment to entry in nil map

	fmt.Println("make:", m1)
	fmt.Println("literal:", m2)
	fmt.Println("nil map:", m3 == nil)
}

// =============================================================================
// 2. CRUD — create, read, update, delete
// =============================================================================

func mapCRUD() {
	legs := map[string]int{
		"cat":   4,
		"bird":  2,
		"snake": 0,
	}

	// Create / Update — same syntax
	legs["spider"] = 8  // create new key
	legs["cat"] = 5     // update existing key

	// Read
	fmt.Println("cat:", legs["cat"])

	// Delete
	delete(legs, "snake")
	fmt.Println("after delete:", legs)

	// Length
	fmt.Println("len:", len(legs))
}

// =============================================================================
// 3. Comma-ok idiom — check if a key exists
//    Reading a missing key returns the zero value, not an error
//    Use the two-value form to distinguish "key missing" from "value is zero"
// =============================================================================

func commaOk() {
	scores := map[string]int{
		"alice": 95,
		"bob":   0, // explicitly zero — not missing
	}

	// Single value — can't tell if key is missing or value is 0
	v1 := scores["bob"]
	v2 := scores["carol"] // carol doesn't exist — also returns 0
	fmt.Println("bob:", v1, "carol:", v2) // both print 0

	// Comma-ok — distinguishes zero value from missing key
	score, ok := scores["bob"]
	fmt.Printf("bob:   score=%d ok=%v\n", score, ok) // 0, true

	score, ok = scores["carol"]
	fmt.Printf("carol: score=%d ok=%v\n", score, ok) // 0, false
}

// =============================================================================
// 4. Iterating a map — order is random every run
// =============================================================================

func mapIteration() {
	population := map[string]int{
		"India":  1400,
		"China":  1300,
		"USA":    330,
		"Brazil": 215,
	}

	fmt.Println("-- random order (may differ each run) --")
	for country, pop := range population {
		fmt.Printf("%-10s %d million\n", country, pop)
	}

	// To iterate in sorted order, collect keys, sort, then iterate
	keys := make([]string, 0, len(population))
	for k := range population {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	fmt.Println("-- sorted order --")
	for _, k := range keys {
		fmt.Printf("%-10s %d million\n", k, population[k])
	}
}

// =============================================================================
// 5. Map of slices — values can be any type including slices
// =============================================================================

func mapOfSlices() {
	// Group animals by number of legs
	byLegs := map[int][]string{}

	animals := []struct {
		name string
		legs int
	}{
		{"cat", 4}, {"dog", 4}, {"bird", 2},
		{"eagle", 2}, {"spider", 8}, {"snake", 0},
	}

	for _, a := range animals {
		byLegs[a.legs] = append(byLegs[a.legs], a.name)
	}

	fmt.Println("0 legs:", byLegs[0])
	fmt.Println("2 legs:", byLegs[2])
	fmt.Println("4 legs:", byLegs[4])
	fmt.Println("8 legs:", byLegs[8])
}

// =============================================================================
// 6. Nested maps
// =============================================================================

func nestedMaps() {
	// Habitat → animal → speed (km/h)
	speeds := map[string]map[string]int{
		"land": {
			"cheetah": 120,
			"lion":    80,
		},
		"water": {
			"sailfish": 110,
			"shark":    50,
		},
		"air": {
			"peregrine falcon": 389,
			"eagle":            150,
		},
	}

	for habitat, animals := range speeds {
		fmt.Printf("-- %s --\n", habitat)
		for animal, speed := range animals {
			fmt.Printf("  %-20s %d km/h\n", animal, speed)
		}
	}

	// Safe nested read — guard each level with comma-ok
	if land, ok := speeds["land"]; ok {
		if speed, ok := land["cheetah"]; ok {
			fmt.Println("cheetah:", speed, "km/h")
		}
	}
}

// =============================================================================
// 7. Map as a set
//    Go has no built-in set type — use map[T]struct{} (struct{} takes zero bytes)
// =============================================================================

func mapAsSet() {
	seen := map[string]struct{}{}

	words := []string{"cat", "dog", "cat", "bird", "dog", "cat"}
	for _, w := range words {
		seen[w] = struct{}{} // struct{}{} is the zero-size value
	}

	fmt.Println("unique words:", seen)
	fmt.Println("len:", len(seen))

	// Check membership
	_, exists := seen["cat"]
	fmt.Println("cat exists:", exists)
}

// =============================================================================
// 8. Generics — type parameters
//    Introduced in Go 1.18. Write one function that works for multiple types.
//    T is a type parameter, constrained by an interface.
// =============================================================================

// Without generics — need separate functions for each type
func sumInts(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func sumFloats(nums []float64) float64 {
	total := 0.0
	for _, n := range nums {
		total += n
	}
	return total
}

// With generics — one function for both (and any other numeric type)
// ~int means "int or any type whose underlying type is int"
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

func sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

// =============================================================================
// 9. Generic functions with common constraints
// =============================================================================

// comparable — built-in constraint: any type that supports == and !=
func contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// any — built-in constraint: any type at all (alias for interface{})
func Map[T any, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// cmp.Ordered — constraint from std lib: any type that supports < > <= >=
func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// =============================================================================
// 10. Generic types — structs with type parameters
// =============================================================================

// Stack works for any type T
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Len() int { return len(s.items) }

// Pair holds two values of potentially different types
type Pair[A, B any] struct {
	First  A
	Second B
}

func NewPair[A, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{First: a, Second: b}
}

// =============================================================================
// main
// =============================================================================

func main() {
	fmt.Println("=== 1. Map Init ===")
	mapInit()

	fmt.Println("\n=== 2. CRUD ===")
	mapCRUD()

	fmt.Println("\n=== 3. Comma-ok Idiom ===")
	commaOk()

	fmt.Println("\n=== 4. Map Iteration ===")
	mapIteration()

	fmt.Println("\n=== 5. Map of Slices ===")
	mapOfSlices()

	fmt.Println("\n=== 6. Nested Maps ===")
	nestedMaps()

	fmt.Println("\n=== 7. Map as Set ===")
	mapAsSet()

	fmt.Println("\n=== 8. Generics — sum ===")
	fmt.Println("sum ints:", sum([]int{1, 2, 3, 4, 5}))
	fmt.Println("sum floats:", sum([]float64{1.1, 2.2, 3.3}))

	fmt.Println("\n=== 9. Generic Functions ===")
	fmt.Println("contains:", contains([]string{"cat", "dog", "bird"}, "dog"))
	fmt.Println("contains:", contains([]int{1, 2, 3}, 5))

	doubled := Map([]int{1, 2, 3, 4}, func(n int) int { return n * 2 })
	fmt.Println("Map double:", doubled)

	upper := Map([]string{"cat", "dog"}, func(s string) string {
		return "[" + s + "]"
	})
	fmt.Println("Map bracket:", upper)

	evens := Filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 })
	fmt.Println("Filter evens:", evens)

	fmt.Println("Min(3,7):", Min(3, 7))
	fmt.Println("Max(3.14,2.71):", Max(3.14, 2.71))

	fmt.Println("\n=== 10. Generic Types ===")
	// Stack of ints
	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	fmt.Println("stack len:", intStack.Len())
	if top, ok := intStack.Pop(); ok {
		fmt.Println("popped:", top)
	}

	// Stack of strings
	var strStack Stack[string]
	strStack.Push("eagle")
	strStack.Push("shark")
	if top, ok := strStack.Peek(); ok {
		fmt.Println("peek:", top)
	}

	// Pair with two different types
	p := NewPair("eagle", 150)
	fmt.Printf("pair: first=%s second=%d\n", p.First, p.Second)
}
