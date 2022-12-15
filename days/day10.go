package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func (d *Day) Day10Part2(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	screen := [6*40]bool{}
	CRT_Pos := 0
	X := 1
	for scanner.Scan() {
		if X - 1 <= CRT_Pos % 40 && CRT_Pos % 40  <= X + 1 {
			screen[CRT_Pos] = true
		}

		CRT_Pos++
		if scanner.Text() == "noop" {
			continue
		}

		if X - 1 <= CRT_Pos % 40 && CRT_Pos % 40 <= X + 1 {
			screen[CRT_Pos] = true
		}

		CRT_Pos++
		numToAdd, _ := strconv.Atoi(scanner.Text()[5:])
		X += numToAdd
		if X - 1 <= CRT_Pos % 40 && CRT_Pos % 40 <= X + 1 {
			screen[CRT_Pos] = true
		}
	}

	if X - 1 <= CRT_Pos % 40 && CRT_Pos % 40 <= X + 1 {
		screen[CRT_Pos] = true
	}

	for i := 0; i < 6; i++ {
		currentLine := screen[40*i:40*(i+1)]
		for _, pixel := range currentLine {
			if pixel {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (d *Day) Day10Part1(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sumOfSignalStrengths := 0
	checkedCyclesCount := 0
	cycle := 1
	X := 1
	for i := 0; scanner.Scan(); i++ {
		if cycle - 20 == checkedCyclesCount * 40 {
			sumOfSignalStrengths += X * cycle
			checkedCyclesCount++
		}

		cycle++
		if scanner.Text() == "noop" {
			continue
		}

		if cycle - 20 == checkedCyclesCount * 40 {
			sumOfSignalStrengths += X * cycle
			checkedCyclesCount++
		}

		cycle++
		numToAdd, _ := strconv.Atoi(scanner.Text()[5:])
		X += numToAdd
		if cycle - 20 == checkedCyclesCount * 40 {
			sumOfSignalStrengths += X * cycle
			checkedCyclesCount++
		}
	}

	if cycle - 20 == checkedCyclesCount * 40 {
		sumOfSignalStrengths += X * cycle
		checkedCyclesCount++
	}

	fmt.Println("sum of signal strengths", sumOfSignalStrengths)
}
