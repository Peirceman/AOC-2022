package days

import (
	"bufio"
	"fmt"
	"os"
)

func (d *Day) Day6Part1(filePath string) {
	file,_ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	counts := make(map[byte]int)
	uniqueCharsCount := 0

	for i := 0; i < 4; i++ {
		if counts[line[i]] == 0 {
			uniqueCharsCount++
		}
		counts[line[i]]++
	}


	for i := 4; i < len(line); i++ {
		counts[line[i-4]]--

		if counts[line[i-4]] == 0 {
			uniqueCharsCount--
		}

		if counts[line[i]] == 0 {
			uniqueCharsCount++
		}

		counts[line[i]]++

		if uniqueCharsCount >= 4 {
			fmt.Println("start of packet marker:", i+1)
			break
		}
	}
}

func (d *Day) Day6Part2(filePath string) {
	file,_ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	counts := make(map[byte]int)
	uniqueCharsCount := 0

	for i := 0; i < 14; i++ {
		if counts[line[i]] == 0 {
			uniqueCharsCount++
		}
		counts[line[i]]++
	}


	for i := 14; i < len(line); i++ {
		counts[line[i-14]]--

		if counts[line[i-14]] == 0 {
			uniqueCharsCount--
		}

		if counts[line[i]] == 0 {
			uniqueCharsCount++
		}

		counts[line[i]]++

		if uniqueCharsCount >= 14 {
			fmt.Println("start of message marker:", i+1)
			break
		}
	}
}
