package main

import (
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

func solvePart1(input string) int {
	grid, trailHeads := constructGrid(input)
	totalScore, _ := constructTrailRoutes(grid, trailHeads)
	return totalScore
}

func solvePart2(input string) int {
	grid, trailHeads := constructGrid(input)
	_, totalRating := constructTrailRoutes(grid, trailHeads)
	return totalRating
}

func constructTrailRoutes(grid [][]int, trailHeads [][]int) (int, int) {
	//for _, value := range grid {
	//	fmt.Println(value)
	//}

	maxRow := len(grid) - 1
	maxCol := len(grid[0]) - 1

	trailMaps := make(map[int][]int)
	trailheadScore := 0
	trailheadScores := make(map[int]int)
	trailRatings := make(map[int]int)

	for startNum, startPosition := range trailHeads {

		trailPaths := make(map[int][]int)
		var possiblePaths [][]int
		//fmt.Println(startPosition)
		determinePossiblePaths(startPosition, grid, maxRow, maxCol, possiblePaths, trailPaths)
		//fmt.Println(trailPaths)
		trailRating := len(trailPaths)
		trailRatings[startNum] = trailRating

		endPoints := make(map[int][]int)
		for _, coord := range trailPaths {
			// Check if the key exists
			if value, exists := endPoints[coord[0]]; exists {
				found := false
				for _, col := range value {
					if col == coord[1] {
						//fmt.Printf("Point already mapped %v\n", coord)
						found = true
						break
					}
				}
				if !found {
					//fmt.Printf("Point not mapped %v\n", coord)
					endPoints[coord[0]] = append(endPoints[coord[0]], coord[1])
				}

			} else {
				var coordSlice []int
				coordSlice = append(coordSlice, coord[1])
				endPoints[coord[0]] = coordSlice

			}
		}
		//fmt.Println(endPoints)
		trailheadScore = updateTrailMaps(trailMaps, endPoints)
		//fmt.Printf("Trailhead score: %d\n", trailheadScore)
		//fmt.Println("-----------------------")
		trailheadScores[startNum] = trailheadScore
	}
	//fmt.Printf("Trailhead scores: %v\n", trailheadScores)
	//fmt.Printf("Trailhead ratings: %v\n", trailRatings)
	totalScore := sumTrailHeadScores(trailheadScores)
	totalRating := sumTrailRatings(trailRatings)
	return totalScore, totalRating
}

func sumTrailRatings(trailRatings map[int]int) int {
	totalScore := 0
	for _, value := range trailRatings {
		totalScore += value
	}
	return totalScore
}

func sumTrailHeadScores(trailheadScores map[int]int) int {
	totalScore := 0
	for _, value := range trailheadScores {
		totalScore += value
	}
	return totalScore
}

func updateTrailMaps(trailMaps map[int][]int, endpoints map[int][]int) int {
	mapKey := 0
	for key, value := range endpoints {
		//fmt.Printf("Key: %d, Value: %v\n", key, value)
		for _, col := range value {
			var mapSlice []int
			mapSlice = append(mapSlice, key, col)
			trailMaps[mapKey+1] = mapSlice
			mapKey++
		}
	}
	return mapKey
}

func determinePossiblePaths(currentPosition []int, grid [][]int, maxRow int, maxCol int, possiblePaths [][]int, trailPaths map[int][]int) [][]int {
	currentValue := currentPosition[2]
	if currentValue == 9 {
		//fmt.Printf("Found end: %v\n", currentPosition)
		updatePathing(trailPaths, currentPosition)
	}
	up := []int{currentPosition[0] - 1, currentPosition[1]}
	down := []int{currentPosition[0] + 1, currentPosition[1]}
	left := []int{currentPosition[0], currentPosition[1] - 1}
	right := []int{currentPosition[0], currentPosition[1] + 1}

	if up[0] >= 0 && up[0] <= maxRow && up[1] >= 0 && up[1] <= maxCol && grid[up[0]][up[1]]-currentValue == 1 {
		up = append(up, grid[up[0]][up[1]])
		possiblePaths = append(possiblePaths, up)
		possiblePaths = determinePossiblePaths(up, grid, maxRow, maxCol, possiblePaths, trailPaths)
	}

	if down[0] >= 0 && down[0] <= maxRow && down[1] >= 0 && down[1] <= maxCol && grid[down[0]][down[1]]-currentValue == 1 {
		down = append(down, grid[down[0]][down[1]])
		possiblePaths = append(possiblePaths, down)
		possiblePaths = determinePossiblePaths(down, grid, maxRow, maxCol, possiblePaths, trailPaths)
	}

	if left[0] >= 0 && left[0] <= maxRow && left[1] >= 0 && left[1] <= maxCol && grid[left[0]][left[1]]-currentValue == 1 {
		left = append(left, grid[left[0]][left[1]])
		possiblePaths = append(possiblePaths, left)
		possiblePaths = determinePossiblePaths(left, grid, maxRow, maxCol, possiblePaths, trailPaths)
	}

	if right[0] >= 0 && right[0] <= maxRow && right[1] >= 0 && right[1] <= maxCol && grid[right[0]][right[1]]-currentValue == 1 {
		right = append(right, grid[right[0]][right[1]])
		possiblePaths = append(possiblePaths, right)
		possiblePaths = determinePossiblePaths(right, grid, maxRow, maxCol, possiblePaths, trailPaths)
	}

	return possiblePaths

}

func updatePathing(trailMap map[int][]int, currentPosition []int) {
	if len(trailMap) == 0 {
		trailMap[1] = currentPosition
	} else {
		trailMap[len(trailMap)+1] = currentPosition
	}
	//fmt.Printf("Current position: %v\n", currentPosition)
	//fmt.Printf("Trail map: %v\n", trailMap)
}

func constructGrid(input string) ([][]int, [][]int) {
	grid := make([][]int, 0)
	trailHeads := make([][]int, 0)
	for r, row := range strings.Fields(input) {
		rowSlice := make([]int, 0)
		for c, col := range strings.Split(row, "") {
			rowSlice = append(rowSlice, convertStrToInt(col))
			if convertStrToInt(col) == 0 {
				trailHeads = append(trailHeads, []int{r, c, 0})
			}
		}
		grid = append(grid, rowSlice)
	}

	return grid, trailHeads
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
