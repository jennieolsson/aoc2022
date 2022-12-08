package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	allElfs := readFile()
	max := findMaximumCalories(allElfs)

	fmt.Println(fmt.Sprintf("Found max is %d", max))
}

func findMaximumCalories(allElfs [][]int) int {
	max := 0

	for _, currentElf := range allElfs {
		currentElfSum := 0
		for _, calorieCount := range currentElf {
			currentElfSum += calorieCount
		}
		if currentElfSum > max {
			max = currentElfSum
		}
	}
	return max
}

func readFile() [][]int {
	allElfs := make([][]int, 0)

	f, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	currentElf := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" && len(currentElf) > 0 {
			allElfs = append(allElfs, currentElf)
			currentElf = make([]int, 0)
		} else {
			calorieCount, err := strconv.Atoi(line)

			if err == nil {
				currentElf = append(currentElf, calorieCount)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return allElfs
}
