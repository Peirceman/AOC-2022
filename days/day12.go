package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type heightMapPositions []*heightMapPosition

func (slice heightMapPositions) Len() int {
	return 4
}

func (slice heightMapPositions) Less(i, j int) bool {
	if slice[i] == nil {
		return slice[j] != nil
	}

	if slice[j] == nil {
		return false
	}

	return slice[i].height < slice[j].height
}

func (slice heightMapPositions) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

type heightMapPosition struct {
	surrounding           [4]*heightMapPosition
	height, distanceToEnd int
	beenVisited           bool
}

type heightMap struct {
	topleft, start, end *heightMapPosition
}

func (heightMap *heightMap) calculateShortestPaths() {
	surroundingPositions := NewSet(heightMap.end.surrounding[:]...)
	nextSurrounding := NewSet[*heightMapPosition]()
	heightMap.end.distanceToEnd = 0
	heightMap.end.beenVisited = true

	for surroundingPositions.Length() > 0 {
		for _, current := range surroundingPositions.ToSlice() {
			if current == nil || current.beenVisited {
				continue
			}

			for _, surrounding := range current.surrounding {
				if surrounding == nil || !surrounding.beenVisited ||
					surrounding.height > current.height+1 {
					continue
				}

				if surrounding.distanceToEnd < current.distanceToEnd {
					current.distanceToEnd = surrounding.distanceToEnd + 1
					current.beenVisited = true
				}
			}

			if current.beenVisited {
				nextSurrounding.AddSlice(current.surrounding[:])
			}
		}

		surroundingPositions = nextSurrounding
		nextSurrounding = NewSet[*heightMapPosition]()
	}
}

func parseHeightMap(scanner *bufio.Scanner) *heightMap {
	result := &heightMap{}
	var firstOfRow *heightMapPosition

	for scanner.Scan() {
		firstOfRow = parseHeightMapRow(scanner.Text(), firstOfRow, result)
		if result.topleft == nil {
			result.topleft = firstOfRow
		}
	}

	return result
}

func parseHeightMapRow(line string, firstAbove *heightMapPosition,
	heightMap *heightMap) *heightMapPosition {
	var left, first *heightMapPosition
	above := firstAbove

	for _, heightChar := range line {
		var height int
		if heightChar == 'S' {
			height = 0
		} else if heightChar == 'E' {
			height = 25
		} else {
			height = int(heightChar - 'a')
		}

		current := &heightMapPosition{
			surrounding:   [4]*heightMapPosition{above, left, nil, nil},
			height:        height,
			distanceToEnd: math.MaxInt,
			beenVisited:   false,
		}

		if heightChar == 'S' {
			heightMap.start = current
		} else if heightChar == 'E' {
			heightMap.end = current
		}

		if left != nil {
			left.surrounding[2] = current
		}

		left = current
		if above != nil {
			above.surrounding[3] = current
			above = above.surrounding[2]
		}

		if first == nil {
			first = current
		}
	}

	return first
}

func (d Day) Day12Part1(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	heightMap := parseHeightMap(scanner)
	heightMap.calculateShortestPaths()

	fmt.Println("shortest distance to end:", heightMap.start.distanceToEnd)
}

func (d Day) Day12Part2(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	heightMap := parseHeightMap(scanner)
	heightMap.calculateShortestPaths()

	shortestDistance := math.MaxInt

	for rowstart := heightMap.topleft; rowstart != nil; rowstart =
		rowstart.surrounding[3] {
		for current := rowstart; current != nil; current =
			current.surrounding[2] {
			if current.height > 0 {
				continue
			}

			shortestDistance = min(shortestDistance, current.distanceToEnd)
		}
	}

	fmt.Println("shortest distance to end:", shortestDistance)
}
