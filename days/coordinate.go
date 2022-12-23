package days

type coordinate struct {
	x, y int
}

func (this coordinate) add(other coordinate) coordinate {
	return coordinate{
		this.x + other.x,
		this.y + other.y,
	}
}

func (coord coordinate) manhattanDistance(other coordinate) int {
	return abs(coord.x-other.x) + abs(coord.y-other.y)
}

func (this coordinate) chebyshevDistance(other coordinate) int {
	return max(abs(this.x-other.x), abs(this.y-other.y))
}
