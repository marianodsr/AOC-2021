package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := parseInput("input.txt")
	solveStepOne(input)
	solveStepTwo(input)
}

func parseInput(filepath string) []string {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		panic("could not parse input")
	}

	rawText := string(bytes)

	splitted := strings.Split(rawText, "\n")

	return splitted

}

func solveStepOne(input []string) {
	var x int = 0
	var depth int = 0

	for _, s := range input {
		words := strings.Split(s, " ")
		n, _ := strconv.Atoi(words[1])

		switch words[0] {
		case "forward":
			x += n
			break
		case "up":
			depth -= n
			break
		case "down":
			depth += n
			break
		default:
			panic("invalid input")
		}

	}

	fmt.Println("Result: ", depth*x)
}

func solveStepTwo(input []string) {

	var x int = 0
	var depth int = 0
	var aim int = 0

	for _, s := range input {
		words := strings.Split(s, " ")
		n, _ := strconv.Atoi(words[1])

		switch words[0] {
		case "forward":
			x += n
			depth = depth + (aim * n)
			break
		case "up":
			aim -= n
			break
		case "down":
			aim += n
			break
		default:
			panic("invalid input")
		}

	}
	fmt.Println("Result: ", depth*x)

}
