package main

import (
	"fmt"
	"github.com/jpillora/puzzler/harness/aoc"
	"math"
	"sort"
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

func extractInputPart1(input string) ([]int, []int) {
	var list1 []int
	var list2 []int

	for _, row := range strings.Split(input, "\n") {
		fields := strings.Fields(row)

		list1 = append(list1, convertStrToInt(fields[0]))
		list2 = append(list2, convertStrToInt(fields[1]))
	}

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})

	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})
	return list1, list2
}

func calculateVariance(list1 []int, list2 []int) int {
	totalVariance := 0
	for i := 0; i < len(list1); i++ {
		variance := int(math.Abs(float64(list1[i] - list2[i])))
		totalVariance += variance
	}
	//fmt.Printf("variances: %v\n", variances)
	fmt.Printf("total variances: %v\n", totalVariance)
	return totalVariance
}

func solvePart1(input string) any {
	list1, list2 := extractInputPart1(input)
	totalVariance := calculateVariance(list1, list2)
	return totalVariance
}

func extractInputPart2(input string) ([]string, []string) {
	var list1 []string
	var list2 []string

	for _, row := range strings.Split(input, "\n") {
		fields := strings.Fields(row)

		list1 = append(list1, fields[0])
		list2 = append(list2, fields[1])
	}
	return list1, list2
}

func mapDuplicateCharacters(slice []string) map[string]int {
	// Initialize an empty map to store counts
	countMap := make(map[string]int)

	// Iterate over the slice
	for _, str := range slice {
		// Increment the count for each string in the map
		countMap[str]++
	}
	return countMap
}

func calculateSimilarity(slice1 []string, map2 map[string]int) int {
	totalSimilarities := 0

	// Iterate through the map
	for _, row := range slice1 {
		occurrencesInList2 := map2[row]
		similarity := convertStrToInt(row) * occurrencesInList2
		totalSimilarities += similarity
	}
	//fmt.Printf("variances: %v\n", variances)
	fmt.Printf("total similarities: %v\n", totalSimilarities)
	return totalSimilarities
}

func solvePart2(input string) any {
	list1, list2 := extractInputPart2(input)
	map2 := mapDuplicateCharacters(list2)
	totalSimilarity := calculateSimilarity(list1, map2)
	return totalSimilarity
}
