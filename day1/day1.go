package day1

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func ExtractDigitGivenRegexExpr(input string, regex string) string {
	re := regexp.MustCompile(regex)
	stringDigit := re.FindString(input)
	return stringDigit
}

func combineStringDigit(firstStringDigit string, secondStringDigit string) (string, error) {
	stringDigit := ""
	if len(firstStringDigit) == 0 && len(secondStringDigit) == 0 {
		return "", errors.New("no digit found")
	} else if len(firstStringDigit) == 0 && len(secondStringDigit) != 0 {
		stringDigit = secondStringDigit + secondStringDigit
	} else if len(firstStringDigit) != 0 && len(secondStringDigit) == 0 {
		stringDigit = firstStringDigit + firstStringDigit
	} else {
		stringDigit = firstStringDigit + secondStringDigit
	}

	return stringDigit, nil
}

func ExtractCalibrationValue(input string) int {
	firstStringDigit := ExtractDigitGivenRegexExpr(input, "(^|\\s)([0-9]+)($|\\s)")
	secondStringDigit := ExtractDigitGivenRegexExpr(input, "((\\d+)(?!.*\\d))")

	stringDigit, err := combineStringDigit(firstStringDigit, secondStringDigit)
	if err != nil {
		log.Fatal(err)
	}

	digit, err := strconv.Atoi(stringDigit)
	// Exit on error
	if err != nil {
		fmt.Println("Error when converting string to int")
		log.Fatal(err)
	}

	return digit
}

// func ReadCalibrationFile() []int {
// }

func PrintHello() {
	fmt.Println("Hello, World!")
}
