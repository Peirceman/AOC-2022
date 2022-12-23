package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

const maxXY = 4000000

var numberRegex = regexp.MustCompile("\\d+")

func fullCollapseRanges(ranges *set[[2]int]) *set[[2]int] {
	newRanges := collapseRanges(ranges)
	for newRanges.Length() > 2 {
		newRanges = collapseRanges(newRanges)
	}

	return collapseRanges(newRanges)
}

func collapseRanges(impossibleRanges *set[[2]int]) *set[[2]int] {
	newRanges := NewSet[[2]int]()
	toRemove := NewSet[[2]int]()

	for impossibleRange := range impossibleRanges.Itterate() {
		newRanges.Add(impossibleRange)

		for otherRange := range impossibleRanges.Itterate() {
			if impossibleRange == otherRange {
				continue
			}

			if impossibleRange[0] > otherRange[1]+1 ||
				otherRange[0] > impossibleRange[1]+1 {
				continue
			}

			newRange := [2]int{
				min(impossibleRange[0], otherRange[0]),
				max(impossibleRange[1], otherRange[1]),
			}

			newRanges.Add(newRange)

			if newRange != impossibleRange {
				toRemove.Add(impossibleRange)
			}

			if newRange != otherRange {
				toRemove.Add(otherRange)
			}
		}
	}

	for aRange := range toRemove.Itterate() {
		newRanges.Remove(aRange)
	}

	return newRanges
}

func impossibleRanges(distances map[coordinate]int,
	closestBeacons map[coordinate]coordinate, row int,
	includeSensor bool) *set[[2]int] {
	result := NewSet[[2]int]()

	for sensor, distance := range distances {
		beacon := closestBeacons[sensor]
		distanceToSide := distance - abs(sensor.y-row)

		if distanceToSide < 0 {
			continue
		}

		if distanceToSide == 0 {
			if beacon.y != row || includeSensor {
				result.Add([2]int{sensor.x, sensor.x})
			}
			continue
		}

		if beacon.y != row || includeSensor {
			result.Add([2]int{max(0, sensor.x-distanceToSide),
				min(maxXY, sensor.x+distanceToSide)})
			continue
		}

		if beacon.x == sensor.x-distanceToSide {
			result.Add([2]int{max(0, sensor.x-distanceToSide+1),
				min(maxXY, sensor.x+distanceToSide)})
			continue
		}

		if beacon.x == sensor.x+distanceToSide {
			result.Add([2]int{max(0, sensor.x-distanceToSide),
				min(maxXY, sensor.x+distanceToSide-1)})
			continue
		}

		result = collapseRanges(result)
	}

	result = fullCollapseRanges(result)
	if !includeSensor || result.Length() > 1 {
		return result
	}

	return nil
}

func (d Day) Day15Part1(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	distancesToBeacon := make(map[coordinate]int)
	closestBeacons := make(map[coordinate]coordinate)

	for scanner.Scan() {
		numbers := numberRegex.FindAllString(scanner.Text(), 4)

		sensor := coordinate{}
		sensor.x, _ = strconv.Atoi(numbers[0])
		sensor.y, _ = strconv.Atoi(numbers[1])

		beacon := coordinate{}
		beacon.x, _ = strconv.Atoi(numbers[2])
		beacon.y, _ = strconv.Atoi(numbers[3])

		distancesToBeacon[sensor] = sensor.manhattanDistance(beacon)
		closestBeacons[sensor] = beacon
	}

	const yToCheck = 2000000
	impossibleXs := impossibleRanges(distancesToBeacon, closestBeacons,
		yToCheck, false)
	impossibleXCount := 0

	for impossibleRange := range impossibleXs.Itterate() {
		impossibleXCount += impossibleRange[1] - impossibleRange[0] + 1
	}

	fmt.Println("amount of possitions without possible beacon",
		impossibleXCount)
}

func (d Day) Day15Part2(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	distancesToBeacon := make(map[coordinate]int)
	closestBeacons := make(map[coordinate]coordinate)

	for scanner.Scan() {
		numbers := numberRegex.FindAllString(scanner.Text(), 4)

		sensor := coordinate{}
		sensor.x, _ = strconv.Atoi(numbers[0])
		sensor.y, _ = strconv.Atoi(numbers[1])

		beacon := coordinate{}
		beacon.x, _ = strconv.Atoi(numbers[2])
		beacon.y, _ = strconv.Atoi(numbers[3])

		distancesToBeacon[sensor] = sensor.manhattanDistance(beacon)
		closestBeacons[sensor] = beacon
	}

	var possible coordinate

	for y := 0; y < maxXY; y++ {
		impossibleXs := impossibleRanges(distancesToBeacon,
			closestBeacons, y, true)
		fmt.Println(impossibleXs)

		if impossibleXs == nil || impossibleXs.Length() < 2 {
			continue
		}

		lowestMax := math.MaxInt
		for notHere := range impossibleXs.Itterate() {
			lowestMax = min(lowestMax, notHere[1])
		}

		if lowestMax > maxXY {
			continue
		}

		possible.x = lowestMax + 1
		possible.y = y
		fmt.Printf("%+v\n", possible)
		fmt.Println("tuning frequency:", possible.x*4000000+possible.y)
		break
	}

	fmt.Printf("%+v\n", possible)
	fmt.Println("tuning frequency:", possible.x*4000000+possible.y)
}
