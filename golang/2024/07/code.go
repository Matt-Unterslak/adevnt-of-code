package main

import (
	"encoding/json"
	"fmt"
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
	// Trim leading and trailing whitespace.
	trimmedStr := strings.TrimSpace(str)

	// string to int.
	convertedInt, err := strconv.Atoi(trimmedStr)
	if err != nil {
		// ... handle error.
		panic(err)
	}

	return convertedInt
}

func convertStrToInt64(str string) int64 {
	// Trim leading and trailing whitespace.
	trimmedStr := strings.TrimSpace(str)

	// string to int64.
	convertedInt, err := strconv.ParseInt(trimmedStr, 10, 64)
	if err != nil {
		// ... handle error.
		panic(err)
	}

	return convertedInt
}

func convertStrToUInt64(str string) uint64 {
	// Trim leading and trailing whitespace.
	trimmedStr := strings.TrimSpace(str)

	// string to uint64.
	convertedInt, err := strconv.ParseUint(trimmedStr, 10, 64)
	if err != nil {
		// ... handle error.
		panic(err)
	}

	return convertedInt
}

func readEquations(input string) [][]int {
	var equations [][]int
	for _, line := range strings.Split(input, "\n") {
		//fmt.Printf("line: %v\n", line)

		var equation []int
		equationSplit := strings.Split(line, ":")
		//fmt.Printf("Equation Split 0: %v\n", equationSplit[0])

		equation = append(equation, convertStrToInt(equationSplit[0]))
		equationValues := strings.Split(equationSplit[1], " ")
		//fmt.Printf("Equation Values: %v\n", equationValues)
		for _, equationValue := range equationValues {
			//fmt.Printf("Equation Value: %v\n", equationValue)
			if equationValue != "" {
				equation = append(equation, convertStrToInt(equationValue))
			}
		}
		//fmt.Printf("equation: %v\n", equation)

		equations = append(equations, equation)
	}
	//fmt.Printf("equations: %v\n", equations)
	return equations
}

// Helper function to compute all possible combinations
func findOperations(numbers []int, target int, currentResult int, currentExpression string, results *[]string) {
	// Base case: if no more numbers are left
	if len(numbers) == 0 {
		if currentResult == target {
			*results = append(*results, currentExpression)
		}
		return
	}

	// Take the first number from the remaining list
	next := numbers[0]
	remaining := numbers[1:]
	//fmt.Printf("next: %v, remaining: %v, current result: %v\n", next, remaining, currentResult)

	// Try addition
	findOperations(remaining, target, currentResult+next, fmt.Sprintf("%s + %d", currentExpression, next), results)

	// Try multiplication
	findOperations(remaining, target, currentResult*next, fmt.Sprintf("%s * %d", currentExpression, next), results)
}

// Helper function to compute all possible combinations
func findOperationsAll(numbers []int, target int, results *[]string) {
	// Base case: when only one number is left, check if it matches the target
	if len(numbers) == 1 {
		if numbers[0] == target {
			*results = append(*results, fmt.Sprintf("%d", numbers[0]))
		}
		return
	}

	// Try all possible splits
	for i := 1; i < len(numbers); i++ {
		// Left and right parts
		left := numbers[:i]
		right := numbers[i:]

		// Compute possible results for the left part
		leftResults := []int{}
		findResults(left, &leftResults)

		// Compute possible results for the right part
		rightResults := []int{}
		findResults(right, &rightResults)

		// Combine results from left and right parts with both + and *
		for _, leftResult := range leftResults {
			for _, rightResult := range rightResults {
				// Try addition
				if leftResult+rightResult == target {
					*results = append(*results, fmt.Sprintf("%v + %v", left, right))
				}
				// Try multiplication
				if leftResult*rightResult == target {
					*results = append(*results, fmt.Sprintf("%v * %v", left, right))
				}
			}
		}
	}
}

// Helper function to compute results for a part of the list
func findResults(numbers []int, results *[]int) {
	if len(numbers) == 1 {
		*results = append(*results, numbers[0])
		return
	}
	for i := 1; i < len(numbers); i++ {
		left := numbers[:i]
		right := numbers[i:]
		// Get results for the left and right parts recursively
		leftResults := []int{}
		findResults(left, &leftResults)
		rightResults := []int{}
		findResults(right, &rightResults)

		// Combine the results with + and *
		for _, leftResult := range leftResults {
			for _, rightResult := range rightResults {
				*results = append(*results, leftResult+rightResult)
				*results = append(*results, leftResult*rightResult)
			}
		}
	}
}

// Recursive function to compute all possible results
func findCombinations(numbers []int, target int, currentResult int, currentExpression string, results *[]string) {
	// Base case: if no more numbers are left, check if the result matches the target
	if len(numbers) == 0 {
		if currentResult == target {
			*results = append(*results, currentExpression)
		}
		return
	}

	// Take the first number from the remaining list
	next := numbers[0]
	remaining := numbers[1:]

	// Recursively try addition
	findCombinations(remaining, target, currentResult+next, currentExpression+" + "+strconv.Itoa(next), results)

	// Recursively try multiplication
	findCombinations(remaining, target, currentResult*next, currentExpression+" * "+strconv.Itoa(next), results)
}

func findTrueEquations(equations [][]int) map[int][]string {
	possibleCombinations := make(map[int][]string)
	for _, equation := range equations {
		// Example inputs
		target := equation[0]
		numbers := equation[1:]

		// Initial setup: Start with the first number as the initial result
		initialNumber := numbers[0]
		initialExpression := strconv.Itoa(initialNumber)
		results := []string{}
		// Recursively compute combinations
		findCombinations(numbers[1:], target, initialNumber, initialExpression, &results)

		if len(results) > 0 {
			possibleCombinations[target] = results
			//if len(results) == 0 {
			//	fmt.Println("No combinations found.")
			//} else {
			//	fmt.Println("Possible combinations:")
			//	for _, res := range results {
			//		fmt.Println(res)
			//	}
			//}
			//fmt.Println("--------------------")
		}
	}
	return possibleCombinations
}

// Function to pretty-print a map
func prettyPrintMap(m map[int][]string) {
	// Marshal the map to JSON with indentation
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling map to JSON:", err)
		return
	}

	// Print the JSON string
	fmt.Println(string(jsonData))
}

func printMap(myMap map[int][]string) {
	// Loop through the map keys and values
	for key, value := range myMap {
		fmt.Printf("Result: %v, Combinations: %v\n", key, value)
	}
}

func getSumOfKeys(myMap map[int][]string) int {
	var sum int = 0
	for key := range myMap {
		sum += key
	}
	return sum
}

func solve(res int, args []int) int {
	operation_result := 0
	last_op := int(math.Pow(2, float64(len(args))) - 1)
	for i := 0; i < last_op; i++ {
		operation_result = args[0]
		for j := 1; j < len(args); j++ {
			current_op := int(math.Pow(2, float64(j-1)))
			if current_op&i != 0 {
				operation_result += args[j]
			} else {
				operation_result *= args[j]
			}
		}
		if res == operation_result {
			return res
		}
	}
	return 0
}
func countTrueEquations(equations [][]int) int {
	totalSolutions := 0
	for _, equation := range equations {
		// Example inputs
		target := equation[0]
		numbers := equation[1:]

		result := solve(target, numbers)
		if result != 0 {
			//fmt.Printf("target: %v, result: %v\n", target, result)
			totalSolutions += result
		}

	}
	return totalSolutions
}

func solvePart1(input string) int {
	equations := readEquations(input)
	return countTrueEquations(equations)
}

func solve2(res int, args []int) int {
	operation_result := 0
	last_op := int(math.Pow(3, float64(len(args))) - 1)
	for i := 0; i < last_op; i++ {
		operation_result = args[0]
		for j := 1; j < len(args); j++ {
			current_op := i
			for k := 1; k < j; k++ {
				current_op /= 3
			}
			if current_op%3 == 0 {
				operation_result += args[j]
			} else if current_op%3 == 1 {
				operation_result *= args[j]
			} else {
				operation_result = concatNumbers(operation_result, args[j])
			}
		}
		if res == operation_result {
			return res
		}
	}
	return 0
}

func concatNumbers(n1, n2 int) int {
	n_aux1, n_aux2 := n1, n2
	for n_aux2 > 0 {
		n_aux2 /= 10
		n_aux1 *= 10
	}
	return n_aux1 + n2
}

func countTrueEquationsPart2(equations [][]int) int {
	totalSolutions := 0
	for _, equation := range equations {
		// Example inputs
		target := equation[0]
		numbers := equation[1:]

		result := solve2(target, numbers)
		if result != 0 {
			//fmt.Printf("target: %v, result: %v\n", target, result)
			totalSolutions += result
		}

	}
	return totalSolutions
}

func solvePart2(input string) int {
	equations := readEquations(input)
	return countTrueEquationsPart2(equations)
}
