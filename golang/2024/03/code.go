package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"regexp"
	"strconv"
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
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return solvePart2(input)
	}
	// solve part 1 here
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

func solvePart1(input string) int {
	// Define the regex pattern to match `mul(...)` and capture the numbers inside.
	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)

	// Find all matches with subgroups.
	matches := re.FindAllStringSubmatch(input, -1)

	// Loop through the matches and extract the groups.
	totalMultiplications := 0
	for _, match := range matches {
		// match[1] is the first number, match[2] is the second number.
		if len(match) > 2 {
			firstNumber := convertStrToInt(match[1])
			secondNumber := convertStrToInt(match[2])
			multiplication := firstNumber * secondNumber
			// fmt.Printf("Multiplication: %v * %v = %v\n", firstNumber, secondNumber, multiplication)
			totalMultiplications += multiplication
			// fmt.Printf("Total Multiplications: %v\n", totalMultiplications)
		}
	}

	// for _, match := range patternMatches {
	//	fmt.Printf("Match Stored: %v\n", match)
	// }

	return totalMultiplications
}

func extractInstructions(input string) [][]int {
	// Define the regex pattern to match "do()" and "don't()" exactly.
	instructionsPattern := `do\(\)|don't\(\)`
	instructionsRe := regexp.MustCompile(instructionsPattern)

	// Find all matches with indices.
	instructionsMatches := instructionsRe.FindAllStringIndex(input, -1)

	// Iterate through matches.
	var doCommandsActive [][]int
	var doActive []int
	doActive = append(doActive, 0)
	isDoCommand := true
	for _, match := range instructionsMatches {
		// match[0] is the start index, match[1] is the end index.
		fullMatch := input[match[0]:match[1]]
		// fmt.Printf("Full match: %s, Start index: %d, End index: %d\n", fullMatch, match[0], match[1])
		if fullMatch == "do()" && !isDoCommand {
			doActive = append(doActive, match[0])
			isDoCommand = true
		} else if fullMatch == "don't()" && isDoCommand {
			if isDoCommand {
				var tmpSlice []int
				tmpSlice = append(tmpSlice, doActive[len(doActive)-1], match[0])
				doCommandsActive = append(doCommandsActive, tmpSlice)
			}
			isDoCommand = false
		}
	}
	doCommandsActive = append(doCommandsActive, []int{doActive[len(doActive)-1], len(input)})
	// fmt.Printf("Do active: %v\n", doCommandsActive)
	return doCommandsActive
}

func solvePart2(input string) int {
	doCommandsActive := extractInstructions(input)

	// Define the regex pattern to match `mul(...)` and capture the numbers inside.
	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)

	// Find all matches with indices.
	matches := re.FindAllStringSubmatchIndex(input, -1)

	// Iterate through matches.
	totalMultiplications := 0
	for _, match := range matches {
		// match[0] and match[1] are the start and end indices of the entire match.
		// match[2] and match[3] are the start and end indices of the first capturing group.
		// match[4] and match[5] are the start and end indices of the second capturing group.

		// fullMatch := input[match[0]:match[1]]
		firstNumber := input[match[2]:match[3]]
		secondNumber := input[match[4]:match[5]]

		for _, doCommand := range doCommandsActive {
			if match[0] >= doCommand[0] && match[1] <= doCommand[1] {
				// fmt.Printf("Full match: %s, First number: %s, Second number: %s, Start index: %d, End index: %d\n",
				//	fullMatch, firstNumber, secondNumber, match[0], match[1])
				multiplication := convertStrToInt(firstNumber) * convertStrToInt(secondNumber)
				// fmt.Printf("Multiplication: %v * %v = %v\n", firstNumber, secondNumber, multiplication)
				totalMultiplications += multiplication
			}
		}
	}

	return totalMultiplications
}
