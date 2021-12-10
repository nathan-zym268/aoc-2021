package main

import (
	"fmt"
	"strconv"

	"lib"
)

func str2intSlice(strSlice []string) []int {
	var result []int
	for _, value := range strSlice {
		intVal, _ := strconv.Atoi(value)
		result = append(result, intVal)
	}
	return result
}

func firstPart(inputData []string) int {
	result := 0
	intData := str2intSlice(inputData)
	for i := 0; i < len(intData)-1; i++ {
		if intData[i+1] > intData[i] {
			result++
		}

	}
	return result
}

func secondPart(inputData []string) int {
	intData := str2intSlice(inputData)
	result := 0
	for i := 0; i < len(intData)-3; i++ {
		x1 := intData[i] + intData[i+1] + intData[i+2]
		x2 := intData[i+1] + intData[i+2] + intData[i+3]
		if x1 < x2 {
			result++
		}

	}
	return result
}

func main() {
	example := lib.ReadData("example/day1.txt")
	data := lib.ReadData("data/day1.txt")
	example_part1 := firstPart(example)
	example_part2 := secondPart(example)
	fmt.Println("Part 1 Example:", example_part1)
	fmt.Println("Part 2 Example: ", example_part2)
	result1 := firstPart(data)
	result2 := secondPart(data)
	fmt.Println("Answer for part 1: ", result1)
	fmt.Println("Answer for part 2: ", result2)
}
