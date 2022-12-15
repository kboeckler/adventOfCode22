package solution

import (
	"regexp"
	"strconv"
)

func init() {
	RegisterSolution(15, day15{})
}

type day15 struct {
}

func (d day15) SolvePart1(input []string) interface{} {
	sensors := make(map[cavePos]bool)
	beacons := make(map[cavePos]bool)
	positions := make([]cavePos, 0)
	for _, row := range input {
		numbers := d.parseNumbersFrom(row)
		sensor := cavePos{numbers[0], numbers[1]}
		beacon := cavePos{numbers[2], numbers[3]}
		sensors[sensor] = true
		beacons[beacon] = true
		positions = append(positions, sensor)
		positions = append(positions, beacon)
	}
	impossiblePositions := make(map[cavePos]bool)
	for sensor, _ := range sensors {
		for _, candidate := range d.findSurroundingCavePosContainingOneOf(sensor, beacons) {
			if !sensors[candidate] && !beacons[candidate] {
				impossiblePositions[candidate] = true
			}
		}
	}
	amount := 0
	for impossible, _ := range impossiblePositions {
		if impossible.y == 2_000_000 {
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

func (d day15) findSurroundingCavePosContainingOneOf(from cavePos, beacons map[cavePos]bool) []cavePos {
	// TODO
	return make([]cavePos, 0)
}

type cavePos struct {
	x, y int
}
