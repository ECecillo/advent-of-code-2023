package day1

import (
	"ececillo/advent-of-code-2023/day1"
	"testing"
)

// ExtractCalibrationValueFromString calls day1.ExtractCalibrationValue
// with a string as input, returning an integer value.
func TestExtractCalibrationValueFromString(t *testing.T) {
	input := "pqr3stu8vwx"
	want := 38
	calibrationValue := day1.ExtractCalibrationValue(input)
	if want != calibrationValue {
		t.Fatalf(`ExtractCalibrationValue(input) = %q, %v, nil`, calibrationValue, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
// func TestExtractCalibrationValueNoDigit(t *testing.T) {
// 	msg, err := Hello("")
// 	if msg != "" || err == nil {
// 		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
// 	}
// }
