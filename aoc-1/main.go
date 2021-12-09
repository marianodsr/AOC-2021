package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("input.txt")
	getStepOneDepth(input)
	getStepTwoDepth(input)
}

func parseInput(filepath string) []int {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		panic("could not parse input")
	}

	rawText := string(bytes)

	splitted := strings.Split(rawText, "\n")

	var nArray []int
	for _, char := range splitted {
		n, err := strconv.Atoi(char)
		if err != nil {
			panic("could not parse input")
		}
		nArray = append(nArray, n)
	}

	return nArray

}

func getStepTwoDepth(input []int) {
	var count int = 0
	var previous int = 0
	for i, n := range input {
		if (i + 2) > (len(input) - 1) {
			break
		}
		current := n + input[i+1] + input[i+2]
		if i == 0 {
			previous = current
			continue
		}
		if current > previous {
			count++
		}
		previous = current
	}

	fmt.Printf("\nDEPTH INCREASED %d amount of times", count)
}

func getStepOneDepth(input []int) {
	var count int = 0
	for i, n := range input {
		if i == 0 {
			continue
		}
		if n > input[i-1] {
			count++
		}
	}

	fmt.Printf("\nDEPTH INCREASED %d amount of times", count)
}
