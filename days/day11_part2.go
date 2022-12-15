package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var divisibilityTestsSet = NewSet[int]()
var divisibilityTest []int

type monkey_part2 struct {
	worryLevels []map[int]int
	opperation func(int) int
	divisibilityTest int
	ifTrueIndex, ifFalseIndex int // index of the monkey to throw to
	ifTrue, ifFalse *monkey_part2
	inspectedItems int
}

func (this *monkey_part2) round() {
	for _, item := range this.worryLevels {
		for test, remainder := range item {
			item[test] = this.opperation(remainder) % test
		}

		if item[this.divisibilityTest] == 0 {
			this.ifTrue.worryLevels =
					append(this.ifTrue.worryLevels, item)
		} else {
			this.ifFalse.worryLevels =
					append(this.ifFalse.worryLevels, item)
		}
	}

	this.inspectedItems += len(this.worryLevels)
	this.worryLevels = make([]map[int]int, 0, cap(this.worryLevels))
}

func newItem(value int) map[int]int {
	result := make(map[int]int)
	result[0] = value
	return result
}

func parseMonkeyPart2(scanner *bufio.Scanner) *monkey_part2 {
	result := &monkey_part2{}

	scanner.Scan() // skip line with monkey index

	// starting items
	scanner.Scan()
	startingItems := splitStartingItems.Split(scanner.Text(), -1)[1:]
	result.worryLevels = make([]map[int]int, 0, 2*len(startingItems))
	for _, item := range startingItems {
		i, _ := strconv.Atoi(item)
		result.worryLevels = append(result.worryLevels, newItem(i))
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
	divisibilityTest, _ := strconv.Atoi(strings.Split(scanner.Text(), "by ")[1])
	divisibilityTestsSet.Add(divisibilityTest)
	result.divisibilityTest = divisibilityTest

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

func (this *monkey_part2) postReadInit(monkeys []*monkey_part2) {
	for _, worryLevel := range this.worryLevels {
		initalValue := worryLevel[0]
		delete(worryLevel, 0)

		for _, test := range divisibilityTest {
			worryLevel[test] = initalValue % test
		}
	}

	this.ifTrue = monkeys[this.ifTrueIndex]
	this.ifFalse = monkeys[this.ifFalseIndex]
}

func (d *Day) Day11Part2(filePath string) {
	file, _ := os.Open(filePath)
	monkeys := make([]*monkey_part2, (lineCounter(file) + 1) / 7)
	file.Close()
	file, _ = os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for i, _ := range monkeys {
		monkeys[i] = parseMonkeyPart2(scanner)

		if !scanner.Scan() {
			break
		}
	}

	divisibilityTest = divisibilityTestsSet.ToSlice()

	for _, monkey := range monkeys {
		monkey.postReadInit(monkeys)
	}

	for i := 0; i < 10000; i++ {
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
