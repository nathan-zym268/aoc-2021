package main

import (
	"fmt"
	"lib"
	"strconv"
	"strings"
)

func firstPart(inputData []string) int {
	horizontal := 0
	depth := 0
	for _, value := range inputData {
		s := strings.Fields(value)
		direction := s[0]
		units, _ := strconv.Atoi(s[1])
		if direction == "forward" {
			horizontal += units
		}
		if direction == "down" {
			depth += units
		}
		if direction == "up" {
			depth -= units
		}

	}

	return depth * horizontal
}

func secondPart(inputData []string) int {
	horizontal := 0
	aim := 0
	depth := 0
	for _, value := range inputData {
		s := strings.Fields(value)
		direction := s[0]
		units, _ := strconv.Atoi(s[1])
		if direction == "forward" {
			horizontal += units
			depth += units * aim
		}
		if direction == "down" {
			aim += units
		}
		if direction == "up" {
			aim -= units
		}

	}

	return depth * horizontal
}
func main() {
	example := lib.ReadData("example/day2.txt")
	data := lib.ReadData("data/day2.txt")
	example_part1 := firstPart(example)
	example_part2 := secondPart(example)
	fmt.Println("Part 1 Example:", example_part1)
	fmt.Println("Part 2 Example: ", example_part2)
	result1 := firstPart(data)
	result2 := secondPart(data)
	fmt.Println("Answer for part 1: ", result1)
	fmt.Println("Answer for part 2: ", result2)
}
