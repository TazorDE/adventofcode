package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
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
	firstList := []int{}
	secondList := []int{}

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

		// append the integers to the first or second list
		firstList = append(firstList, intWords[0])
		secondList = append(secondList, intWords[1])
	}

	// check that both lists have the same length
	fmt.Println("First list length:", len(firstList))
	fmt.Println("Second list length:", len(secondList))
	if len(firstList) != len(secondList) {
		fmt.Println("The lists have different lengths")
		return
	}

	fmt.Println("Part 1:")
	part1(firstList, secondList)

	fmt.Println("Part 2:")
	part2(firstList, secondList)
}

func part1(firstList []int, secondList []int) {

	// sort both the first and second list
	fmt.Println("Sorting the lists...")
	sort.Ints(firstList)
	sort.Ints(secondList)

	sum := 0

	// calculate the absolute distance between corresponding elements in the two lists
	fmt.Println("Calculating the distance between corresponding elements in the lists...")
	for i := 0; i < len(firstList); i++ {
		tempsum := firstList[i] - secondList[i]
		if tempsum < 0 {
			tempsum *= -1
		}
		sum += tempsum
	}

	// print the sum
	fmt.Println("The sum is:", sum)
}

func part2(firstList []int, secondList []int) {
	// calculate a similarity score between the two lists
	score := 0

	// for each value in the first list check how many times it appears in the second list
	fmt.Println("Calculating the similarity score...")
	for i := 0; i < len(firstList); i++ {
		appearances := 0
		for j := 0; j < len(secondList); j++ {
			if firstList[i] == secondList[j] {
				appearances++
			}
		}
		// multiply the number of appearances by the value in the first list and add it to the score
		score += (firstList[i] * appearances)
	}

	// print the score
	fmt.Println("The similarity score is:", score)
}
