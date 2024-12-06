package main

import (
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
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return solvePart2(input)
	}
	// solve part 1 here
	return solvePart1(input)
}

func readInput(input string) ([]int, [][]int) {
	var grid [][]int
	var startPosition []int
	for rowIndex, row := range strings.Split(input, "\n") {
		var rowSlice []int
		for colIndex := 0; colIndex < len(row); colIndex++ {
			// fmt.Printf("Row Index: %v, Col Index: %v, Value: %s\n", rowIndex, colIndex, string(row[colIndex]))
			value := string(row[colIndex])
			var gridValue int
			switch value {
			case "^":
				gridValue = 1
				startPosition = append(startPosition, rowIndex, colIndex, 1)
			case ">":
				gridValue = 1
				startPosition = append(startPosition, rowIndex, colIndex, 2)
			case "v":
				gridValue = 1
				startPosition = append(startPosition, rowIndex, colIndex, 3)
			case "<":
				gridValue = 1
				startPosition = append(startPosition, rowIndex, colIndex, 4)
			case "#":
				gridValue = -1
			default:
				gridValue = 0
			}

			rowSlice = append(rowSlice, gridValue)
		}
		grid = append(grid, rowSlice)

	}
	return startPosition, grid
}

func readInputPart2(input string) ([]int, [][]int, [][]int) {
	var grid [][]int
	var possibleObstaclePositions [][]int
	var startPosition []int
	for rowIndex, row := range strings.Split(input, "\n") {
		var rowSlice []int
		for colIndex := 0; colIndex < len(row); colIndex++ {
			// fmt.Printf("Row Index: %v, Col Index: %v, Value: %s\n", rowIndex, colIndex, string(row[colIndex]))
			value := string(row[colIndex])
			var gridValue int
			switch value {
			case "^":
				gridValue = 1
				startPosition = append(startPosition, rowIndex, colIndex, 1)
			case ">":
				gridValue = 1
				startPosition = append(startPosition, rowIndex, colIndex, 2)
			case "v":
				gridValue = 1
				startPosition = append(startPosition, rowIndex, colIndex, 3)
			case "<":
				gridValue = 1
				startPosition = append(startPosition, rowIndex, colIndex, 4)
			case "#":
				gridValue = -1
			default:
				gridValue = 0
				possibleObstaclePositions = append(possibleObstaclePositions, []int{rowIndex, colIndex})
			}

			rowSlice = append(rowSlice, gridValue)
		}
		grid = append(grid, rowSlice)

	}
	return startPosition, grid, possibleObstaclePositions
}

func cloneGrid(grid [][]int) [][]int {
	// Create a new slice with the same outer size
	newGrid := make([][]int, len(grid))

	for i := range grid {
		// Allocate a new inner slice for each row
		newGrid[i] = make([]int, len(grid[i]))
		copy(newGrid[i], grid[i]) // Copy the contents of the row
	}

	return newGrid
}

func printGrid(grid [][]int) {
	for i, row := range grid {
		fmt.Printf("Row %v: %v\n", i, row)
	}
	fmt.Println("--------------------")
}

func calculateNewGridPosition(previousX int, previousY int, previousDirection int, grid [][]int) ([]int, int) {
	currentDirection := previousDirection
	currentX := previousX
	currentY := previousY

	// 1 = up, 2 = right, 3 = down, 4 = left
	guardedGridPoints := 0
	switch currentDirection {
	case 1:
		currentX -= 1
	case 2:
		currentY += 1
	case 3:
		currentX += 1
	case 4:
		currentY -= 1
	}
	if currentX < 0 || currentX >= len(grid[0]) || currentY < 0 || currentY >= len(grid) {
		return []int{currentX, currentY, currentDirection}, guardedGridPoints
	} else if grid[currentX][currentY] == -1 {
		currentDirection += 1
		if currentDirection == 5 {
			currentDirection = 1
		}
		return []int{previousX, previousY, currentDirection}, 0
	} else {
		if grid[currentX][currentY] == -1 {
			fmt.Printf("Current Position: %v, %v, %v\n", currentX, currentY, currentDirection)
			panic("Invalid grid position")
		}
		if grid[currentX][currentY] == 0 {
			guardedGridPoints = 1
		}
		grid[currentX][currentY] += 1
		return []int{currentX, currentY, currentDirection}, guardedGridPoints
	}

	//if grid[currentX][currentY] == 0 {
	//	grid[currentX][currentY] = 1
	//} else if grid[currentX][currentY] == 1 {
	//	grid[currentX][currentY] = 0
	//}

}

func guardMovement(startPosition []int, grid [][]int) int {
	guardedGridPoints := 1
	currentX := startPosition[0]
	currentY := startPosition[1]
	currentDirection := startPosition[2]
	maxX := len(grid[0])
	maxY := len(grid)
	updatedGrid := cloneGrid(grid)
	//printGrid(updatedGrid)
	for {
		newPosition, newPositionGuarded := calculateNewGridPosition(currentX, currentY, currentDirection, updatedGrid)

		currentX = newPosition[0]
		currentY = newPosition[1]
		currentDirection = newPosition[2]
		guardedGridPoints += newPositionGuarded

		// fmt.Printf("Current Position: %v, %v, %v\n", currentX, currentY, currentDirection)
		if currentX < 0 || currentX >= maxX || currentY < 0 || currentY >= maxY {
			break
		}
	}
	//printGrid(updatedGrid)
	return guardedGridPoints
}

func solvePart1(input string) int {
	// fmt.Printf("Input: %v/n", input)
	startPosition, grid := readInput(input)
	//fmt.Printf("Start: %v\n", startPosition)
	guardedPoints := guardMovement(startPosition, grid)
	return guardedPoints
}

func guardMovementObstacle(startPosition []int, grid [][]int) []int {
	guardedGridPoints := 1
	currentX := startPosition[0]
	currentY := startPosition[1]
	currentDirection := startPosition[2]
	maxX := len(grid[0])
	maxY := len(grid)
	updatedGrid := cloneGrid(grid)
	//updatedGrid[6][3] = -1
	//printGrid(updatedGrid)
	cycles := 0
	var positionsTraversed map[string]int
	positionsTraversed = make(map[string]int)
	var stringPosition string
	stringPosition = strconv.Itoa(currentX) + strconv.Itoa(currentY)
	positionsTraversed[stringPosition] = 1
	for {
		newPosition, newPositionGuarded := calculateNewGridPosition(currentX, currentY, currentDirection, updatedGrid)
		//fmt.Printf("New Position: %v\n", newPosition)
		currentX = newPosition[0]
		currentY = newPosition[1]
		currentDirection = newPosition[2]
		guardedGridPoints += newPositionGuarded
		cycles += 1
		stringPosition = strconv.Itoa(currentX) + strconv.Itoa(currentY)
		if positionsTraversed[stringPosition] > 40 {
			return []int{6, 3}
		} else if value, ok := positionsTraversed[stringPosition]; ok {
			positionsTraversed[stringPosition] = value + 1
		} else {
			positionsTraversed[stringPosition] = 1
		}
		positionsTraversed[stringPosition] += 1

		//fmt.Printf("Current Position: %v, %v, %v\n", currentX, currentY, currentDirection)
		if currentX < 0 || currentX >= maxX || currentY < 0 || currentY >= maxY || cycles > 1000000 {
			break
		}

	}
	//printGrid(updatedGrid)
	return make([]int, 0)

}

func tryPossibleObstacles(startPosition []int, grid [][]int, possibleObstaclePositions [][]int) int {
	var cyclicalObstacles [][]int
	for _, possibleObstaclePosition := range possibleObstaclePositions {
		//fmt.Printf("Possible Obstacle: %v\n", possibleObstaclePosition)
		updatedGrid := cloneGrid(grid)
		updatedGrid[possibleObstaclePosition[0]][possibleObstaclePosition[1]] = -1
		detectCyclicalObstacle := guardMovementObstacle(startPosition, updatedGrid)
		if len(detectCyclicalObstacle) > 0 {
			//fmt.Printf("Cyclical Obstacle Detected: %v\n", possibleObstaclePosition)
			cyclicalObstacles = append(cyclicalObstacles, possibleObstaclePosition)
		}

		//fmt.Printf("Cyclical Obstacles: %v\n", cyclicalObstacles)
	}
	return len(cyclicalObstacles)

}

func solvePart2(input string) int {
	// fmt.Printf("Input: %v/n", input)
	startPosition, grid, possibleObstaclePositions := readInputPart2(input)
	//fmt.Printf("Start: %v\n", startPosition)
	//fmt.Printf("Possible Obstacles: %v\n", possibleObstaclePositions)
	//var testingObstacles [][]int
	//testingObstacles = append(testingObstacles, []int{6, 3})
	cyclicalObstacles := tryPossibleObstacles(startPosition, grid, possibleObstaclePositions)
	return cyclicalObstacles
}
