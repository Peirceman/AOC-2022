package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (d *Day) Day5Part1(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		if (scanner.Text()[1] == '1') {
			break
		}

		lines = append(lines, scanner.Text())
	}

	crateStacks := make([]*doublyLinkedList[int], (len(lines[0])+1)/4)

	for i := 0; i < len(crateStacks); i++ {
		crateStacks[i] = NewDoublyLinkedList[int]()
	}

	lineLength := len(lines[0])
	for _, line := range lines {

		rowIndex := 0
		for j := 1; j < lineLength; j += 4 {
			if (line[j] == ' ') {
				rowIndex++
				continue
			}

			crateStacks[rowIndex].InsertConstant(0, int(line[j]))
			rowIndex++
		}
	}

	// printStacks(crateStacks)

	scanner.Scan() // skip empty line

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		count,_ := strconv.Atoi(parts[1])
		originIndex,_ := strconv.Atoi(parts[3])
		destinationIndex,_ := strconv.Atoi(parts[5])
		destinationIndex--
		origin := crateStacks[originIndex - 1]
		destination := crateStacks[destinationIndex]

		crateStacks[destinationIndex] = destination.Append(origin.SubList(origin.length - count, origin.length).Reverse())
		origin.MaxLength(origin.length - count)
		// fmt.Printf("count: %d, originIndex: %d, destinationIndex: %d\n", count, originIndex, destinationIndex)
		// printStacks(crateStacks)
	}

	fmt.Print("top crate names: ")
	for _,stack := range crateStacks {
		fmt.Print(string(*stack.end.value))
	}
	fmt.Println()
}

func printStacks(crateStacks []*doublyLinkedList[int]) {
	maxHeight := 0
	for _, stack := range crateStacks {
		if stack.length > maxHeight {
			maxHeight = stack.length
		}
	}

	for i := maxHeight - 1; i > -1; i-- {
		for _, crateStack := range crateStacks {
			if (i >= crateStack.length) {
				fmt.Print(" ")
				continue
			}
			numPtr,_ := crateStack.Get(i)
			fmt.Print(string(*numPtr), "")
		}
		fmt.Println()
	}
}

func (d *Day) Day5Part2(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		if (scanner.Text()[1] == '1') {
			break
		}

		lines = append(lines, scanner.Text())
	}

	crateStacks := make([]*doublyLinkedList[int], (len(lines[0])+1)/4)

	for i := 0; i < len(crateStacks); i++ {
		crateStacks[i] = NewDoublyLinkedList[int]()
	}

	lineLength := len(lines[0])
	for _, line := range lines {

		rowIndex := 0
		for j := 1; j < lineLength; j += 4 {
			if (line[j] == ' ') {
				rowIndex++
				continue
			}

			crateStacks[rowIndex].InsertConstant(0, int(line[j]))
			rowIndex++
		}
	}

	// printStacks(crateStacks)

	scanner.Scan() // skip empty line

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		count,_ := strconv.Atoi(parts[1])
		originIndex,_ := strconv.Atoi(parts[3])
		destinationIndex,_ := strconv.Atoi(parts[5])
		destinationIndex--
		origin := crateStacks[originIndex - 1]
		destination := crateStacks[destinationIndex]

		crateStacks[destinationIndex] = destination.Append(origin.SubList(origin.length - count, origin.length))
		origin.MaxLength(origin.length - count)
		// fmt.Printf("count: %d, originIndex: %d, destinationIndex: %d\n", count, originIndex, destinationIndex)
		// printStacks(crateStacks)
	}

	fmt.Print("top crate names: ")
	for _,stack := range crateStacks {
		fmt.Print(string(*stack.end.value))
	}
	fmt.Println()
}
