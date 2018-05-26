// Power prints a decimal representation of negative powers of 2.
package main

import "fmt"

const lines = 10

func main() {
	Power(lines)
}

// Power prints n lines of the negative powers of 2.
func Power(n int) {
	if n < 1 {
		return
	}
	// Line 0 (printing ".5") has 1 digit at array index 0, so the last line
	// (line n - 1) requires n digits.
	digits := make([]int, n)

	digits[0] = 5
	fmt.Println(".5")
	for line := 1; line < n; line++ {
		fmt.Print(".")
		rem := 0
		// There are line + 1 digits to print.
		for d := 0; d <= line; d++ {
			rem = 10*rem + digits[d]
			digits[d] = rem / 2
			rem -= 2 * digits[d]
			fmt.Print(string(digits[d] + '0'))
		}
		fmt.Println()
	}
}
