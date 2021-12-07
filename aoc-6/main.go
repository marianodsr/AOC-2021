package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	args := os.Args[1:]

	days := args[0]

	ndays, _ := strconv.Atoi(days)

	input := parseInput("input.txt")

	table := dumpTable(input)
	table = timeTravelToAmountOfDays(ndays, table)

	fmt.Printf("RESULT AFTER %d DAYS:\n", ndays)
	fmt.Println(getAmountOfFishs(table))
}

func timeTravelToAmountOfDays(days int, table map[int]int) map[int]int {
	for i := 0; i < days; i++ {
		table = nextDay(table)
	}

	return table
}

func getAmountOfFishs(table map[int]int) int {
	var count int
	for _, v := range table {
		count += v
	}
	return count
}

func nextDay(table map[int]int) map[int]int {
	newTable := make(map[int]int)
	for k, v := range table {
		//gives birth to new lanternfish
		if k == 0 {
			newTable[8] = v
			newTable[6] += v
			continue
		}
		newTable[k-1] += v

	}

	return newTable
}

func dumpTable(input []int) map[int]int {

	table := make(map[int]int)

	for _, n := range input {
		table[n] = table[n] + 1
	}

	return table
}

func parseInput(filepath string) []int {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		panic("could not open input file")
	}

	rawText := string(bytes)

	splitted := strings.Split(rawText, ",")

	var nArray []int

	for _, char := range splitted {
		n, _ := strconv.Atoi(char)
		nArray = append(nArray, n)
	}

	return nArray

}
