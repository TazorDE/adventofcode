package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	// check if os.Args contains the value '--test'
	testmode := slices.Contains(os.Args, "--test")
	filename := "input"
	if testmode {
		filename = "testinput"
	}

	// read the 'input' file
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File opened successfully")

	fmt.Println("Part 1:")
	part1(string(file))

	fmt.Println("Part 2:")
	part2(string(file))
}

func part1(memoryContent string) {
	multiplicationPattern := regexp.MustCompile(`mul\(\d+,\d+\)`)

	matches := multiplicationPattern.FindAllString(memoryContent, -1)
	numericMatcher := regexp.MustCompile(`\d+`)

	sum := 0

	for i := 0; i < len(matches); i++ {
		// get the two numbers
		numbers := numericMatcher.FindAllString(matches[i], -1)
		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Error converting to int")
			return
		}

		// multiply the numbers
		sum += num1 * num2
	}

	fmt.Println("Sum:", sum)
}

func part2(memoryContent string) {
	doDontPattern := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	statements := doDontPattern.FindAllString(memoryContent, -1)

	if statements == nil {
		fmt.Println("No statements found")
		return
	}

	enabled := true
	sum := 0

	for i := 0; i < len(statements); i++ {
		// enable or disable the multiplication
		if statements[i] == "do()" {
			enabled = true
			continue
		} else if statements[i] == "don't()" {
			enabled = false
			continue
		}

		// get the two numbers
		numbers := regexp.MustCompile(`\d+`).FindAllString(statements[i], -1)

		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Error converting to int")
			return
		}

		// multiply the numbers if enabled
		if enabled {
			sum += num1 * num2
		} else {
			continue
		}
	}
	fmt.Println("Sum:", sum)
}
