package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read the 'input' file
	orderingRules, oerr := os.ReadFile("orderingRules")
	pageNumbers, perr := os.ReadFile("pageNumbers")
	if oerr != nil {
		fmt.Println(oerr)
		return
	}
	if perr != nil {
		fmt.Println(perr)
		return
	}

	fmt.Println("Files opened successfully")

	fmt.Println("Part 1:")
	part1(string(orderingRules), string(pageNumbers))
}

func part1(orderingRules string, pageNumbers string) {
	// generate a list of the updates
	updates := [][]int{}
	rawUpdates := strings.Split(pageNumbers, "\n")

	for i := 0; i < len(rawUpdates); i++ {
		// split each line into numbers
		numberStrings := strings.Split(rawUpdates[i], ",")
		nums := []int{}
		for j := 0; j < len(numberStrings); j++ {
			num, err := strconv.Atoi(numberStrings[j])
			if err != nil {
				fmt.Println("Error converting to int")
				return
			}
			nums = append(nums, num)
		}

		if len(nums)%2 == 0 {
			fmt.Println("Error: update has an even number of numbers. Index:", i)
			return
		}

		updates = append(updates, nums)
	}

	// generate a list of the rules, the rules are given in the format "a|b", where a and b are integers and a must appear before b
	rules := [][]int{}
	rawRules := strings.Split(orderingRules, "\n")

	for i := 0; i < len(rawRules); i++ {
		// split each line into numbers
		numberStrings := strings.Split(rawRules[i], "|")
		nums := []int{}
		for j := 0; j < len(numberStrings); j++ {
			num, err := strconv.Atoi(numberStrings[j])
			if err != nil {
				fmt.Println("Error converting to int")
				return
			}
			nums = append(nums, num)
		}

		rules = append(rules, nums)
	}

	// check each update against the rules to see if it is valid
	validUpdates := [][]int{}
	incorrectUpdates := [][]int{}
	for i := 0; i < len(updates); i++ {
		// check if the update is valid
		valid := true
		for j := 0; j < len(rules); j++ {
			if !valid {
				continue
			}
			ruleNum1 := rules[j][0]
			ruleNum2 := rules[j][1]

			for k := 0; k < len(updates[i])-1; k++ {
				if !valid {
					continue
				}

				// ignore the rule if the rule numbers aren't in the update
				if !containsBothNumbers(updates[i], ruleNum1, ruleNum2) {
					continue
				}

				index1 := getArrayIndex(updates[i], ruleNum1)
				index2 := getArrayIndex(updates[i], ruleNum2)

				if index1 > index2 {
					valid = false
				}
			}
		}
		if valid {
			validUpdates = append(validUpdates, updates[i])
		} else {
			incorrectUpdates = append(incorrectUpdates, updates[i])
		}
	}

	// sum the middle numbers of the valid updates
	sum := 0
	for i := 0; i < len(validUpdates); i++ {
		middleNum := getMiddleNumber(validUpdates[i])
		if middleNum != -1 {
			sum += middleNum
		} else {
			fmt.Println("Error: update has an even number of numbers. Index:", i)
		}
	}
	fmt.Println("Sum:", sum)
}
func containsBothNumbers(arr []int, num1 int, num2 int) bool {
	num1Found := false
	num2Found := false

	for i := 0; i < len(arr); i++ {
		if arr[i] == num1 {
			num1Found = true
		}
		if arr[i] == num2 {
			num2Found = true
		}
	}

	return num1Found && num2Found
}

func getArrayIndex(arr []int, num int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return i
		}
	}
	return -1
}

func getMiddleNumber(arr []int) int {
	if len(arr)%2 == 0 {
		return -1
	}
	return arr[len(arr)/2]
}
