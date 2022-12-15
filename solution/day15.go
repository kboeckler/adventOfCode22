package solution

import (
	"math"
	"regexp"
	"strconv"
)

func init() {
	RegisterSolution(15, day15{})
}

type day15 struct {
}

func (d day15) SolvePart1(input []string) interface{} {
	return d.solve(input, 2_000_000)
}

func (d day15) solve(input []string, row int) interface{} {
	sensorPositions := make(map[cavePos]bool)
	beaconPositions := make(map[cavePos]bool)
	sensors := make([]sensor, 0)
	positions := make([]cavePos, 0)
	for _, row := range input {
		numbers := d.parseNumbersFrom(row)
		sensorPos := cavePos{numbers[0], numbers[1]}
		beaconPos := cavePos{numbers[2], numbers[3]}
		length := sensorPos.distTo(beaconPos.x, beaconPos.y)
		sensorPositions[sensorPos] = true
		sensors = append(sensors, sensor{sensorPos, beaconPos, length})
		beaconPositions[beaconPos] = true
		positions = append(positions, sensorPos)
		positions = append(positions, beaconPos)
	}
	minX, maxX := 9999999, -999999
	for _, sens := range sensors {
		if sens.pos.x < minX {
			minX = sens.pos.x
		}
		if sens.closestBeacon.x < minX {
			minX = sens.closestBeacon.x
		}
		if sens.pos.x > maxX {
			maxX = sens.pos.x
		}
		if sens.closestBeacon.x > maxX {
			maxX = sens.closestBeacon.x
		}
	}
	maxLength := 0
	for _, sens := range sensors {
		if sens.length > maxLength {
			maxLength = sens.length
		}
	}
	amount := 0
	for x := minX - maxLength; x < maxX+maxLength; x++ {
		isImpossible := false
		for _, sens := range sensors {
			pos := cavePos{x, row}
			if sens.pos.distTo(pos.x, pos.y) <= sens.length && !sensorPositions[pos] && !beaconPositions[pos] {
				isImpossible = true
				break
			}
		}
		if isImpossible {
			amount++
		}
	}
	return amount
}

func (d day15) SolvePart2(input []string) interface{} {
	return 0
}

func (d day15) parseNumbersFrom(row string) []int {
	re, _ := regexp.Compile("\\d+")
	matches := re.FindAllStringSubmatch(row, -1)
	numbers := make([]int, 0)
	for _, single := range matches {
		number, _ := strconv.Atoi(single[0])
		numbers = append(numbers, number)
	}
	return numbers
}

type cavePos struct {
	x, y int
}

func (p cavePos) distTo(x, y int) int {
	return int(math.Abs(float64(p.x-x)) + math.Abs(float64(p.y-y)))
}

type sensor struct {
	pos           cavePos
	closestBeacon cavePos
	length        int
}

func (p cavePos) neighbors() []cavePos {
	neighbors := make([]cavePos, 4)
	neighbors[0] = cavePos{p.x + 1, p.y}
	neighbors[1] = cavePos{p.x - 1, p.y}
	neighbors[2] = cavePos{p.x, p.y + 1}
	neighbors[3] = cavePos{p.x, p.y - 1}
	return neighbors
}
