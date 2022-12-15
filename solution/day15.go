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
	return d.solve1(input, 2_000_000)
}

func (d day15) SolvePart2(input []string) interface{} {
	return d.solve2(input, 4_000_000, 4_000_000)
}

func (d day15) solve1(input []string, row int) interface{} {
	sensorPositions, beaconPositions, sensors := d.parseInput(input)
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
		pos := cavePos{x, row}
		if sensorPositions[pos] || beaconPositions[pos] {
			continue
		}
		isImpossible := false
		for _, sens := range sensors {
			if sens.pos.distTo(pos.x, pos.y) <= sens.length {
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

func (d day15) solve2(input []string, maxX, maxY int) interface{} {
	sensorPositions, beaconPositions, sensors := d.parseInput(input)
	candidates := make([]cavePos, 0)
	for _, sens := range sensors {
		outerX := sens.pos.x - sens.length - 1
		outerY := sens.pos.y
		if outerX >= 0 && outerX <= maxX && outerY >= 0 && outerY <= maxY {
			candidates = append(candidates, cavePos{outerX, outerY})
		}
		for i := 0; i < sens.length+1; i++ {
			outerX += 1
			outerY -= 1
			if outerX >= 0 && outerX <= maxX && outerY >= 0 && outerY <= maxY {
				candidates = append(candidates, cavePos{outerX, outerY})
			}
		}
		for i := 0; i < sens.length+1; i++ {
			outerX += 1
			outerY += 1
			if outerX >= 0 && outerX <= maxX && outerY >= 0 && outerY <= maxY {
				candidates = append(candidates, cavePos{outerX, outerY})
			}
		}
		for i := 0; i < sens.length+1; i++ {
			outerX -= 1
			outerY += 1
			if outerX >= 0 && outerX <= maxX && outerY >= 0 && outerY <= maxY {
				candidates = append(candidates, cavePos{outerX, outerY})
			}
		}
		for i := 0; i < sens.length; i++ {
			outerX -= 1
			outerY -= 1
			if outerX >= 0 && outerX <= maxX && outerY >= 0 && outerY <= maxY {
				candidates = append(candidates, cavePos{outerX, outerY})
			}
		}
	}
	var distressBeaconPos *cavePos
	for _, candidate := range candidates {
		if sensorPositions[candidate] || beaconPositions[candidate] {
			continue
		}
		isImpossible := false
		for _, sens := range sensors {
			if sens.pos.distTo(candidate.x, candidate.y) <= sens.length {
				isImpossible = true
				break
			}
		}
		if !isImpossible {
			distressBeaconPos = &candidate
			break
		}
	}
	return distressBeaconPos.x*4_000_000 + distressBeaconPos.y
}

func (d day15) parseInput(input []string) (map[cavePos]bool, map[cavePos]bool, []sensor) {
	sensorPositions := make(map[cavePos]bool)
	beaconPositions := make(map[cavePos]bool)
	sensors := make([]sensor, 0)
	for _, row := range input {
		numbers := d.parseNumbersFrom(row)
		sensorPos := cavePos{numbers[0], numbers[1]}
		beaconPos := cavePos{numbers[2], numbers[3]}
		length := sensorPos.distTo(beaconPos.x, beaconPos.y)
		sensorPositions[sensorPos] = true
		sensors = append(sensors, sensor{sensorPos, beaconPos, length})
		beaconPositions[beaconPos] = true
	}
	return sensorPositions, beaconPositions, sensors
}

func (d day15) parseNumbersFrom(row string) []int {
	re, _ := regexp.Compile("-?\\d+")
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
