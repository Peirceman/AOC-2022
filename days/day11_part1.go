package days

import (
	"bufio"
	"bytes"
	"io"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type monkey_part1 struct {
	worryLevels []int
	opperation func(int) int
	divisibilityTest int
	ifTrueIndex, ifFalseIndex int // index of the monkey to throw to
	ifTrue, ifFalse *monkey_part1
	inspectedItems int
}

func (this *monkey_part1) round() {
	for _, worryLevel := range this.worryLevels {
		worryLevel = this.opperation(worryLevel)
		worryLevel /= 3
		if worryLevel % this.divisibilityTest == 0 {
			this.ifTrue.worryLevels =
					append(this.ifTrue.worryLevels, worryLevel)
		} else {
			this.ifFalse.worryLevels =
					append(this.ifFalse.worryLevels, worryLevel)
		}
	}

	this.inspectedItems += len(this.worryLevels)
	this.worryLevels = make([]int, 0, cap(this.worryLevels))
}

func (this *monkey_part1) readIndexes(monkeys []*monkey_part1) {
	this.ifTrue = monkeys[this.ifTrueIndex]
	this.ifFalse = monkeys[this.ifFalseIndex]
}

func parseMonkeyPart1(scanner *bufio.Scanner) *monkey_part1 {
	result := &monkey_part1{}

	scanner.Scan() // skip line with monkey index

	// starting items
	scanner.Scan()
	startingItems := splitStartingItems.Split(scanner.Text(), -1)[1:]
	result.worryLevels = make([]int, 0, 2*len(startingItems))
	for _, item := range startingItems {
		i, _ := strconv.Atoi(item)
		result.worryLevels = append(result.worryLevels, i)
	}

	// opperation
	scanner.Scan()
	opperation := strings.Split(strings.Split(scanner.Text(), "old ")[1], " ")
	number, err := strconv.Atoi(opperation[1])
	if err == nil {
		if opperation[0][0] == '+' {
			result.opperation = func (a int) int {
				return a + number
			}
		} else {
			result.opperation = func (a int) int {
				return a * number
			}
		}
	} else {
		result.opperation = func (a int) int {
			return a * a
		}
	}

	// test
	scanner.Scan()
	result.divisibilityTest, _ =
			strconv.Atoi(strings.Split(scanner.Text(), "by ")[1])

	// if true
	scanner.Scan()
	result.ifTrueIndex, _ =
			strconv.Atoi(strings.Split(scanner.Text(), "key ")[1])

	// if false
	scanner.Scan()
	result.ifFalseIndex, _ =
			strconv.Atoi(strings.Split(scanner.Text(), "key ")[1])

	return result
}

func lineCounter(r io.Reader) int {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count

		case err != nil:
			os.Exit(1)
			return count
		}
	}
}

func (d *Day) Day11Part1(filePath string) {
	file, _ := os.Open(filePath)
	monkeys := make([]*monkey_part1, (lineCounter(file) + 1) / 7)
	file.Close()
	file, _ = os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for i, _ := range monkeys {
		monkeys[i] = parseMonkeyPart1(scanner)

		if !scanner.Scan() {
			break
		}
	}

	for _, monkey := range monkeys {
		monkey.readIndexes(monkeys)
	}

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			monkey.round()
		}
	}

	maxItemsInspected1 := 0
	maxItemsInspected2 := 0
	for _, monkey := range monkeys {
		if monkey.inspectedItems > maxItemsInspected2 {
			if monkey.inspectedItems > maxItemsInspected1 {
				maxItemsInspected2 = maxItemsInspected1
				maxItemsInspected1 = monkey.inspectedItems
			} else {
				maxItemsInspected2 = monkey.inspectedItems
			}
		}
	}

	fmt.Println("product of 2 highest num of inspected items:",
			maxItemsInspected2 * maxItemsInspected1)
}
