// Power prints a decimal representation of negative powers of 2.
package main

import "fmt"

const lines = 10

func main() {
	Power(lines)
}

// Power prints n lines of the negative powers of 2.
func Power(n int) {
	digits := make([]int, n+1)

	digits[0] = 5
	fmt.Println(".5")
	for line := 1; line < n; line++ {
		fmt.Print(".")
		rem := 0
		for i := 0; i <= line; i++ {
			rem = 10*rem + digits[i]
			digits[i] = rem / 2
			rem -= 2 * digits[i]
			fmt.Print(string(digits[i] + '0'))
		}
		fmt.Println()
	}
}
