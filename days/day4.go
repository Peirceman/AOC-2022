package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (d *Day) Day4Part1(filePath string) {
	file,_ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	containedCount := 0
	for scanner.Scan() {
		elf1Str := strings.Split(strings.Split(scanner.Text(), ",")[0], "-")
		elf2Str := strings.Split(strings.Split(scanner.Text(), ",")[1], "-")
		elf1Min, _ := strconv.Atoi(elf1Str[0])
		elf1Max, _ := strconv.Atoi(elf1Str[1])
		elf2Min, _ := strconv.Atoi(elf2Str[0])
		elf2Max, _ := strconv.Atoi(elf2Str[1])

		if (elf1Min <= elf2Min && elf1Max >= elf2Max) || (elf2Min <= elf1Min && elf2Max >= elf1Max) {
			containedCount++;
		}
	}

	fmt.Println("amount of fully contained regions pairs:", containedCount)
}

func (d *Day) Day4Part2(filePath string) {
	file,_ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	overlapCount := 0
	for scanner.Scan() {
		elf1Str := strings.Split(strings.Split(scanner.Text(), ",")[0], "-")
		elf2Str := strings.Split(strings.Split(scanner.Text(), ",")[1], "-")
		elf1Min, _ := strconv.Atoi(elf1Str[0])
		elf1Max, _ := strconv.Atoi(elf1Str[1])
		elf2Min, _ := strconv.Atoi(elf2Str[0])
		elf2Max, _ := strconv.Atoi(elf2Str[1])

		if (elf1Max >= elf2Min && elf1Min <= elf2Max) || (elf2Max >= elf1Min && elf2Min <= elf1Max) {
			overlapCount++
		}
	}

	fmt.Println("amount of overlaping pairs:", overlapCount)
}
