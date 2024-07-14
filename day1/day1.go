package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func ExtractCalibrationValueFromString(calibrationLine string) (int, error) {
	re := regexp.MustCompile(`[0-9]+`)
	arrayOfStringNumbers := re.FindAllString(calibrationLine, -1)

	if len(arrayOfStringNumbers) == 0 {
		return 0, nil
	}

	sum, err := sumDigits(arrayOfStringNumbers)
	if err != nil {
		return -1, err
	}

	fmt.Println("Sum", sum)

	return sum, nil
}

func sumDigits(digits []string) (int, error) {
	sum := 0
	digit := ""
	if len(digits) == 1 {
		firstDigit := digits[0]
		if len(firstDigit) == 2 {
			digit = string(firstDigit)
		} else {
			digit = string(firstDigit[0]) + string(firstDigit[0])
		}
	} else {
		firstDigit := digits[0]
		secondDigit := digits[len(digits)-1]
		digit = string(firstDigit[0]) + string(secondDigit[len(secondDigit)-1])
	}

	digitInt, err := strconv.Atoi(digit)
	if err != nil {
		return -1, fmt.Errorf("error when converting string to int: %v", err)
	}
	sum += digitInt

	return sum, nil
}

func ReadCalibrationFileAndReturnSum(calibrationFilePath string) int {
	file, err := os.Open(calibrationFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumOfAllCalibrationValue := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		calibrationValue, err := ExtractCalibrationValueFromString(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sumOfAllCalibrationValue += calibrationValue
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sumOfAllCalibrationValue
}
