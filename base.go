package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

// convertNumToBase10 converts a number in the specified base to base 10.
func convertNumToBase10(base int8, number int64) (int64, error) {
	index, answer := 0, 0.0
	originalNum := number
	for number != 0 {
		currentNum := number % 10
		if int8(currentNum) >= base {
			return 0, fmt.Errorf("Error: %d not in specified base: %d", originalNum, base)
		}

		number /= 10
		answer += float64(currentNum) * math.Pow(float64(base), float64(index))
		index++
	}
	return int64(answer), nil
}

// convertToBase converts a number from its current base to a desired base.
func convertToBase(base int8, number int64, desiredBase int) (int, error) {
	if base != 10 {
		numBase10, err := convertNumToBase10(base, number)
		if err != nil {
			return 0, fmt.Errorf("convertNumToBase10 error: %v", err)
		}
		number = numBase10
	}

	result := 0
	counter := 1
	remainder := 0
	for number != 0 {
		remainder = int(number) % desiredBase
		number = number / int64(desiredBase)
		result += remainder * counter
		counter *= 10
	}
	return result, nil
}

// hexToAny convert hexadecimals to any desired base between base 2-10
func hexToAny(hexNum string, desiredBase int) (int, error) {
	var result int
	var err error
	Chk := 0
	decNum := 0
	i := 0

	hexNumLen := len(hexNum)
	hexNumLen = hexNumLen - 1

	for hexNumLen >= 0 {
		rem := hexNum[hexNumLen]
		var strNew string
		var remValue rune

		if rem >= '0' && rem <= '9' {
			remStr := string(rem)
			remInt, _ := strconv.Atoi(remStr)
			remIntByte := make([]byte, 4)
			binary.LittleEndian.PutUint32(remIntByte, uint32(remInt))
			remRune, _ := utf8.DecodeRune(remIntByte)
			remValue = remRune * 1
		} else if rem >= 'A' && rem <= 'F' {
			strNew = string(rem)
			remRune, _ := utf8.DecodeRuneInString(strNew)
			remValue = remRune - 55
		} else if rem >= 'a' && rem <= 'f' {
			strNew = string(rem)
			remRune, _ := utf8.DecodeRuneInString(strNew)
			remValue = remRune - 87
		} else {
			Chk = 1
			break
		}

		iFloat := float64(i)
		res := int(math.Pow(16, iFloat))
		decNum = decNum + (int(remValue) * res)
		hexNumLen = hexNumLen - 1
		i = i + 1
	}

	if Chk == 0 {
		result, err = convertToBase(10, int64(decNum), desiredBase)
		if err != nil {
			log.Fatalf("Can't convert resulted decimal number from the inputted hexadecimal to your desired base ")
		}
	}

	return result, nil
}