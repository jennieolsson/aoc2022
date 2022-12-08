package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	allElfs := readFile()
	maxValues := findMaximumCalories(allElfs, 3)

	fmt.Println("Found max is")
	fmt.Println(maxValues)
}

func findMaximumCalories(allElfs []int, numbersOfMax int) []int {
	sort.Ints(allElfs)
	startIndex := len(allElfs) - (numbersOfMax)
	return allElfs[startIndex:len(allElfs)]
}

func readFile() []int {
	f, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	allElfs := make([]int, 0)
	currentElfSum := 0
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 && currentElfSum > 0 {
			allElfs = append(allElfs, currentElfSum)
			currentElfSum = 0
		} else {
			calorieCount, err := strconv.Atoi(line)

			if err == nil {
				currentElfSum += calorieCount
			}
		}
	}

	if currentElfSum > 0 {
		allElfs = append(allElfs, currentElfSum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return allElfs
}
