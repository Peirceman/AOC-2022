package days

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
)

type coordinate struct {
	x, y int
}

func (this coordinate) add(other coordinate) coordinate {
	return coordinate {
		this.x + other.x,
		this.y + other.y,
	}
}

// this means that if the distance < 2, the 2 coordinates are touching
func (this coordinate) chebyshevDistance(other coordinate) int {
	return int(max(abs(this.x - other.x), abs(this.y - other.y)))
}

// moves 1 step to the other coordinate if the two coordinates ar to far away
func (this coordinate) moveTo(other coordinate) coordinate {
	if this.chebyshevDistance(other) < 2 {
		return this
	}

	result := coordinate{}
	result.x = this.x + clamp(other.x - this.x, -1, 1)
	result.y = this.y + clamp(other.y - this.y, -1, 1)

	return result
}

func directionToInt(dir uint8) coordinate {
	x := 0
	y := 0

	switch dir {
		case 'U':
			y = 1
		case 'L':
			x = -1
		case 'R':
			x = 1
		case 'D':
			y = -1
	}

	return coordinate { x, y }
}

func (d *Day) Day9Part1(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	head := coordinate { 0, 0, }
	tail := coordinate { 0, 0, }
	visited := NewSet[coordinate](tail)

	for scanner.Scan() {
		velocity := directionToInt(scanner.Text()[0])
		count, _ := strconv.Atoi(scanner.Text()[2:])

		for ; count > 0; count-- {
			head = head.add(velocity)
			tail = tail.moveTo(head)
			// fmt.Println(head, tail)
			visited.Add(tail)
		}
	}

	fmt.Println(visited.Length())
}

func (d *Day) Day9Part2(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rope := [10]coordinate{}
	for i, _ := range rope {
		rope[i] = coordinate { 0, 0 }
	}

	visited := NewSet[coordinate](rope[9])

	for scanner.Scan() {
		velocity := directionToInt(scanner.Text()[0])
		count, _ := strconv.Atoi(scanner.Text()[2:])

		for ; count > 0; count-- {
			rope[0] = rope[0].add(velocity)

			for i := 1; i < len(rope); i++ {
				rope[i] = rope[i].moveTo(rope[i-1])
			}

			visited.Add(rope[9])
		}
	}

	fmt.Println(visited.Length())
}

