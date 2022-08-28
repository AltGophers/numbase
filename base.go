package main

import (
	"fmt"
	"math"
)

// toBase10 converts a number in a specified base to decimal (base 10)
func toBase10(base int8, number int64) (int64, error) {
	index, answer := 0, 0.0
	inputNum := number
	for number != 0 {
		currentNum := number % 10
		if int8(currentNum) >= base {
			return 0, fmt.Errorf("Error: %d not in specified base: %d", inputNum, base)
		}

		number /= 10
		answer += float64(currentNum) * math.Pow(float64(base), float64(index))
		index++
	}
	return int64(answer), nil
}
