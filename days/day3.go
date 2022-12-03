package days

import (
	"fmt"
	"os"
	"bufio"
)

func priority(item int) int {
	if (item <= 'Z') {
		return item - 'A' + 27
	}

	return item - 'a' + 1
}

func priorityOfDuplicate(line string) int {
	compartement1 := NewSet()

	for _, item := range line[:len(line) / 2] {
		compartement1.Add(int(item))
	}

	for _, itemRune := range line[len(line) / 2:] {
		item := int(itemRune)
		if compartement1.Contains(item) {
			return priority(item)
		}
	}

	// unreachable
	return -1
}

func duplicatesInGroup(rucksacks [3]string) int {
	rucksack1Contents := NewSet()

	for _, item := range rucksacks[0] {
		rucksack1Contents.Add(int(item))
	}

	rucksacks1And2Duplicates := NewSet()

	for _, itemRune := range rucksacks[1] {
		item := int(itemRune)
		if rucksack1Contents.Contains(item) {
			rucksacks1And2Duplicates.Add(item)
		}
	}

	for _, itemRune := range rucksacks[2] {
		item := int(itemRune)
		if rucksacks1And2Duplicates.Contains(item) {
			return priority(item)
		}
	}

	// unreachable
	return -1
}

func (d *Day) Day3Part1(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	totalDuplicatePriorities := 0
	for scanner.Scan() {
		totalDuplicatePriorities += priorityOfDuplicate(scanner.Text())
	}

	fmt.Println("total priority duplicates:", totalDuplicatePriorities)
}

func (d *Day) Day3Part2(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sumOfBadgePriorities := 0
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()

		sumOfBadgePriorities += duplicatesInGroup([3]string{line1, line2, line3})
	}

	fmt.Println("Sum of badge priorities:", sumOfBadgePriorities)
}
