package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"slices"
	"sort"
	"strconv"
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
	//exampleInput := "12345"
	//fmt.Printf("input: %s\n", input)
	checksum := diskOptimization(input)

	//fmt.Printf("checksum: %v\n", checksum)
	return checksum
}

func solvePart2(input string) int {
	checksum := diskOptimizationPart2(input)
	//fmt.Printf("checksum: %v\n", checksum)
	return checksum
}

func diskOptimization(input string) int {
	disk, _, _, _, _ := initializeDisk(input)
	//fmt.Printf("disk: %v\n", disk)

	for {
		freeSpacePosition := findNextFreeSpacePosition(disk)
		if freeSpacePosition == -1 {
			break
		}

		//Find rightmost file block after the free space
		filePosition := -1
		for i := len(disk) - 1; i > freeSpacePosition; i-- {
			if disk[i] != -1 {
				filePosition = i
				break
			}
		}

		if filePosition == -1 {
			break
		}

		// Move the file block
		disk[freeSpacePosition] = disk[filePosition]
		disk[filePosition] = -1
		//fmt.Printf("File Position: %v, Free Space Position: %v\n", filePosition, freeSpacePosition)
		//break
	}
	//fmt.Printf("new disk: %v\n", disk)

	checksum := 0
	for index, block := range disk {
		if block != -1 {
			checksum += block * index
		}
	}
	//fmt.Printf("checksum: %v\n", checksum)
	return checksum
}

func diskOptimizationPart2(input string) int {
	disk, fileSizes, _, _, fileIds := initializeDisk(input)
	//fmt.Printf("disk: %v\n", disk)
	//fmt.Printf("File Sizes: %v\n", fileSizes)
	//fmt.Printf("Space Sizes: %v\n", spaceSizes)
	//fmt.Printf("Space Indices: %v\n", spaceIndices)

	// Process files in decreasing ID order
	for _, currentId := range fileIds {
		if fileSizes[currentId] == nil {
			continue
		}
		size := fileSizes[currentId][0][1]
		fileStartIndex := fileSizes[currentId][0][0]
		//fmt.Printf("Size: %v, Start Index: %v\n", size, fileStartIndex)

		// Find leftmost valid position
		bestPosition := -1
		spaceCount := 0
		for i, block := range disk {
			if block == -1 {
				spaceCount += 1
				if spaceCount >= size {
					bestPosition = i - size + 1
					break
				}
			} else {
				spaceCount = 0
			}
		}

		if bestPosition != -1 && bestPosition < fileStartIndex {
			// Clear old position
			for i := fileStartIndex; i < fileStartIndex+size; i++ {
				disk[i] = -1
			}

			//	Fill new position
			for i := bestPosition; i < bestPosition+size; i++ {
				disk[i] = currentId
			}
		}
	}

	//fmt.Printf("Disk: %v\n", disk)
	//fmt.Printf("File Sizes: %v\n", fileSizes)
	//fmt.Printf("File Ids: %v\n", fileIds)
	//fmt.Printf("Space Indices: %v\n", spaceIndices)
	//fmt.Printf("Space Sizes: %v\n", spaceSizes)

	checksum := 0
	for index, block := range disk {
		if block != -1 {
			checksum += block * index
		}
	}

	return checksum
}

func findNextFreeSpacePosition(disk []int) int {
	freeSpacePosition := -1

	for i, block := range disk {
		if block == -1 {
			freeSpacePosition = i
			break
		}
	}
	return freeSpacePosition
}

func initializeDisk(input string) ([]int, map[int][][]int, map[int][]int, []int, []int) {
	var disk []int
	fileId := 0
	fileMap := make(map[int][][]int)
	spaceSizes := make(map[int][]int)
	var spaceIndices []int
	var fileIds []int
	for i := 0; i < len(input); i++ {
		diskBlock := convertStrToInt(string(input[i]))
		//fmt.Printf("diskBlock: %v\n", diskBlock)
		if i%2 == 0 {
			fileIds = append(fileIds, fileId)
			var fileSlice []int
			fileSlice = append(fileSlice, len(disk), diskBlock)
			fileMap[fileId] = append(fileMap[fileId], fileSlice)
			for j := 0; j < diskBlock; j++ {
				disk = append(disk, fileId)
			}

			fileId += 1
		} else {
			if diskBlock != 0 {
				spaceSizes[len(disk)] = append(spaceSizes[len(disk)], diskBlock)
				spaceIndices = append(spaceIndices, len(disk))
				for j := 0; j < diskBlock; j++ {
					disk = append(disk, -1)
				}
			}
		}
	}
	slices.Sort(spaceIndices)
	// Sort the slice in descending order
	sort.Slice(fileIds, func(i, j int) bool {
		return fileIds[i] > fileIds[j] // Descending order
	})
	return disk, fileMap, spaceSizes, spaceIndices, fileIds
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
