package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type axisAlignedLine struct {
	start      coordinate
	horizontal bool
	end        int
}

// 498,4  498,6
// 498,4
// false
// y
func newAxisAlignedLine(start, end coordinate) *axisAlignedLine {
	if start.y == end.y {
		return &axisAlignedLine{
			start:      coordinate{min(start.x, end.x), start.y},
			horizontal: true,
			end:        max(start.x, end.x),
		}
	}

	return &axisAlignedLine{
		start:      coordinate{start.x, min(start.y, end.y)},
		horizontal: false,
		end:        max(start.y, end.y),
	}
}

func (line *axisAlignedLine) cointainsPoint(point coordinate) bool {
	if line.horizontal {
		return point.y == line.start.y &&
			line.start.x <= point.x && point.x <= line.end
	}

	return point.x == line.start.x &&
		line.start.y <= point.y && point.y <= line.end
}

func (line *axisAlignedLine) String() string {
	if line.horizontal {
		return strconv.Itoa(line.start.x) + "," + strconv.Itoa(line.start.y) +
			" -> " + strconv.Itoa(line.end) + "," + strconv.Itoa(line.start.y)
	}

	return strconv.Itoa(line.start.x) + "," + strconv.Itoa(line.start.y) +
		" -> " + strconv.Itoa(line.start.x) + "," + strconv.Itoa(line.end)
}

func collides(lines []*axisAlignedLine, otherSands []coordinate,
	sand coordinate) bool {
	for _, line := range lines {
		if line.cointainsPoint(sand) {
			return true
		}
	}

	for _, other := range otherSands {
		if other == sand {
			return true
		}
	}

	return false
}

func (d Day) Day14Part1(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]*axisAlignedLine, 0, 1000)
	sands := make([]coordinate, 0, 1000)

	lowestX, highestX, highestY := math.MinInt, math.MaxInt, math.MinInt

	for scanner.Scan() {
		pointStrs := strings.Split(scanner.Text(), " -> ")
		points := make([]coordinate, len(pointStrs))

		for i, point := range pointStrs {
			xYValues := strings.Split(point, ",")
			points[i] = coordinate{}
			points[i].x, _ = strconv.Atoi(xYValues[0])
			points[i].y, _ = strconv.Atoi(xYValues[1])

			lowestX = min(lowestX, points[i].x)
			highestX = max(highestX, points[i].x)
			highestY = max(highestY, points[i].y)
		}

		for i := range points[:len(points)-1] {
			lines = append(lines, newAxisAlignedLine(points[i], points[i+1]))
		}
	}

	sandOrigin := coordinate{500, 0}
	sand := sandOrigin

	for sand.y < highestY && lowestX < sand.x && sand.x < highestX {
		newPos := coordinate{sand.x, sand.y + 1}

		if !collides(lines, sands, newPos) {
			sand = newPos
			continue
		}

		newPos.x--
		if !collides(lines, sands, newPos) {
			sand = newPos
			continue
		}

		newPos.x += 2
		if !collides(lines, sands, newPos) {
			sand = newPos
			continue
		}

		sands = append(sands, sand)
		sand = sandOrigin
	}

	fmt.Println("amount of pieces that came to a rest:", len(sands))
}

func fillWithLine(occupiedSpaces *set[coordinate], start, end coordinate) {
	if start.y == end.y {
		max := max(start.x, end.x)
		for x := min(start.x, end.x); x <= max; x++ {
			occupiedSpaces.Add(coordinate{x, start.y})
		}
	}

	max := max(start.y, end.y)
	for y := min(start.y, end.y); y <= max; y++ {
		occupiedSpaces.Add(coordinate{start.x, y})
	}
}

func (d Day) Day14Part2(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	occupiedSpaces := NewSet[coordinate]()

	highestY := math.MinInt

	for scanner.Scan() {
		pointStrs := strings.Split(scanner.Text(), " -> ")
		points := make([]coordinate, len(pointStrs))

		for i, point := range pointStrs {
			xYValues := strings.Split(point, ",")
			points[i] = coordinate{}
			points[i].x, _ = strconv.Atoi(xYValues[0])
			points[i].y, _ = strconv.Atoi(xYValues[1])

			highestY = max(highestY, points[i].y)
		}

		for i := range points[:len(points)-1] {
			fillWithLine(occupiedSpaces, points[i], points[i+1])
		}
	}

	sandOrigin := coordinate{500, 0}
	sand := sandOrigin
	droppedPiecesCount := 0

	for {
		if sand.y > highestY {
			occupiedSpaces.Add(sand)
			droppedPiecesCount++
			sand = sandOrigin
			continue
		}

		newPos := coordinate{sand.x, sand.y + 1}
		if !occupiedSpaces.Contains(newPos) {
			sand = newPos
			continue
		}

		newPos.x--
		if !occupiedSpaces.Contains(newPos) {
			sand = newPos
			continue
		}

		newPos.x += 2
		if !occupiedSpaces.Contains(newPos) {
			sand = newPos
			continue
		}

		occupiedSpaces.Add(sand)
		droppedPiecesCount++
		if sand == sandOrigin {
			break
		}
		sand = sandOrigin
	}

	fmt.Println("amount of pieces that came to a rest:", droppedPiecesCount)
}
