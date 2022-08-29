package main

import (
	"fmt"
	"log"
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

// fromAnyBasetoAnyBase converts a number in a specified base to a desired base
func fromAnyBasetoAnyBase(base int8, number int64, desiredBase int) (int, error) {

	if base != 10 {
		numBase10, err := toBase10(base, number)
		if err != nil {
			log.Fatal(err)
		}
		number = numBase10
	}

	result := 0
	counter := 1
	remainder := 0
	for number != 0 {
		remainder = int(number) % int(desiredBase)
		number = number / int64(desiredBase)
		result += remainder * counter
		counter *= 10
	}
	return result, nil
}
