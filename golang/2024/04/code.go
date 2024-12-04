package main

import (
	"fmt"
	"github.com/jpillora/puzzler/harness/aoc"
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
	// when you're ready to do part 2, remove this "not implemented" block.
	if part2 {
		return solvePart2(input)
	}
	// solve part 1 here.
	return solvePart1(input)
}

func createWordSearchSpace(input string) [][]string {
	var wordSearch [][]string
	for _, line := range strings.Fields(input) {
		lineChars := strings.Split(line, "")
		wordSearch = append(wordSearch, lineChars)
	}

	return wordSearch
}

func isReverseMatch(input, searchWord string) bool {
	// Normalize the strings: remove spaces and convert to lowercase
	input = strings.ToLower(strings.ReplaceAll(input, " ", ""))
	searchWord = strings.ToLower(strings.ReplaceAll(searchWord, " ", ""))

	if input == searchWord {
		return true
	}
	// Reverse the input string
	runes := []rune(input)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	reversedInput := string(runes)

	// Compare the reversed input with the search word
	return reversedInput == searchWord || input == searchWord
}

func rightSearch(searchSpace [][]string, rowIndex int, colIndex int, searchWord string, wordLength int, stringsFound [][]int) [][]int {
	// fmt.Printf("Can go right\n")
	rightString := searchSpace[rowIndex][colIndex : colIndex+wordLength]

	if isReverseMatch(strings.Join(rightString, ""), searchWord) {
		//fmt.Printf("Char: %v, Index: (%v, %v)\n", searchSpace[rowIndex][colIndex], rowIndex, colIndex)
		//fmt.Printf("Right String: %v\n", rightString)
		var tmpSlice []int
		// Generate consecutive numbers
		for i := colIndex; i < colIndex+wordLength; i++ {
			tmpSlice = append(tmpSlice, rowIndex, i)
		}
		stringsFound = append(stringsFound, tmpSlice)
		//fmt.Printf("Strings Found: %v\n", stringsFound)
		//fmt.Println("---------------------------------")
	}

	return stringsFound
}

func leftSearch(searchSpace [][]string, rowIndex int, colIndex int, searchWord string, wordLength int, stringsFound [][]int) [][]int {
	// fmt.Printf("Can go left\n")
	leftString := searchSpace[rowIndex][colIndex-wordLength : colIndex+1]

	if isReverseMatch(strings.Join(leftString, ""), searchWord) {
		//fmt.Printf("Char: %v, Index: (%v, %v)\n", searchSpace[rowIndex][colIndex], rowIndex, colIndex)
		//fmt.Printf("Left String: %v\n", leftString)
		var tmpSlice []int
		// Generate consecutive numbers
		for i := colIndex - wordLength; i <= colIndex; i++ {
			tmpSlice = append(tmpSlice, rowIndex, i)
		}
		stringsFound = append(stringsFound, tmpSlice)
		//fmt.Printf("Strings Found: %v\n", stringsFound)
		//fmt.Println("---------------------------------")
	}

	return stringsFound
}

func upSearch(searchSpace [][]string, rowIndex int, colIndex int, searchWord string, wordLength int, stringsFound [][]int) [][]int {
	// fmt.Printf("Char: %v, Index: (%v, %v)\n", searchSpace[rowIndex][colIndex], rowIndex, colIndex)
	// fmt.Printf("Can go up\n")

	// Extract 'word length' elements from the specified column
	var upString []string
	var tmpSlice []int
	var searchRow int
	var searchCol int
	for i := rowIndex - wordLength; i <= rowIndex; i++ {
		searchRow = i
		searchCol = colIndex
		upString = append(upString, searchSpace[searchRow][searchCol])
		tmpSlice = append(tmpSlice, searchRow, searchCol)
	}

	if isReverseMatch(strings.Join(upString, ""), searchWord) {
		//fmt.Printf("Char: %v, Index: (%v, %v)\n", searchSpace[rowIndex][colIndex], rowIndex, colIndex)
		//fmt.Printf("Up String: %v\n", upString)
		stringsFound = append(stringsFound, tmpSlice)
		//fmt.Printf("Strings Found: %v\n", stringsFound)
		//fmt.Println("---------------------------------")
	}

	return stringsFound
}

func upRightSearch(searchSpace [][]string, rowIndex int, colIndex int, searchWord string, wordLength int, stringsFound [][]int) [][]int {
	// fmt.Printf("Can go up-right\n")

	// Extract 'word length' elements from the specified column
	var upRightString []string
	var tmpSlice []int
	var searchRow int
	var searchCol int
	for i := 0; i <= wordLength; i++ {
		searchRow = rowIndex - i
		searchCol = colIndex + i
		//fmt.Printf("i: %v, searchRow: %v, searchCol: %v\n", i, searchRow, searchCol)
		upRightString = append(upRightString, searchSpace[searchRow][searchCol])
		tmpSlice = append(tmpSlice, searchRow, searchCol)
	}

	if isReverseMatch(strings.Join(upRightString, ""), searchWord) {
		//fmt.Printf("Char: %v, Index: (%v, %v)\n", searchSpace[rowIndex][colIndex], rowIndex, colIndex)
		//fmt.Printf("Up Right String: %v\n", upRightString)
		stringsFound = append(stringsFound, tmpSlice)
		//fmt.Printf("Strings Found: %v\n", stringsFound)
		//fmt.Println("---------------------------------")
	}

	return stringsFound
}

func upLeftSearch(searchSpace [][]string, rowIndex int, colIndex int, searchWord string, wordLength int, stringsFound [][]int) [][]int {
	//fmt.Printf("Can go up-left\n")

	// Extract 'word length' elements from the specified column
	var upLeftString []string
	var tmpSlice []int
	var searchRow int
	var searchCol int
	for i := wordLength; i >= 0; i-- {
		searchRow = rowIndex - i
		searchCol = colIndex - i
		upLeftString = append(upLeftString, searchSpace[searchRow][searchCol])
		tmpSlice = append(tmpSlice, searchRow, searchCol)
	}

	if isReverseMatch(strings.Join(upLeftString, ""), searchWord) {
		//fmt.Printf("Char: %v, Index: (%v, %v)\n", searchSpace[rowIndex][colIndex], rowIndex, colIndex)
		//fmt.Printf("Up Left String: %v\n", upLeftString)
		stringsFound = append(stringsFound, tmpSlice)
		//fmt.Printf("Strings Found: %v\n", stringsFound)
		//fmt.Println("---------------------------------")
	}

	return stringsFound
}

func downSearch(searchSpace [][]string, rowIndex int, colIndex int, searchWord string, wordLength int, stringsFound [][]int) [][]int {
	// fmt.Printf("Can down\n")

	// Extract 'word length' elements from the specified column
	var downString []string
	var tmpSlice []int
	var searchRow int
	var searchCol int
	for i := rowIndex; i <= rowIndex+wordLength; i++ {
		searchRow = i
		searchCol = colIndex
		downString = append(downString, searchSpace[searchRow][searchCol])
		tmpSlice = append(tmpSlice, searchRow, searchCol)
	}

	if isReverseMatch(strings.Join(downString, ""), searchWord) {
		//fmt.Printf("Char: %v, Index: (%v, %v)\n", searchSpace[rowIndex][colIndex], rowIndex, colIndex)
		//fmt.Printf("Down String: %v\n", downString)
		stringsFound = append(stringsFound, tmpSlice)
		//fmt.Printf("Strings Found: %v\n", stringsFound)
		//fmt.Println("---------------------------------")
	}

	return stringsFound
}

func downRightSearch(searchSpace [][]string, rowIndex int, colIndex int, searchWord string, wordLength int, stringsFound [][]int) [][]int {
	// fmt.Printf("Can go down-right\n")

	// Extract 'word length' elements from the specified column
	var downRightString []string
	var tmpSlice []int
	var searchRow int
	var searchCol int
	for i := 0; i <= wordLength; i++ {
		searchRow = rowIndex + i
		searchCol = colIndex + i
		downRightString = append(downRightString, searchSpace[searchRow][searchCol])
		tmpSlice = append(tmpSlice, searchRow, searchCol)
	}

	if isReverseMatch(strings.Join(downRightString, ""), searchWord) {
		//fmt.Printf("Char: %v, Index: (%v, %v)\n", searchSpace[rowIndex][colIndex], rowIndex, colIndex)
		//fmt.Printf("Down Right String: %v\n", downRightString)
		stringsFound = append(stringsFound, tmpSlice)
		//fmt.Printf("Strings Found: %v\n", stringsFound)
		//fmt.Println("---------------------------------")
	}

	return stringsFound
}

func downLeftSearch(searchSpace [][]string, rowIndex int, colIndex int, searchWord string, wordLength int, stringsFound [][]int) [][]int {
	// fmt.Printf("Can go down-left\n")

	// Extract 'word length' elements from the specified column
	var downLeftString []string
	var tmpSlice []int
	var searchRow int
	var searchCol int
	for i := wordLength; i >= 0; i-- {
		searchRow = rowIndex + i
		searchCol = colIndex - i
		downLeftString = append(downLeftString, searchSpace[searchRow][searchCol])
		tmpSlice = append(tmpSlice, searchRow, searchCol)
	}

	if isReverseMatch(strings.Join(downLeftString, ""), searchWord) {
		//fmt.Printf("Char: %v, Index: (%v, %v)\n", searchSpace[rowIndex][colIndex], rowIndex, colIndex)
		//fmt.Printf("Down Left String: %v\n", downLeftString)
		stringsFound = append(stringsFound, tmpSlice)
		//fmt.Printf("Strings Found: %v\n", stringsFound)
		//fmt.Println("---------------------------------")
	}

	return stringsFound
}

func wordSearchSpace(searchWord string, searchWordLength int, searchSpace [][]string, rowIndex int, colIndex int, stringsFound [][]int) [][]int {
	searchRadius := searchWordLength - 1
	maxRowRadius := len(searchSpace[colIndex]) - 1
	maxColRadius := len(searchSpace[rowIndex]) - 1
	// check if we can go right
	if colIndex+searchRadius <= maxColRadius {
		stringsFound = rightSearch(searchSpace, rowIndex, colIndex, searchWord, searchWordLength, stringsFound)
	}

	// check if we can go left
	if colIndex-searchRadius >= 0 {
		stringsFound = leftSearch(searchSpace, rowIndex, colIndex, searchWord, searchRadius, stringsFound)
	}

	// check if we can go up
	if rowIndex-searchRadius >= 0 {
		stringsFound = upSearch(searchSpace, rowIndex, colIndex, searchWord, searchRadius, stringsFound)
	}

	// check if we can go down
	if rowIndex+searchRadius <= maxRowRadius {
		stringsFound = downSearch(searchSpace, rowIndex, colIndex, searchWord, searchRadius, stringsFound)
	}

	// check if we can go up-right
	if colIndex+searchRadius <= maxColRadius && rowIndex-searchRadius >= 0 {
		stringsFound = upRightSearch(searchSpace, rowIndex, colIndex, searchWord, searchRadius, stringsFound)
	}

	// check if we can go down-right
	if colIndex+searchRadius <= maxColRadius && rowIndex+searchRadius <= maxRowRadius {
		stringsFound = downRightSearch(searchSpace, rowIndex, colIndex, searchWord, searchRadius, stringsFound)
	}

	// check if we can go down-left
	if colIndex-searchRadius >= 0 && rowIndex+searchRadius <= maxRowRadius {
		stringsFound = downLeftSearch(searchSpace, rowIndex, colIndex, searchWord, searchRadius, stringsFound)
	}

	// check if we can go up-left
	if colIndex-searchRadius >= 0 && rowIndex-searchRadius >= 0 {
		stringsFound = upLeftSearch(searchSpace, rowIndex, colIndex, searchWord, searchRadius, stringsFound)
	}

	return stringsFound
}

func xSearch(searchSpace [][]string, rowIndex int, colIndex int, searchWord string, wordLength int, stringsFound [][]int) [][]int {
	//fmt.Printf("Looking for X shape word: %v\n", searchWord)
	//fmt.Printf("Char: %v, Index: (%v, %v)\n", searchSpace[rowIndex][colIndex], rowIndex, colIndex)

	// Extract 'word length' elements from the specified column
	upRightDownLeftRow := rowIndex + wordLength
	upLeftDownRightRow := rowIndex - wordLength
	upRightCol := colIndex - wordLength
	var upRightDownLeftString []string
	var upLeftDownRightString []string
	var tmpSlice []int
	var coordSlice []int
	var searchRowDownLeft int
	var searchRowUpLeft int
	var searchCol int
	for i := 0; i <= wordLength+1; i++ {
		searchRowDownLeft = upRightDownLeftRow - i
		searchRowUpLeft = upLeftDownRightRow + i
		searchCol = upRightCol + i
		// fmt.Printf("i: %v, searchRow: %v, searchCol: %v\n", i, searchRowDownLeft, searchCol)
		upRightDownLeftString = append(upRightDownLeftString, searchSpace[searchRowDownLeft][searchCol])
		tmpSlice = append(tmpSlice, searchRowDownLeft, searchCol)

		// fmt.Printf("i: %v, searchRow: %v, searchCol: %v\n", i, searchRowUpLeft, searchCol)
		upLeftDownRightString = append(upLeftDownRightString, searchSpace[searchRowUpLeft][searchCol])
		tmpSlice = append(tmpSlice, searchRowUpLeft, searchCol)
	}

	//fmt.Printf("Up Right Down Left String: %v\n", upRightDownLeftString)
	//fmt.Printf("Up Left Down Right String: %v\n", upLeftDownRightString)

	if isReverseMatch(strings.Join(upRightDownLeftString, ""), searchWord) && isReverseMatch(strings.Join(upLeftDownRightString, ""), searchWord) {
		//fmt.Printf("Char: %v, Index: (%v, %v)\n", searchSpace[rowIndex][colIndex], rowIndex, colIndex)
		//fmt.Printf("Up Right String: %v\n", upRightString)
		coordSlice = append(coordSlice, rowIndex, colIndex)
		stringsFound = append(stringsFound, coordSlice)
		//fmt.Printf("Strings Found: %v\n", stringsFound)
		//fmt.Println("---------------------------------")
	}

	return stringsFound
}

func wordSearchXSpace(searchWord string, searchWordLength int, searchSpace [][]string, rowIndex int, colIndex int, stringsFound [][]int) [][]int {
	searchRadius := searchWordLength - 2
	maxRowRadius := len(searchSpace[colIndex]) - 1
	maxColRadius := len(searchSpace[rowIndex]) - 1

	var xesFound [][]int

	// check if we can go up-right, down-right, up-left and down-left
	if colIndex+searchRadius <= maxColRadius && rowIndex+searchRadius <= maxRowRadius && colIndex-searchRadius >= 0 && rowIndex-searchRadius >= 0 {
		xesFound = xSearch(searchSpace, rowIndex, colIndex, searchWord, searchRadius, xesFound)
		stringsFound = append(stringsFound, xesFound...)

		//fmt.Printf("Char: %v, Index: (%v, %v)\n", searchSpace[rowIndex][colIndex], rowIndex, colIndex)
		//fmt.Printf("Xes Found: %v\n", xesFound)
		//fmt.Printf("Strings Found: %v\n", stringsFound)
		//fmt.Println("---------------------------------")
	}

	return stringsFound
}

// Helper function to compare two slices of integers for equality
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Function to remove duplicate slices
func removeDuplicates(matrix [][]int) [][]int {
	var result [][]int

	// Iterate over the matrix (slice of slices)
	for _, row := range matrix {
		// Check if this row is already in the result
		duplicate := false
		for _, existingRow := range result {
			if slicesEqual(row, existingRow) {
				duplicate = true
				break
			}
		}
		// If not a duplicate, add it to the result
		if !duplicate {
			result = append(result, row)
		}
	}
	return result
}

func solvePart1(input string) int {
	wordSearch := createWordSearchSpace(input)
	// fmt.Printf("Word Search: %v\n", wordSearch)

	// find all XMAS in the word search
	// can be any direction with a straight line
	wordToFind := "XMAS"
	wordToFindLength := len(wordToFind)
	var stringsFound [][]int
	for l, line := range wordSearch {
		for c, _ := range line {
			stringsFound = wordSearchSpace(wordToFind, wordToFindLength, wordSearch, l, c, stringsFound)
		}
	}

	fmt.Printf("Matches Found: %v\n", len(stringsFound))

	uniqueMatches := removeDuplicates(stringsFound)
	fmt.Printf("Unique Matches Found: %v\n", len(uniqueMatches))

	return len(uniqueMatches)
}

func solvePart2(input string) int {
	wordSearch := createWordSearchSpace(input)
	// fmt.Printf("Word Search: %v\n", wordSearch)

	// find all X-MAS in the word search
	// MAS in the shape of an X
	wordToFind := "MAS"
	wordToFindLength := len(wordToFind)
	var stringsFound [][]int
	for l, line := range wordSearch {
		for c, _ := range line {
			stringsFound = wordSearchXSpace(wordToFind, wordToFindLength, wordSearch, l, c, stringsFound)
		}
	}

	fmt.Printf("Matches Found: %v\n", len(stringsFound))
	return len(stringsFound)
}
