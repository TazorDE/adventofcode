package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
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

	reports := [][]int{}
	// split the file content into lines
	fmt.Println("Splitting the file content into lines...")
	lines := strings.Split(string(file), "\n")
	for i := 0; i < len(lines); i++ {
		// split the line into words, ignoring the empty words
		words := strings.Fields(lines[i])

		// convert each word to an integer
		intWords := []int{}
		for j := 0; j < len(words); j++ {
			num, err := strconv.Atoi(words[j])
			if err != nil {
				fmt.Println(err)
				return
			}
			intWords = append(intWords, num)
		}

		// append the integers to the reports list
		reports = append(reports, intWords)
	}
	fmt.Println("Part 1:")
	part1(reports)

}

func part1(reports [][]int) {
	safeReports := 0
	for i := 0; i < len(reports); i++ {
		levels := reports[i]

		if len(levels) < 2 {
			fmt.Println("Invalid input")
			return
		}

		isSafe := true

		if !checkConsistentChange(levels) {
			continue
		}

		for j := 0; j < len(levels)-1; j++ {
			if !checkDifference(levels[j], levels[j+1], 1, 3, false) {

				isSafe = false
				break
			}
		}

		if isSafe {
			safeReports++
		}
		continue
	}
	fmt.Println("Safe reports:", safeReports)
}

func checkDifference(num1 int, num2 int, minDiff int, maxDiff int, sameAllowed bool) bool {
	if !sameAllowed {
		// if number 1 is equal to number 2, fail the test
		if num1 == num2 {
			return false
		}
	}

	difference := num1 - num2
	// get the absolute value of the difference
	if difference < 0 {
		difference = -difference
	}

	if difference < minDiff || difference > maxDiff {
		return false
	}

	return true
}

func checkConsistentChange(levels []int) bool {
	// check if the levels are consistently increasing or decreasing
	consistentIncrease := true
	consistentDecrease := true

	// check if the levels are consistently increasing or decreasing
	for i := 0; i < len(levels)-1; i++ {
		if levels[i] < levels[i+1] {
			consistentDecrease = false
			continue
		}
		if levels[i] > levels[i+1] {
			consistentIncrease = false
			continue
		}
	}

	if !consistentIncrease && !consistentDecrease {
		return false
	}
	return true
}
