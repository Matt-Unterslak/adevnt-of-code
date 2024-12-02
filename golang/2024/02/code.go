package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input.
// 2. with: true (part2), and example input.
// 3. with: false (part1), and user input.
// 4. with: true (part2), and user input.
// the return value of each run is printed to stdout.
func run(part2 bool, input string) any {
	// solve part 2.
	if part2 {
		return solvePart2(input)
	}
	// solve part 1.
	return solvePart1(input)
}

func convertStrToInt(str string) int {
	// string to int.
	convertedInt, err := strconv.Atoi(str)
	if err != nil {
		// ... handle error.
		panic(err)
	}

	return convertedInt
}

func extractInputPart1(input string) [][]int {
	// Split the string into lines.
	lines := strings.Split(input, "\n")

	// Create a 2D matrix (slice of slices of integers).
	matrix := make([][]int, len(lines))
	for lineIndex, line := range lines {
		var rows []int
		row := strings.Fields(line)
		// Convert each line into a slice of int.
		for _, value := range row {
			rows = append(rows, convertStrToInt(value))
		}

		// Add row to the matrix.
		matrix[lineIndex] = rows
	}

	return matrix
}

func isReportMonotonic(index int, value int, previousValues []int) bool {
	var previousValue int
	if index == 0 {
		previousValue = previousValues[index]
	} else {
		previousValue = previousValues[index-1]
	}

	if (value > 0 && previousValue > 0) || (value < 0 && previousValue < 0) {
		return true
	}

	return false
}

func getReportSafety(report []int) int {
	isSafe := true
	var previousCompares []int
	for reportIndex := 0; reportIndex < len(report)-1; reportIndex++ {
		currentValue := report[reportIndex]
		nextValue := report[reportIndex+1]
		currentCompare := currentValue - nextValue
		previousCompares = append(previousCompares, currentCompare)

		isMonotonicCalculation := isReportMonotonic(reportIndex, currentCompare, previousCompares)
		levelGap := int(math.Abs(float64(currentCompare)))

		// fmt.Printf(
		// "Current value: %v, next value: %v, current gap: %v, previous gap: %v\n"
		// , currentValue, nextValue, currentCompare, previousCompare
		// )
		// fmt.Printf("level gap: %v, monotonic: %v\n", levelGap, isMonotonic)
		if isMonotonicCalculation && levelGap > 0 && levelGap <= 3 {
			isSafe = true
		} else {
			isSafe = false

			break
		}
	}

	if isSafe {
		return 1
	}

	return 0
}

func calculateLevelSafety(matrix [][]int) int {
	totalSafe := 0
	for _, row := range matrix {
		reportSafe := getReportSafety(row)
		// fmt.Printf("Report safety: %v\n", reportSafe)
		totalSafe += reportSafe
	}

	return totalSafe
}

func solvePart1(input string) any {
	inputRows := extractInputPart1(input)
	// fmt.Printf("input data: %v\n", inputRows)
	return calculateLevelSafety(inputRows)
}

func removeCurrentValueFromReport(currentIndex int, report []int) []int {
	// Create a new slice by copying.
	newReport := make([]int, 0, len(report)-1)
	newReport = append(newReport, report[:currentIndex]...)
	newReport = append(newReport, report[currentIndex+1:]...)

	// fmt.Printf("Excluding index %d (%d): %v\n", currentIndex, report[currentIndex], newReport)
	return newReport
}

func getReportSafetyWithDampener(report []int) int {
	var safeValues []int

	safeReport := getReportSafety(report)
	if safeReport == 1 {
		return 1
	}
	for c := 0; c < len(report); c++ {
		// Create a new slice excluding the current element.
		removeCurrentLevel := removeCurrentValueFromReport(c, report)
		isSafe := getReportSafety(removeCurrentLevel)
		// fmt.Printf("Reduced report: %v, safety: %v\n", removeCurrentLevel, isSafe)
		safeValues = append(safeValues, isSafe)
	}
	// fmt.Printf("new safe: %v\n", safeValues)

	if slices.Contains(safeValues, 1) {
		// fmt.Printf("unsafe report detected: %v, report is now safe\n", report)
		return 1
	}
	// fmt.Printf("unsafe report detected: %v\n", report)
	return 0
}

func calculateLevelSafetyWithDampener(matrix [][]int) int {
	totalSafe := 0
	for _, row := range matrix {
		reportSafe := getReportSafetyWithDampener(row)
		// fmt.Printf("Report safety: %v\n", reportSafe)
		totalSafe += reportSafe
	}

	return totalSafe
}

func solvePart2(input string) any {
	inputRows := extractInputPart1(input)
	// fmt.Printf("input data: %v\n", inputRows)
	return calculateLevelSafetyWithDampener(inputRows)
}
