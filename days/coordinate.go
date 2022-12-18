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

// this means that if the distance < 2, the 2 coordinates are touching
func (this coordinate) chebyshevDistance(other coordinate) int {
	return int(max(abs(this.x-other.x), abs(this.y-other.y)))
}

// moves 1 step to the other coordinate if the two coordinates ar to far away
func (this coordinate) moveTo(other coordinate) coordinate {
	if this.chebyshevDistance(other) < 2 {
		return this
	}

	result := coordinate{}
	result.x = this.x + clamp(other.x-this.x, -1, 1)
	result.y = this.y + clamp(other.y-this.y, -1, 1)

	return result
}
