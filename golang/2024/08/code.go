package main

import (
	"fmt"
	"github.com/jpillora/puzzler/harness/aoc"
	"math"
	"sort"
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
	lines := (strings.Split(input, "\n"))
	return findAntinodes(lines)
}

func solvePart2(input string) int {
	lines := (strings.Split(input, "\n"))
	return findAntinodesInLine(lines)
}

func findAntinodes(lines []string) int {
	solution := 0
	matrix := make([][]string, 0)
	antennaMap := make(map[string][][]int)
	for _, line := range lines {
		if len(line) > 0 {
			matrix = append(matrix, strings.Split(line, ""))
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != "." {
				antennaMap[matrix[i][j]] = append(antennaMap[matrix[i][j]], []int{i, j})
			}
		}
	}

	for _, v := range antennaMap {
		for i := range v {
			for j := range v {
				if i != j {
					x := v[i][0] + (v[i][0] - v[j][0])
					y := v[i][1] + (v[i][1] - v[j][1])
					isValid := x >= 0 && y >= 0 && x < len(matrix) && y < len(matrix[0])
					if isValid {
						matrix[x][y] = "#"
					}
				}
			}
		}
	}
	for _, row := range matrix {
		for _, e := range row {
			if e == "#" {
				solution++
			}
		}
	}
	//fmt.Printf("Part One: %d\n", solution)
	return solution
}

func findAntinodesInLine(lines []string) int {
	solution := 0
	matrix := make([][]string, 0)
	antennaMap := make(map[string][][]int)
	for _, line := range lines {
		if len(line) > 0 {
			matrix = append(matrix, strings.Split(line, ""))
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != "." {
				antennaMap[matrix[i][j]] = append(antennaMap[matrix[i][j]], []int{i, j})
			}
		}
	}

	for _, v := range antennaMap {
		for i := range v {
			for j := range v {
				matrix[v[i][0]][v[i][1]] = "#"
				if i != j {
					x := v[i][0] + (v[i][0] - v[j][0])
					y := v[i][1] + (v[i][1] - v[j][1])
					isValid := x >= 0 && y >= 0 && x < len(matrix) && y < len(matrix[0])
					for isValid {
						if isValid {
							matrix[x][y] = "#"
						}
						x += (v[i][0] - v[j][0])
						y += (v[i][1] - v[j][1])
						isValid = x >= 0 && y >= 0 && x < len(matrix) && y < len(matrix[0])
					}
				}
			}
		}
	}
	for _, row := range matrix {
		for _, e := range row {
			if e == "#" {
				solution++
			}
		}
	}
	//fmt.Printf("Part Two: %d\n", solution)
	return solution

}

func removeDuplicates(input []int) []int {
	seen := make(map[int]bool) // Map to track seen elements
	result := []int{}          // Slice to store unique values

	for _, val := range input {
		if !seen[val] {
			seen[val] = true
			result = append(result, val) // Add value to the result if not seen
		}
	}
	sort.Ints(result)
	return result
}

func calculateAntinodeLocations(antennas map[string][][]int, maxRow int, maxCol int) map[string][][]int {
	// this approaach missed the point 7,0 in the example.
	// This resulted in a final answer of 205 instead of 254 for part 1!!
	antinodes := make(map[string][][]int)
	for antennaName, antennaLocations := range antennas {
		for i := 0; i < len(antennaLocations)-1; i++ {
			currentAntennaLocation := antennaLocations[i]
			otherAntennaLocations := antennaLocations[i+1:]
			if len(otherAntennaLocations) == 0 {
				break
			}

			for _, otherAntennaLocation := range otherAntennaLocations {
				// calculate slope
				gradient := float64(otherAntennaLocation[1]-currentAntennaLocation[1]) / float64(otherAntennaLocation[0]-currentAntennaLocation[0])
				yIntercept := float64(otherAntennaLocation[1]) - gradient*float64(otherAntennaLocation[0])
				//fmt.Printf("Line connecting [%v, %v] and [%v, %v]: y = %vx + %v\n", currentAntennaLocation[0], currentAntennaLocation[1], otherAntennaLocation[0], otherAntennaLocation[1], gradient, yIntercept)

				for row := 0; row <= maxRow; row++ {
					//if row == 7 {
					//	fmt.Printf("Row: %v\n", row)
					//}
					//fmt.Printf("Row: %v, Col: %v\n", row, col)
					//if row == currentAntennaLocation[0] || row == otherAntennaLocation[0] {
					//	continue
					//}

					col := gradient*float64(row) + yIntercept
					//if row == 7 {
					//	fmt.Printf("Row: %v, Col: %v\n", row, col)
					//}
					if col < 0 || col > float64(maxCol) || col != float64(int(col)) {
						continue
					}

					if (row == currentAntennaLocation[0] && int(col) == currentAntennaLocation[1]) || (row == otherAntennaLocation[0] && int(col) == otherAntennaLocation[1]) {
						continue
					}
					//fmt.Printf("Row: %v, Col: %v\n", row, col)

					// distance between currentAntennaLocation and [row, col]
					// s = sqrt((x2 - x1)^2 + (y2-y1)^2)
					xChange := float64(row) - float64(currentAntennaLocation[0])
					xChangeSquared := xChange * xChange
					yChange := col - float64(currentAntennaLocation[1])
					yChangeSquared := yChange * yChange
					distanceCurrentAntenna := math.Pow(math.Abs((xChangeSquared - yChangeSquared)), 0.5)
					//fmt.Printf("Distance Antenna 1: %v, x2-x1: %v, (x2-x1)^2: %v, y2-y1: %v, (y2-y1)^2: %v\n", distanceCurrentAntenna, xChange, xChangeSquared, yChange, yChangeSquared)
					xChange2 := float64(row) - float64(otherAntennaLocation[0])
					xChangeSquared2 := xChange2 * xChange2
					yChange2 := col - float64(otherAntennaLocation[1])
					yChangeSquared2 := yChange2 * yChange2
					distanceOtherAntenna := math.Pow(math.Abs((xChangeSquared2 - yChangeSquared2)), 0.5)

					if row == 7 && col == 0 {
						fmt.Printf("Distance Antenna 1: %v, Distance Antenna 2: %v\n", distanceCurrentAntenna, distanceOtherAntenna)
					}
					if distanceCurrentAntenna == 2*distanceOtherAntenna || distanceOtherAntenna == 2*distanceCurrentAntenna {
						//if (distanceCurrentAntenna == 2*distanceOtherAntenna || distanceOtherAntenna == 2*distanceCurrentAntenna) && distanceCurrentAntenna != 0 && distanceOtherAntenna != 0 {
						if distanceCurrentAntenna == 0 || distanceOtherAntenna == 0 {
							//fmt.Printf("Zero Distance!!!!!\n")
							//fmt.Printf("Potential Antinode: [%v, %v]\n", row, col)
							//fmt.Printf("Distance Antenna 1: %v, Distance Antenna 2: %v\n", distanceCurrentAntenna, distanceOtherAntenna)

							if (math.Abs(float64(row)-float64(currentAntennaLocation[0])) == 1 && math.Abs(float64(col)-float64(currentAntennaLocation[1])) == 1) || (math.Abs(float64(row)-float64(otherAntennaLocation[0])) == 1 && math.Abs(float64(col)-float64(otherAntennaLocation[1])) == 1) {
								antinodes[antennaName] = append(antinodes[antennaName], []int{row, int(col)})
								//fmt.Printf("Antinode: [%v, %v]\n", row, col)
							}
							//antinodes[antennaName] = append(antinodes[antennaName], []int{row, int(col)})

						} else {
							//fmt.Printf("Distance Antenna 1: %v, Distance Antenna 2: %v\n", distanceCurrentAntenna, distanceOtherAntenna)
							//fmt.Printf("Found an Antinode: [%v, %v]\n", row, col)
							antinodes[antennaName] = append(antinodes[antennaName], []int{row, int(col)})
						}

					}
				}
			}
		}
	}
	return antinodes
}

func createGrid(input string) ([][]string, map[string][][]int, int, int) {
	var grid [][]string
	antennas := make(map[string][][]int)
	for row, line := range strings.Split(input, "\n") {
		var gridLine []string
		for col, char := range line {
			gridLine = append(gridLine, string(char))
			if char != '.' {
				antennas[string(char)] = append(antennas[string(char)], []int{row, col})
			}
		}
		grid = append(grid, gridLine)
	}
	maxRow := len(grid) - 1
	maxCol := len(grid[0]) - 1
	return grid, antennas, maxRow, maxCol
}

func printGrid(grid [][]string) {
	for i, line := range grid {
		fmt.Printf("Row %d: %v\n", i, line)
	}
}

func printMap(myMap map[string][][]int) {
	// Loop through the map keys and values
	for key, value := range myMap {
		fmt.Printf("Antenna: %v, Locations: %v\n", key, value)
	}
}
