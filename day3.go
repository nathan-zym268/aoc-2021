package main

import (
	"fmt"
	"lib"
	"strconv"
)

func firstPart(inputData []string) int64 {
	string_len := len([]rune(inputData[0]))
	input_len := len(inputData)
	count_1s := make([]int, string_len)
	for _, value1 := range inputData {

		for i, value2 := range []rune(value1) {
			// string type and rune type :(
			s, _ := strconv.Atoi(string(value2))
			count_1s[i] += s
		}
	}
	var gamma_string string
	var epsilon_string string
	for _, value := range count_1s {
		if value > input_len/2 {
			gamma_string += "1"
			epsilon_string += "0"
		} else {
			gamma_string += "0"
			epsilon_string += "1"
		}
	}
	gamma, err := strconv.ParseInt(gamma_string, 2, 64)
	epsilon, err := strconv.ParseInt(epsilon_string, 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	return gamma * epsilon
}

func secondPart(inputData []string) int64 {
	// make the input string into a int slice
	var intData [][]int
	for _, value1 := range inputData {
		var intSlice []int
		for _, value2 := range []rune(value1) {
			s, _ := strconv.Atoi(string(value2))
			intSlice = append(intSlice, s)
		}
		intData = append(intData, intSlice)
	}
	o2Rating := intData
	for i := range intData[0] {
		if len(o2Rating) == 1 {
			break
		}
		sum1 := 0
		for _, value := range o2Rating {
			sum1 += value[i]
		}
		var newO2Rating [][]int
		if float64(sum1) >= float64(len(o2Rating))/2 {
			for _, value := range o2Rating {
				if value[i] == 1 {
					newO2Rating = append(newO2Rating, value)
				}
			}
		} else {
			for _, value := range o2Rating {
				if value[i] == 0 {
					newO2Rating = append(newO2Rating, value)
				}
			}

		}
		o2Rating = newO2Rating
	}
	co2Rating := intData
	for i := range intData[0] {
		if len(co2Rating) == 1 {
			break
		}
		sum1 := 0
		for _, value := range co2Rating {
			sum1 += value[i]
		}
		var newCO2Rating [][]int
		if float64(sum1) < float64(len(co2Rating))/2 {
			for _, value := range co2Rating {
				if value[i] == 1 {
					newCO2Rating = append(newCO2Rating, value)
				}
			}
		} else {
			for _, value := range co2Rating {
				if value[i] == 0 {
					newCO2Rating = append(newCO2Rating, value)
				}
			}

		}
		co2Rating = newCO2Rating
	}
	o2_string := ""
	for _, value := range o2Rating[0] {
		o2_string += strconv.Itoa(value)
	}
	co2_string := ""
	for _, value := range co2Rating[0] {
		co2_string += strconv.Itoa(value)
	}
	o2, _ := strconv.ParseInt(o2_string, 2, 64)
	co2, _ := strconv.ParseInt(co2_string, 2, 64)

	return o2 * co2
}

func main() {
	example := lib.ReadData("example/day3.txt")
	data := lib.ReadData("data/day3.txt")
	example_part1 := firstPart(example)
	example_part2 := secondPart(example)
	fmt.Println("Part 1 Example:", example_part1)
	fmt.Println("Part 2 Example: ", example_part2)
	result1 := firstPart(data)
	result2 := secondPart(data)
	fmt.Println("Answer for part 1: ", result1)
	fmt.Println("Answer for part 2: ", result2)
}
