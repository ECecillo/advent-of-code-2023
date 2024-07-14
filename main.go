package main

import "ececillo/advent-of-code-2023/day1"

// Uncomment to run other days
// import "ececillo/advent-of-code-2023/day2"
// import "ececillo/advent-of-code-2023/day3"

func main() {
	result := day1.ReadCalibrationFileAndReturnSum("data/day1.txt")
	if result != 0 {
		println(result)
	} else {
		println("No result")
	}
}
