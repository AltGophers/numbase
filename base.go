package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

// convertToBase converts a number from its current base to a desired base between 2 - 16
func convertToBase(base int8, num string, desiredBase int8) (string, error) {
	var result string
	var number int64
	var err error

	if base < 10 {
		if number, err = strconv.ParseInt(num, 10, 64); err != nil {
			return "", ErrInvalidBase(base)
		}
		numBase10, err := convertNumToBase10(base, number)
		if err != nil {
			return "", fmt.Errorf("convertToBase10 error: %v", err)
		}
		number = numBase10

	} else if base > 10 {
		numBase10, err := convertBasesAboveTenToBase10(num, base)
		if err != nil {
			return "", fmt.Errorf("convertToBase10 error: %v", err)
		}
		number = numBase10
	} else {
		number, err = strconv.ParseInt(num, 10, 64)
		if err != nil {
			return "", ErrInvalidBase(base)
		}
	}

	switch {
	case desiredBase > 10:
		result = convertToBaseGreaterThan10(number, desiredBase)
	case desiredBase < 10:
		numResult := convertToBaseLessThan10(number, desiredBase)
		result = fmt.Sprint(numResult)
	default:
		result = fmt.Sprint(number)
	}

	return result, nil
}

// convertNumToBase10 converts a number in the specified base to base 10.
func convertNumToBase10(base int8, number int64) (int64, error) {
	index, answer := 0, 0.0

	for number != 0 {
		currentNum := number % 10
		if int8(currentNum) >= base {
			return 0, ErrInvalidBase(base)
		}

		number /= 10
		answer += float64(currentNum) * math.Pow(float64(base), float64(index))
		index++
	}
	return int64(answer), nil
}

// convertToBaseGreaterThan10 converts a number in base 10 to a number greater than base 10
func convertToBaseGreaterThan10(number int64, desiredBase int8) string {
	numResult := make([]byte, 0, 2)
	m := map[int]byte{
		10: 'A',
		11: 'B',
		12: 'C',
		13: 'D',
		14: 'E',
		15: 'F',
	}

	remainder := 0
	for number != 0 {
		remainder = int(number % int64(desiredBase))
		if remainder >= 10 {
			numResult = append(numResult, m[remainder])
		} else {
			numResult = append(numResult, []byte(strconv.Itoa(remainder))...)
		}
		number = number / int64(desiredBase)
	}

	var numResultRvsed []byte
	for i := len(numResult) - 1; i >= 0; i-- {
		numResultRvsed = append(numResultRvsed, numResult[i])
	}

	return string(numResultRvsed)
}

// convertToBaseLessThan10 converts a number in base 10 to a number less than base 10
func convertToBaseLessThan10(number int64, desiredBase int8) int64 {
	numResult := int64(0)
	counter := int64(1)
	remainder := 0
	for number != 0 {
		remainder = int(number % int64(desiredBase))
		number = number / int64(desiredBase)
		numResult += int64(remainder) * counter
		counter *= 10
	}
	return numResult
}

// convertBasesAboveTenToBase10 convert a number from its current base > 10 to base 10
func convertBasesAboveTenToBase10(num string, currentBase int8) (int64, error) {
	decNum := int64(0)
	i := 0

	numLen := len(num) - 1
	for numLen >= 0 {
		rem := num[numLen]
		var remValue rune

		if rem >= '0' && rem <= '9' {
			remInt, _ := strconv.Atoi(string(rem))
			remIntByte := make([]byte, 4)
			binary.LittleEndian.PutUint32(remIntByte, uint32(remInt))
			remValue, _ = utf8.DecodeRune(remIntByte)
		} else if rem >= 'A' && rem <= 'F' {
			if !isValidBaseDigit(rem, currentBase) {
				return 0, ErrInvalidBase(currentBase)
			}
			remRune, _ := utf8.DecodeRuneInString(string(rem))
			remValue = remRune - 55 // Set remValue to the hexadecimal value of remRune
		} else if rem >= 'a' && rem <= 'f' {
			if !isValidBaseDigit(rem, currentBase) {
				return 0, ErrInvalidBase(currentBase)
			}
			remRune, _ := utf8.DecodeRuneInString(string(rem))
			remValue = remRune - 87 // Set remValue to the hexadecimal value of remRune
		} else {
			return 0, ErrInvalidBase(currentBase)
		}

		res := int64(math.Pow(float64(currentBase), float64(i)))
		decNum = decNum + (int64(remValue) * res)
		numLen--
		i++
	}

	return decNum, nil
}

func isValidBaseDigit(digit byte, base int8) bool {
	bases := []byte{'A', 'B', 'C', 'D', 'E', 'F'}
	sliceDigit := []byte{digit}

	return bytes.Contains(bases[:base-10], bytes.ToUpper(sliceDigit))
}

func ErrInvalidBase(base int8) error {
	return fmt.Errorf("conversion error: Specified number is not in base %d", base)
}
