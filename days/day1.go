package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func sum(slice []int) int {
	var result int
	for _, num := range slice {
		result += num
	}

	return result
}

func (d *Day) Day1Part1(filePath string) {
	file, _ := os.Open(filePath)
	// close file at the end of this function
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	highestCalories := math.MinInt

	currentElf := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			highestCalories = max(currentElf, highestCalories)
			currentElf = 0
			continue
		}

		calories, _ := strconv.Atoi(scanner.Text())
		currentElf += calories
	}

	highestCalories = max(currentElf, highestCalories)

	fmt.Printf("Part1: %d\n", highestCalories)
}

func (d *Day) Day1Part2(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	topCalories := []int{math.MinInt, math.MinInt, math.MinInt}
	currentElf := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			topCalories, _ =
				insertSorted(topCalories, currentElf, intCompareFunc)
			topCalories = topCalories[1:]
			currentElf = 0
			continue
		}

		calories, _ := strconv.Atoi(scanner.Text())
		currentElf += calories
	}
	topCalories, _ = insertSorted(topCalories, currentElf, intCompareFunc)
	topCalories = topCalories[1:]

	fmt.Println("Top calores:", sum(topCalories))
}
