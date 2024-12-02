package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"math"
	"strconv"
	"strings"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return solvePart2(input)
	}
	// solve part 1 here
	return solvePart1(input)
}

func convertStrToInt(str string) int {
	// string to int
	i, err := strconv.Atoi(str)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}

func extractInputPart1(input string) [][]int {
	// Split the string into lines
	lines := strings.Split(input, "\n")

	// Create a 2D matrix (slice of slices of runes)
	matrix := make([][]int, len(lines))
	for i, line := range lines {
		var rows []int
		row := strings.Fields(line)
		// Convert each line into a slice of int
		for _, value := range row {
			rows = append(rows, convertStrToInt(value))
		}

		// Add row to the matrix
		matrix[i] = rows
	}
	return matrix
}

func isReportMonotonic(index int, value int, previousValues []int) bool {
	previousValue := 0
	if index == 0 {
		previousValue = previousValues[index]
	} else {
		previousValue = previousValues[index-1]
	}

	if (value > 0 && previousValue > 0) || (value < 0 && previousValue < 0) {
		return true
	} else {
		return false
	}
}

func getReportSafety(report []int) int {
	isSafe := true
	var previousCompares []int
	previousCompare := 0
	isMonotonic := true
	for c := 0; c < len(report)-1; c++ {
		currentValue := report[c]
		nextValue := report[c+1]
		currentCompare := currentValue - nextValue
		previousCompares = append(previousCompares, currentCompare)

		if c == 0 {
			previousCompare = previousCompares[c]
		} else {
			previousCompare = previousCompares[c-1]
		}

		if (currentCompare > 0 && previousCompare > 0) || (currentCompare < 0 && previousCompare < 0) {
			isMonotonic = true
		} else {
			isMonotonic = false
		}

		levelGap := int(math.Abs(float64(currentCompare)))

		//fmt.Printf("Current value: %v, next value: %v, current gap: %v, previous gap: %v\n", currentValue, nextValue, currentCompare, previousCompare)
		//fmt.Printf("level gap: %v, monotonic: %v\n", levelGap, isMonotonic)
		if isMonotonic && levelGap > 0 && levelGap <= 3 {
			isSafe = true
		} else {
			isSafe = false
			break
		}
	}

	if isSafe {
		return 1
	} else {
		return 0
	}
}

func calculateLevelSafety(matrix [][]int) int {
	totalSafe := 0
	for _, row := range matrix {
		reportSafe := getReportSafety(row)
		//fmt.Printf("Report safety: %v\n", reportSafe)
		totalSafe += reportSafe
	}
	return totalSafe
}

func solvePart1(input string) any {
	inputRows := extractInputPart1(input)
	//fmt.Printf("input data: %v\n", inputRows)
	return calculateLevelSafety(inputRows)
}

func solvePart2(input string) any {
	inputRows := extractInputPart1(input)
	//fmt.Printf("input data: %v\n", inputRows)
	return calculateLevelSafety(inputRows)
}
