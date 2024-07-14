package day1

import (
	"ececillo/advent-of-code-2023/day1"
	"testing"
)

func TestExtractCalibrationValueFromString(t *testing.T) {
	input := "pqr3stu8vwx"
	want := 38
	calibrationValue, err := day1.ExtractCalibrationValueFromString(input)
	if want != calibrationValue || err != nil {
		t.Fatalf(`ExtractCalibrationValueFromString returned %q, %v, want match for %#q, nil`, calibrationValue, err, want)
	}
}

func TestExtractCalibrationWithNoValue(t *testing.T) {
	input := "pqrstuvwx"
	calibrationValue, err := day1.ExtractCalibrationValueFromString(input)
	if calibrationValue != 0 || err != nil {
		t.Fatalf(`ExtractCalibrationValueFromString throw error %q, %v`, calibrationValue, err)
	}
}
