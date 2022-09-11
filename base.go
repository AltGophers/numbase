package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"strconv"
	"unicode/utf8"
)

// toBase10 converts a number in a specified base to decimal (base 10)
func toBase10(base int8, number int64) (int64, error) {
	index, answer := 0, 0.0
	inputNum := number
	for number != 0 {
		currentNum := number % 10
		if int8(currentNum) >= base {
			return 0, fmt.Errorf("error: %d not in specified base: %d", inputNum, base)
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

// to convert hexadecimals to any desired base between base 2-10

func hextoAny(hexdecNum string, desiredBase int) (int, error) {

	var result int
	var err error
	Chk := 0
	decnum := 0
	i := 0

	hexdecNumLen := len(hexdecNum)
	hexdecNumLen = hexdecNumLen - 1

	for hexdecNumLen >= 0 {
		
		rem := hexdecNum[hexdecNumLen]
		var strnew string
		var remValue rune

		if rem >= '0' && rem <= '9' {
			remStr := string(rem)
			remInt, _ := strconv.Atoi(remStr)
			remIntByte := make([]byte, 4)
			binary.LittleEndian.PutUint32(remIntByte, uint32(remInt))
			remRune, _ := utf8.DecodeRune(remIntByte)
			remValue = remRune * 1

		} else if rem >= 'A' && rem <= 'F' {
			strnew = string(rem)
			remRune, _ := utf8.DecodeRuneInString(strnew)
			remValue = remRune - 55

		} else if rem >= 'a' && rem <= 'f' {
			strnew = string(rem)
			remRune, _ := utf8.DecodeRuneInString(strnew)
			remValue = remRune - 87

		} else {
			Chk = 1
			break
		}
		iFloat := float64(i)
		res := int(math.Pow(16, iFloat))
		decnum = decnum + (int(remValue) * res)
		hexdecNumLen = hexdecNumLen - 1
		i = i + 1
	}

	if Chk == 0 {
		result, err = fromAnyBasetoAnyBase(10, int64(decnum), desiredBase)
		if err != nil {
			log.Fatalf("Can't convert resulted decimal number from the inputted hexadecimal to your desired base ")
		}
	}

	return result, nil

}


