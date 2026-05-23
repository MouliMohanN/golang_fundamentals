package main

import "fmt"

func main() {
	// -------------------------------------------------------------------------
	// 1. C-style for loop — init; condition; post
	// -------------------------------------------------------------------------
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Counting down
	for i := 5; i > 0; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Step by 2
	for i := 0; i <= 10; i += 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// -------------------------------------------------------------------------
	// 2. While-style — Go has no `while` keyword, use for with only a condition
	// -------------------------------------------------------------------------
	n := 1
	for n < 64 {
		n *= 2
	}
	fmt.Println("first power of 2 >= 64:", n)

	// -------------------------------------------------------------------------
	// 3. Infinite loop — for without any condition, must break out manually
	// -------------------------------------------------------------------------
	count := 0
	for {
		count++
		if count == 5 {
			break
		}
	}
	fmt.Println("broke out at count:", count)

	// -------------------------------------------------------------------------
	// 4. continue — skip the current iteration, move to the next
	// -------------------------------------------------------------------------
	fmt.Print("even numbers: ")
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue // skip odd numbers
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// -------------------------------------------------------------------------
	// 5. range over slice — yields index and value
	// -------------------------------------------------------------------------
	animals := []string{"cat", "dog", "eagle", "shark"}

	for i, animal := range animals {
		fmt.Printf("index=%d  animal=%s\n", i, animal)
	}

	// Index only
	for i := range animals {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Value only — discard index with _
	for _, animal := range animals {
		fmt.Print(animal, " ")
	}
	fmt.Println()

	// -------------------------------------------------------------------------
	// 6. range over string — yields index (byte position) and rune (character)
	//    Important: index is byte offset, not character number
	// -------------------------------------------------------------------------
	word := "hello,世界"
	for i, ch := range word {
		fmt.Printf("byte index=%d  char=%c  rune=%d\n", i, ch, ch)
	}

	// -------------------------------------------------------------------------
	// 7. range over map — order is random every run
	// -------------------------------------------------------------------------
	legs := map[string]int{"cat": 4, "bird": 2, "spider": 8}
	for animal, count := range legs {
		fmt.Printf("%s has %d legs\n", animal, count)
	}

	// -------------------------------------------------------------------------
	// 8. Nested loops
	// -------------------------------------------------------------------------
	fmt.Println("\n-- multiplication table (3x3) --")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}

	// -------------------------------------------------------------------------
	// 9. Labeled break — break out of an outer loop from inside an inner loop
	// -------------------------------------------------------------------------
	fmt.Println("\n-- labeled break --")
outer:
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i+j == 4 {
				fmt.Printf("breaking outer at i=%d j=%d\n", i, j)
				break outer // exits the outer loop entirely
			}
			fmt.Printf("i=%d j=%d\n", i, j)
		}
	}

	// -------------------------------------------------------------------------
	// 10. Labeled continue — continue the outer loop from inside an inner loop
	// -------------------------------------------------------------------------
	fmt.Println("\n-- labeled continue --")
loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue loop // skip rest of inner loop, continue outer
			}
			fmt.Printf("i=%d j=%d\n", i, j)
		}
	}
}
