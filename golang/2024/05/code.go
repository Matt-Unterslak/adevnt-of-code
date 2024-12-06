package main

import (
	"container/list"
	"fmt"
	"github.com/jpillora/puzzler/harness/aoc"
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
// the return value of each run is printed to stdout.
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return solvePart2(input)
	}
	// solve part 1 here
	return solvePart1(input)
}

func constructOrderingMap(orderingRules []string) map[string][]string {
	// Initialize the map
	result := make(map[string][]string)

	// Iterate over the slice
	for _, item := range orderingRules {
		// Split the string by "|"
		parts := strings.Split(item, "|")
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]

			// Append the value to the corresponding key's slice
			result[key] = append(result[key], value)
		}
	}

	return result
}

func checkPageOrdering(slice1 []string, slice2 []string) bool {
	// Create a map to track elements in slice2
	exists := make(map[string]bool)
	for _, value := range slice2 {
		exists[value] = true
	}

	// Check if all elements in slice1 exist in the map
	for _, value := range slice1 {
		if exists[value] {
			return false
		}
	}

	return true
}

func getCorrectAndIncorrectUpdates(pagesToPrint []string, orderRulesMap map[string][]string) ([][]string, [][]string) {
	var correctUpdates [][]string
	var incorrectUpdates [][]string
	// for each page, print the page
	for _, pages := range pagesToPrint {
		//fmt.Printf("Pages: %v\n", pages)
		pagesToCheck := strings.Split(pages, ",")
		updateCorrect := true
		for pageIndex := len(pagesToCheck) - 1; pageIndex > 0; pageIndex-- {
			page := pagesToCheck[pageIndex]
			//fmt.Printf("Page: %v, Page Index: %d\n", page, pageIndex)
			previousPages := pagesToCheck[:pageIndex]
			//fmt.Printf(" PreviousPages: %v\n", previousPages)
			pageMap := orderRulesMap[page]
			//fmt.Printf("Page Map: %v\n", pageMap)

			updateCorrect = checkPageOrdering(previousPages, pageMap)
			//fmt.Printf("Ordering Correct: %v\n", updateCorrect)
			if !updateCorrect {
				incorrectUpdates = append(incorrectUpdates, pagesToCheck)
				break
			}
		}
		if updateCorrect {
			correctUpdates = append(correctUpdates, pagesToCheck)
		}
	}

	return correctUpdates, incorrectUpdates
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

func getMiddleNumber(slice1 []string) int {
	median := len(slice1) / 2
	return convertStrToInt(slice1[median])
}

func solvePart1(input string) int {
	splitInput := strings.Split(input, "\n\n")
	orderRules := strings.Fields(splitInput[0])
	orderRulesMap := constructOrderingMap(orderRules)

	// Print the resulting map
	//fmt.Printf("Order Rules Map: %v\n", orderRulesMap)

	pagesToPrint := strings.Fields(splitInput[1])
	//fmt.Printf("Pages to print: %v\n", pagesToPrint)

	//fmt.Println("--------------------------------")

	correctUpdates, _ := getCorrectAndIncorrectUpdates(pagesToPrint, orderRulesMap)

	//fmt.Printf("Correct Updates: %v\n", correctUpdates)
	//fmt.Printf("Incorrect Updates: %v\n", incorrectUpdates)

	sumMiddleNumbers := 0
	for _, correctUpdate := range correctUpdates {
		middleNumber := getMiddleNumber(correctUpdate)
		//fmt.Printf("Middle Number: %v\n", middleNumber)
		sumMiddleNumbers += middleNumber
	}

	fmt.Printf("Sum of Middle Numbers %v\n", sumMiddleNumbers)
	//fmt.Println("--------------------------------")

	return sumMiddleNumbers
}

type Deque[T any] struct {
	data *list.List
}

// NewDeque creates a new empty deque
func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{data: list.New()}
}

// PushFront adds an element to the front of the deque
func (d *Deque[T]) PushFront(value T) {
	d.data.PushFront(value)
}

// PushBack adds an element to the back of the deque
func (d *Deque[T]) PushBack(value T) {
	d.data.PushBack(value)
}

// PopFront removes and returns the element at the front of the deque
func (d *Deque[T]) PopFront() (T, bool) {
	frontElement := d.data.Front()
	if frontElement == nil {
		var zeroValue T
		return zeroValue, false
	}
	d.data.Remove(frontElement)
	return frontElement.Value.(T), true
}

// PopBack removes and returns the element at the back of the deque
func (d *Deque[T]) PopBack() (T, bool) {
	backElement := d.data.Back()
	if backElement == nil {
		var zeroValue T
		return zeroValue, false
	}
	d.data.Remove(backElement)
	return backElement.Value.(T), true
}

// PeekFront returns the element at the front without removing it
func (d *Deque[T]) PeekFront() (T, bool) {
	frontElement := d.data.Front()
	if frontElement == nil {
		var zeroValue T
		return zeroValue, false
	}
	return frontElement.Value.(T), true
}

// PeekBack returns the element at the back without removing it
func (d *Deque[T]) PeekBack() (T, bool) {
	backElement := d.data.Back()
	if backElement == nil {
		var zeroValue T
		return zeroValue, false
	}
	return backElement.Value.(T), true
}

// IsEmpty checks if the deque is empty
func (d *Deque[T]) IsEmpty() bool {
	return d.data.Len() == 0
}

// Size returns the number of elements in the deque
func (d *Deque[T]) Size() int {
	return d.data.Len()
}

// PrintDeque prints all elements in the deque
func (d *Deque[T]) PrintDeque() {
	for e := d.data.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

func excludeCurrentOptimized(slice []string, currentIndex int) []string {
	// Pre-allocate a new slice of the exact required size
	result := make([]string, 0, len(slice)-1)

	// Add elements before the current index
	if currentIndex > 0 {
		result = append(result, slice[:currentIndex]...)
	}

	// Add elements after the current index
	if currentIndex < len(slice)-1 {
		result = append(result, slice[currentIndex+1:]...)
	}

	return result
}

// Remove removes the element at the specified index from the slice
func Remove(slice []string, index int) []string {
	// Check if the index is within range
	if index < 0 || index >= len(slice) {
		return slice // Return the original slice if the index is out of range
	}

	// Remove the element by combining slices before and after the index
	return append(slice[:index], slice[index+1:]...)
}

func findNextTail(correctedUpdate *Deque[string], incorrectUpdate []string, orderRulesMap map[string][]string) []string {

	if len(incorrectUpdate) == 1 {
		correctedUpdate.PushFront(incorrectUpdate[0])
		return make([]string, 0)
	}

	var potentialTail [][]string
	for index, page := range incorrectUpdate {
		otherPages := excludeCurrentOptimized(incorrectUpdate, index)
		//fmt.Printf("Page: %v, Other Pages: %v\n", page, otherPages)
		pageMap := orderRulesMap[page]
		if checkPageOrdering(otherPages, pageMap) {
			var tmpSlice []string
			tmpSlice = append(tmpSlice, page, strconv.Itoa(index))
			potentialTail = append(potentialTail, tmpSlice)
		}
	}
	//fmt.Printf("Potential Tail: %v\n", potentialTail)

	if len(potentialTail) == 0 {
		panic("No potential tail found")
	} else if len(potentialTail) > 1 {
		panic("More than one potential tail found")
	}

	correctedUpdate.PushFront(potentialTail[0][0])

	return Remove(incorrectUpdate, convertStrToInt(potentialTail[0][1]))
}

func fixIncorrectUpdate(incorrectUpdate []string, orderRulesMap map[string][]string) []string {
	//fmt.Printf("Incorrect Update: %v\n", incorrectUpdate)

	// Create a deque of strings
	correctedUpdate := NewDeque[string]()
	remainingPages := incorrectUpdate

	cycles := 10
	for {
		if len(remainingPages) == 0 || cycles <= 0 {
			break
		}
		remainingPages = findNextTail(correctedUpdate, remainingPages, orderRulesMap)
		//fmt.Printf("Remaining Pages: %v\n", remainingPages)

		cycles--
	}

	//fmt.Println("Corrected Update")
	//correctedUpdate.PrintDeque()
	return incorrectUpdate
}

func fixIncorrectUpdates(incorrectUpdates [][]string, orderRulesMap map[string][]string) [][]string {
	var correctedUpdates [][]string
	for _, incorrectUpdate := range incorrectUpdates {
		correctedUpdate := fixIncorrectUpdate(incorrectUpdate, orderRulesMap)
		correctedUpdates = append(correctedUpdates, correctedUpdate)
		//fmt.Println("--------------------------------")
	}

	return correctedUpdates
}

func solvePart2(input string) int {
	splitInput := strings.Split(input, "\n\n")
	orderRules := strings.Fields(splitInput[0])
	orderRulesMap := constructOrderingMap(orderRules)

	// Print the resulting map
	//fmt.Printf("Order Rules Map: %v\n", orderRulesMap)

	pagesToPrint := strings.Fields(splitInput[1])
	//fmt.Printf("Pages to print: %v\n", pagesToPrint)

	//fmt.Println("--------------------------------")

	_, incorrectUpdates := getCorrectAndIncorrectUpdates(pagesToPrint, orderRulesMap)

	//fmt.Printf("Correct Updates: %v\n", correctUpdates)
	//fmt.Printf("Incorrect Updates: %v\n", incorrectUpdates)
	//fmt.Println("--------------------------------")

	correctedUpdates := fixIncorrectUpdates(incorrectUpdates, orderRulesMap)

	sumMiddleNumbers := 0
	for _, correctedUpdate := range correctedUpdates {
		middleNumber := getMiddleNumber(correctedUpdate)
		//fmt.Printf("Middle Number: %v\n", middleNumber)
		sumMiddleNumbers += middleNumber
	}

	//fmt.Printf("Sum of Middle Numbers %v\n", sumMiddleNumbers)
	//fmt.Println("--------------------------------")

	// correct answer: 5169
	return sumMiddleNumbers
}
