package intermediate

import "fmt"

func formatting() {
	// will remove underscores btw
	i := 222_222.89
	a := "skjdskdj"
	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	// two %s here are like escape so you get float + %
	fmt.Printf("%v%%\n", i)

	fmt.Printf("%v\n", a)
	// gives you string in double quotes bc that is how its stored
	fmt.Printf("%#v\n", a)
	fmt.Printf("%T\n", a)

	n := 18

	fmt.Printf("// --- Integer Formatting Verbs\n\n")

	// Binary
	fmt.Printf("%%b   → %b\n", n) // Base 2

	// Decimal
	fmt.Printf("%%d   → %d\n", n)  // Base 10
	fmt.Printf("%%+d  → %+d\n", n) // Always show sign

	// Octal
	fmt.Printf("%%o   → %o\n", n)  // Base 8
	fmt.Printf("%%#o  → %#o\n", n) // Base 8, with leading 0o

	// Hexadecimal
	fmt.Printf("%%x   → %x\n", n)  // Base 16, lowercase
	fmt.Printf("%%X   → %X\n", n)  // Base 16, uppercase
	fmt.Printf("%%#x  → %#x\n", n) // Base 16, with leading 0x
	fmt.Printf("%%#X  → %#X\n", n) // Base 16, with leading 0X

	// Width / padding
	fmt.Printf("%%4d  → %4d\n", n)  // Pad with spaces (width 4, right-justified)
	fmt.Printf("%%-4d → %-4d\n", n) // Pad with spaces (width 4, left-justified)
	fmt.Printf("%%04d → %04d\n", n) // Pad with zeroes (width 4)

	txt := "World"

	// --- String Formatting Verbs

	// %s   → plain string
	fmt.Printf("%%s   → %s\n", txt)

	// %q   → double-quoted string
	fmt.Printf("%%q   → %q\n", txt)

	// %8s  → width 8, right-justified
	fmt.Printf("%%8s  → %8s\n", txt)

	// %-8s → width 8, left-justified
	fmt.Printf("%%-8s → %-8s\n", txt)

	// %x   → hex dump (no spaces)
	fmt.Printf("%%x   → %x\n", txt)

	// % x  → hex dump (with spaces)
	fmt.Printf("%% x  → % x\n", txt)

	flt := 9.18

	// --- Float Formatting Verbs

	// %e   → scientific notation with 'e'
	fmt.Printf("%%e     → %e\n", flt)

	// %f   → decimal point, no exponent
	fmt.Printf("%%f     → %f\n", flt)

	// %.2f → default width, precision 2
	fmt.Printf("%%.2f   → %.2f\n", flt)

	// %6.2f → width 6, precision 2
	fmt.Printf("%%6.2f  → %6.2f\n", flt)

	// %g   → exponent as needed, only necessary digits
	fmt.Printf("%%g     → %g\n", flt)

}
