package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	inputFilePath = "resources/input.txt"
)

var digitStrings = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

type calibrationValueParser func(line string) (int, error)

func getCalibrationSum(parser calibrationValueParser) (int, error) {
	f, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var sum int
	for scanner.Scan() {
		cv, err := parser(scanner.Text())
		if err != nil {
			return 0, err
		}
		sum += cv
	}

	return sum, nil
}

func getCalibrationValuePart1(line string) (int, error) {
	var (
		left, right byte
		i           = 0
		j           = len(line) - 1
	)

	for i < len(line) && j >= 0 {
		if left == 0 && isAlphaNumeric(line[i]) {
			left = line[i]
		}
		if right == 0 && isAlphaNumeric(line[j]) {
			right = line[j]
		}
		if left != 0 && right != 0 {
			break
		}
		i++
		j--
	}

	cv, err := strconv.Atoi(string([]byte{left, right}))
	if err != nil {
		return 0, err
	}

	return cv, nil
}

func getCalibrationValuePart2(line string) (int, error) {
	bytes := []byte{getDigitFromLeftSide(line), getDigitFromRightSide(line)}
	cv, err := strconv.Atoi(string(bytes))
	if err != nil {
		return 0, err
	}
	return cv, nil
}

func getDigitFromLeftSide(line string) byte {
	for startIdx := range line {
		if isAlphaNumeric(line[startIdx]) {
			return line[startIdx]
		} else {
			for _, str := range digitStrings {
				if strings.Index(line, str) == startIdx {
					return longFormToByte(str)
				}
			}
		}
	}
	return 0
}

func getDigitFromRightSide(line string) byte {
	var reverseLine string
	for idx := range line {
		reverseIdx := len(line) - 1 - idx
		if isAlphaNumeric(line[reverseIdx]) {
			return line[reverseIdx]
		} else {
			reverseLine = line[reverseIdx:]
			for _, str := range digitStrings {
				if strings.Index(reverseLine, str) != -1 {
					return longFormToByte(str)
				}
			}
		}
	}
	return 0
}

func longFormToByte(lf string) byte {
	switch lf {
	case "one":
		return byte('1')
	case "two":
		return byte('2')
	case "three":
		return byte('3')
	case "four":
		return byte('4')
	case "five":
		return byte('5')
	case "six":
		return byte('6')
	case "seven":
		return byte('7')
	case "eight":
		return byte('8')
	case "nine":
		return byte('9')
	default:
		return 0
	}
}

func isAlphaNumeric(b byte) bool {
	r := rune(b)
	return r >= '0' && r <= '9'
}

func main() {
	p1Sum, err := getCalibrationSum(getCalibrationValuePart1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[PART 1] Calibration value sum is: %d\n", p1Sum)

	p2Sum, err := getCalibrationSum(getCalibrationValuePart2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[PART 2] Calibration value sum is: %d\n", p2Sum)
}
