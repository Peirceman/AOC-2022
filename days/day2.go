package days

import (
	"fmt"
	"os"
	"bufio"
)

func (d *Day) Day2Part1(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		opponentPlay := int(scanner.Text()[0] - 'A')
		myPlay := int(scanner.Text()[2] - 'X')
		score += myPlay + 1

		// draw
		if myPlay == opponentPlay {
			score += 3
			continue
		}

		// win
		if myPlay - 1 == opponentPlay || myPlay + 2 == opponentPlay {
			score += 6
		}

		// loss
	}

	fmt.Println("total score:", score)
}

func (d *Day) Day2Part2(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		opponentPlay := int(scanner.Text()[0] - 'A' + 1)
		result := int(scanner.Text()[2] - 'X')
		score += result * 3

		// draw
		if (result == 1) {
			score += opponentPlay
			continue
		}

		// win
		if (result == 2) {
			score += (opponentPlay % 3) + 1
			continue
		}

		// loss

		if (opponentPlay > 1) {
			score += opponentPlay - 1
			continue
		}

		score += 3
	}

	fmt.Println("total score:", score)
}
