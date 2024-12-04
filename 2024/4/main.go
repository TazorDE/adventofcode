package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
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

	fmt.Println("Part 1:")
	part1(string(file))

	fmt.Println("Part 2:")
	// part2(string(file))
}

func part1(wordSearch string) {
	// split the input into lines
	lines := strings.Split(wordSearch, "\n")
	arr2d := [][]string{}

	// This word search allows words to be horizontal, vertical, diagonal, written backwards, or even overlapping other words.
	// if the word is found, increment a counter
	xmasRegex := regexp.MustCompile(`XMAS`)

	// create a 2D array to store the letters
	for i := 0; i < len(lines); i++ {
		// split each line into letters
		arr2d = append(arr2d, strings.Split(lines[i], ""))
	}

	// find horizontal matches
	horizontalCounter := 0
	for i := 0; i < len(lines); i++ {
		// if lines[i] contains "XMAS" increase the counter
		matches := xmasRegex.FindAllString(lines[i], -1)
		reverseMatches := xmasRegex.FindAllString(reverseString(lines[i]), -1)

		if matches != nil {
			horizontalCounter += len(matches)
		}
		if reverseMatches != nil {
			horizontalCounter += len(reverseMatches)
		}
	}

	// find vertical matches
	verticalCounter := 0
	for i := 0; i < len(arr2d); i++ {
		// create a string from the column
		column := ""
		for j := 0; j < len(arr2d); j++ {
			column += arr2d[j][i]
		}

		// if column contains "XMAS" increase the counter
		matches := xmasRegex.FindAllString(column, -1)
		reverseMatches := xmasRegex.FindAllString(reverseString(column), -1)

		if matches != nil {
			verticalCounter += len(matches)
		}
		if reverseMatches != nil {
			verticalCounter += len(reverseMatches)
		}
	}

	// find diagonal matches
	diagonalCounter := 0

	// diagonals to right
	rightDiagonals := []string{}
	for i := 0; i < len(arr2d); i++ {
		diagonal := ""
		for j := 0; j < len(arr2d)-i; j++ {
			diagonal += arr2d[j+i][j]
		}
		rightDiagonals = append(rightDiagonals, diagonal)
		diagonal = ""
		for j := 0; j < len(arr2d)-i; j++ {
			diagonal += arr2d[j][j+i]
		}
		rightDiagonals = append(rightDiagonals, diagonal)
	}
	leftDiagonals := []string{}
	for i := 0; i < len(arr2d); i++ {
		diagonal := ""
		for j := 0; j < len(arr2d)-i; j++ {
			diagonal += arr2d[j+i][len(arr2d)-j-1]
		}
		leftDiagonals = append(leftDiagonals, diagonal)
		diagonal = ""
		for j := 0; j < len(arr2d)-i; j++ {
			diagonal += arr2d[j][len(arr2d)-j-i-1]
		}
		leftDiagonals = append(leftDiagonals, diagonal)
	}

	diagonals := append(rightDiagonals, leftDiagonals...)

	for i := 0; i < len(diagonals); i++ {
		matches := xmasRegex.FindAllString(diagonals[i], -1)
		reverseMatches := xmasRegex.FindAllString(reverseString(diagonals[i]), -1)

		if matches != nil {
			diagonalCounter += len(matches)
		}
		if reverseMatches != nil {
			diagonalCounter += len(reverseMatches)
		}
	}

	fmt.Println("Horizontal matches:", horizontalCounter)
	fmt.Println("Vertical matches:", verticalCounter)
	fmt.Println("Diagonal matches:", diagonalCounter)
	fmt.Println("Total matches:", horizontalCounter+verticalCounter+diagonalCounter)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// swap the letters
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
