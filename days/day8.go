package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type forest struct {
	topLeft, bottomRight *tree
}

func (this *forest) calculateScenicScores() {
	// calculate scores from top left to bottom right
	for rowStart := this.topLeft; rowStart != nil; rowStart = rowStart.get(3) {
		for current := rowStart; current != nil; current = current.get(2) {
			current.calcScenicScoreDirection(0)
			current.calcScenicScoreDirection(1)
		}
	}

	// now calculate from bottom right to top left
	for rowStart := this.bottomRight; rowStart != nil;
			rowStart = rowStart.get(0) {
		for current := rowStart; current != nil; current = current.get(1) {
			current.calcScenicScoreDirection(2)
			current.calcScenicScoreDirection(3)
		}
	}
}

func (this *forest) maxScenicScore() int {
	this.calculateScenicScores()
	maxScore := math.MinInt

	for rowStart := this.topLeft; rowStart != nil; rowStart = rowStart.get(3) {
		for current := rowStart; current != nil; current = current.get(2) {
			maxScore = max(maxScore, current.getScenicScore())
		}
	}

	return maxScore
}

type tree struct {
	surrounding [4]*tree
	maxTreeHights [4]int
	height int
	scenicScores [4]int
	lastVisible [4]*tree
}

func (this *tree) calcScenicScoreDirection(direction int) (int, *tree) {
	if this.scenicScores[direction] >= 0 {
		return this.scenicScores[direction], this.lastVisible[direction]
	}
	bordering := this.surrounding[direction]
	if bordering == nil {
		this.scenicScores[direction] = 0
		this.lastVisible[direction] = nil
		return 0, nil
	}
	this.scenicScores[direction] = 1

	var score int
	for true {
		if bordering == nil {
			this.lastVisible[direction] = nil
			break
		}

		if bordering.height >= this.height {
			this.lastVisible[direction] = bordering
			break
		}

		score, bordering = bordering.calcScenicScoreDirection(direction)
		this.scenicScores[direction] += score
		this.lastVisible[direction] = bordering
	}

	return this.scenicScores[direction], this.lastVisible[direction]
}

func (this *tree) isVisible() bool {
	return this.height > this.maxTreeHights[0] ||
		   this.height > this.maxTreeHights[1] ||
		   this.height > this.maxTreeHights[2] ||
		   this.height > this.maxTreeHights[3]
}

func (this *tree) get(direction int) *tree {
	return this.surrounding[direction]
}

func (this *tree) getMaxHeight(direction int) int {
	return this.maxTreeHights[direction]
}

func (this *tree) getScenicScore() int {
	return this.scenicScores[0] * this.scenicScores[1] *
		   this.scenicScores[2] * this.scenicScores[3]
}

func (d *Day) Day8Part1(filePath string) {
	input := parseForest(filePath)
	visibleTreeCount := 0

	for firstOfRow := input.topLeft; firstOfRow != nil;
			firstOfRow = firstOfRow.get(3) {
		for current := firstOfRow; current != nil; current = current.get(2) {
			if current.isVisible() {
				visibleTreeCount++
			}
		}
	}

	fmt.Println("amount of visible trees:", visibleTreeCount)
}

func (d *Day) Day8Part2(filePath string) {
	input := parseForest(filePath)
	fmt.Println("best possible scenic score:", input.maxScenicScore())
}

func parseForest(filePath string) *forest {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var topLeft *tree
	var bottomRight *tree
	var firstOfRow *tree
	for scanner.Scan() {
		firstOfRow = parseForestRow(scanner.Text(), firstOfRow)
		if topLeft == nil {
			topLeft = firstOfRow
		}
	}

	for bottomRight = firstOfRow; bottomRight.get(2) != nil;
			bottomRight = bottomRight.get(2) {}

	for rowRight := bottomRight; rowRight != nil; rowRight = rowRight.get(0) {
		for current := rowRight; current != nil; current = current.get(1) {
			right := current.get(2)
			below := current.get(3)
			rightMaxHeight := -1
			belowMaxHeight := -1

			if right != nil {
				rightMaxHeight = max(right.getMaxHeight(2), right.height)
			}

			if below != nil {
				belowMaxHeight = max(below.getMaxHeight(3), below.height)
			}

			current.maxTreeHights[2] = rightMaxHeight
			current.maxTreeHights[3] = belowMaxHeight
		}
	}

	return &forest {
		topLeft, bottomRight,
	}
}

// returns the start of the row
func parseForestRow(line string, firstAbove *tree) *tree {
	var left *tree
	var current *tree
	var first *tree
	above := firstAbove
	for _, treeHeightRune := range line {
		aboveMaxHeight := -1
		leftMaxHeight := -1

		if above != nil {
			aboveMaxHeight = max(above.getMaxHeight(0), above.height)
		}

		if left == nil {
			// it is the first line and the first ree
		} else {
			leftMaxHeight = max(left.getMaxHeight(1), left.height)
		}

		height := int(treeHeightRune) - '0'

		current = &tree {
			[4]*tree {above, left, nil, nil},
			[4]int{aboveMaxHeight, leftMaxHeight, -1, -1},
			height,
			[4]int{-1, -1, -1, -1},
			[4]*tree{},
		}

		if left != nil {
			left.surrounding[2] = current
		} else if above == nil {
			first = current
		}

		if above != nil {
			above.surrounding[3] = current
			above = above.get(2)
		}

		left = current
	}

	if firstAbove == nil {
		return first
	}

	return firstAbove.get(3)
}
