package math_01

import "fmt"

func Sum(i, j int) int {
	sum := i + j

	if sum > 0 {
		fmt.Printf("sum is positive: %d\n", sum)
	} else if sum < 0 {
		fmt.Printf("sum is negative: %d\n", sum)
	}

	return sum
}
