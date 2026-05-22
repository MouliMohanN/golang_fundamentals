package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	// -------------------------------------------------------------------------
	// 1. Print family — Print, Println, Printf
	// -------------------------------------------------------------------------
	fmt.Print("Print: no newline, ")
	fmt.Print("continues on same line\n")

	fmt.Println("Println: auto newline, adds spaces between", "multiple", "args")

	fmt.Printf("Printf: format string, name=%s age=%d\n", "Mouli", 25)

	// -------------------------------------------------------------------------
	// 2. General format verbs
	// -------------------------------------------------------------------------
	fmt.Println("\n-- general verbs --")

	type Point struct{ X, Y int }
	p := Point{3, 7}

	fmt.Printf("%%v  (default)      : %v\n", p)
	fmt.Printf("%%+v (with fields)  : %+v\n", p)
	fmt.Printf("%%#v (Go syntax)    : %#v\n", p)
	fmt.Printf("%%T  (type)         : %T\n", p)
	fmt.Printf("%%%%  (literal %%)  : %%\n")

	// -------------------------------------------------------------------------
	// 3. Boolean
	// -------------------------------------------------------------------------
	fmt.Println("\n-- boolean --")
	fmt.Printf("%%t: %t  %t\n", true, false)

	// -------------------------------------------------------------------------
	// 4. Integer verbs
	// -------------------------------------------------------------------------
	fmt.Println("\n-- integer verbs --")
	n := 255

	fmt.Printf("%%d  (decimal)      : %d\n", n)
	fmt.Printf("%%b  (binary)       : %b\n", n)
	fmt.Printf("%%o  (octal)        : %o\n", n)
	fmt.Printf("%%O  (octal 0o)     : %O\n", n)
	fmt.Printf("%%x  (hex lower)    : %x\n", n)
	fmt.Printf("%%X  (hex upper)    : %X\n", n)
	fmt.Printf("%%c  (char/rune)    : %c\n", 72) // 'H'
	fmt.Printf("%%U  (Unicode)      : %U\n", '世') // U+4E16
	fmt.Printf("%%q  (quoted char)  : %q\n", 65)  // 'A'

	// -------------------------------------------------------------------------
	// 5. Float verbs
	// -------------------------------------------------------------------------
	fmt.Println("\n-- float verbs --")
	f := 123456.789

	fmt.Printf("%%f  (decimal)       : %f\n", f)
	fmt.Printf("%%F  (decimal upper) : %F\n", f)
	fmt.Printf("%%e  (scientific)    : %e\n", f)
	fmt.Printf("%%E  (sci upper)     : %E\n", f)
	fmt.Printf("%%g  (compact)       : %g\n", f)
	fmt.Printf("%%G  (compact upper) : %G\n", f)

	// -------------------------------------------------------------------------
	// 6. String and byte slice verbs
	// -------------------------------------------------------------------------
	fmt.Println("\n-- string/bytes verbs --")
	s := "hello"

	fmt.Printf("%%s  (plain string) : %s\n", s)
	fmt.Printf("%%q  (quoted)       : %q\n", s)
	fmt.Printf("%%x  (hex bytes)    : %x\n", s)
	fmt.Printf("%%X  (hex upper)    : %X\n", s)

	// -------------------------------------------------------------------------
	// 7. Pointer
	// -------------------------------------------------------------------------
	fmt.Println("\n-- pointer --")
	val := 42
	fmt.Printf("%%p  (pointer addr) : %p\n", &val)

	// -------------------------------------------------------------------------
	// 8. Width, precision, and padding
	// -------------------------------------------------------------------------
	fmt.Println("\n-- width & padding --")

	// Width: minimum field width (right-aligned by default)
	fmt.Printf("[%10d]  right-align width 10\n", 42)
	fmt.Printf("[%-10d]  left-align  width 10\n", 42)
	fmt.Printf("[%010d]  zero-pad    width 10\n", 42)

	// Always show sign
	fmt.Printf("[%+d] [%+d]  force sign\n", 42, -42)

	// Space for positive numbers (aligns with negatives)
	fmt.Printf("[% d] [% d]  space for positive\n", 42, -42)

	// Float precision
	fmt.Printf("[%f]      default precision\n", 3.14159)
	fmt.Printf("[%.2f]    2 decimal places\n", 3.14159)
	fmt.Printf("[%10.2f]  width 10, 2 decimals\n", 3.14159)
	fmt.Printf("[%-10.2f] left-align, 2 decimals\n", 3.14159)

	// String width
	fmt.Printf("[%10s]   right-align string\n", "go")
	fmt.Printf("[%-10s]   left-align  string\n", "go")
	fmt.Printf("[%.3s]    truncate to 3 chars\n", "golang")

	// Dynamic width and precision using *
	fmt.Printf("[%*d]     dynamic width=8\n", 8, 42)
	fmt.Printf("[%.*f]    dynamic precision=3\n", 3, 3.14159)

	// -------------------------------------------------------------------------
	// 9. Sprintf — format to a string (does not print)
	// -------------------------------------------------------------------------
	fmt.Println("\n-- Sprintf --")

	greeting := fmt.Sprintf("Hello, %s! You are %d years old.", "Mouli", 25)
	hex := fmt.Sprintf("0x%X", 255)
	padded := fmt.Sprintf("|%10s|%-10s|", "right", "left")

	fmt.Println(greeting)
	fmt.Println(hex)
	fmt.Println(padded)

	// Sprintf is useful for building strings before logging or returning them
	label := fmt.Sprintf("user_%05d", 42) // "user_00042"
	fmt.Println(label)

	// Sprint / Sprintln — like Print/Println but return a string
	s1 := fmt.Sprint("a", "b", "c")        // "abc"
	s2 := fmt.Sprintln("a", "b", "c")      // "a b c\n"
	fmt.Printf("Sprint=%q Sprintln=%q\n", s1, s2)

	// -------------------------------------------------------------------------
	// 10. Fprintf — format to any io.Writer
	// -------------------------------------------------------------------------
	fmt.Println("\n-- Fprintf --")

	// Write to stdout explicitly
	fmt.Fprintf(os.Stdout, "Fprintf to stdout: %s\n", "works")

	// Write to stderr (useful for error messages)
	fmt.Fprintf(os.Stderr, "Fprintf to stderr: %s\n", "error output")

	// Write to a bytes.Buffer (in-memory writer)
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "buffered: %d + %d = %d", 1, 2, 3)
	fmt.Println(buf.String())

	// Write to a strings.Builder (efficient string building)
	var sb strings.Builder
	for i := 1; i <= 5; i++ {
		fmt.Fprintf(&sb, "item%d ", i)
	}
	fmt.Println(sb.String())

	// Fprint / Fprintln variants
	fmt.Fprint(os.Stdout, "Fprint: no format\n")
	fmt.Fprintln(os.Stdout, "Fprintln:", "auto", "spaced")

	// -------------------------------------------------------------------------
	// 11. Errorf — create a formatted error
	// -------------------------------------------------------------------------
	fmt.Println("\n-- Errorf --")

	err := fmt.Errorf("user %d not found in region %q", 42, "us-east")
	fmt.Println(err)
	fmt.Printf("type: %T\n", err)

	// %w wraps an error for errors.Is / errors.As unwrapping
	baseErr := fmt.Errorf("connection timeout")
	wrapped := fmt.Errorf("fetchUser failed: %w", baseErr)
	fmt.Println(wrapped)

	// -------------------------------------------------------------------------
	// 12. Scanf / Sscanf — scan input
	// -------------------------------------------------------------------------
	fmt.Println("\n-- Sscanf (scan from string) --")

	var parsedName string
	var parsedAge int

	// Sscanf parses from a string (safe alternative to Scanf for demos)
	input := "Alice 30"
	n, err2 := fmt.Sscanf(input, "%s %d", &parsedName, &parsedAge)
	fmt.Printf("scanned %d values: name=%s age=%d err=%v\n", n, parsedName, parsedAge, err2)

	// Sscan (no format string, whitespace delimited)
	var x, y int
	fmt.Sscan("10 20", &x, &y)
	fmt.Printf("Sscan: x=%d y=%d\n", x, y)

	// NOTE: fmt.Scanf reads from stdin — good for interactive programs:
	// fmt.Scanf("%s %d", &name, &age)
}
