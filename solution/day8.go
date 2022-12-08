package solution

import "fmt"

func init() {
	RegisterSolution(8, day8{})
}

type day8 struct {
}

func (d day8) SolvePart1(input []string) interface{} {
	width := len(input[0])
	height := len(input)
	visibleTree := make(map[int]bool, 0)
	for y := 0; y < height; y++ {
		highest := input[y][0]
		visibleTree[coordToIndex(0, y, width)] = true
		for x := 1; x < width; x++ {
			current := input[y][x]
			if current > highest {
				visibleTree[coordToIndex(x, y, width)] = true
				highest = current
			}
		}
		highest = input[y][width-1]
		visibleTree[coordToIndex(width-1, y, width)] = true
		for x := width - 2; x >= 0; x-- {
			current := input[y][x]
			if current > highest {
				visibleTree[coordToIndex(x, y, width)] = true
				highest = current
			}
		}
	}
	for x := 0; x < width; x++ {
		highest := input[0][x]
		visibleTree[coordToIndex(x, 0, width)] = true
		for y := 1; y < height; y++ {
			current := input[y][x]
			if current > highest {
				visibleTree[coordToIndex(x, y, width)] = true
				highest = current
			}
		}
		highest = input[height-1][x]
		visibleTree[coordToIndex(x, height-1, width)] = true
		for y := height - 2; y >= 0; y-- {
			current := input[y][x]
			if current > highest {
				visibleTree[coordToIndex(x, y, width)] = true
				highest = current
			}
		}
	}
	return len(visibleTree)
}

func (d day8) SolvePart2(input []string) interface{} {
	width := len(input[0])
	height := len(input)
	vision := make(map[int]int, 0)
	for y := 0; y < height; y++ {
		highestX := 0
		highestY := y
		highest := input[highestY][highestX]
		for x := 1; x < width; x++ {
			current := input[y][x]
			if current > highest {
				highest = current
				highestX = x
				highestY = y
			} else {
				vision[coordToIndex(x, y, width)] = vision[coordToIndex(x, y, width)] + x - highestX
			}
		}
		highestX = width - 1
		highestY = y
		highest = input[highestY][highestX]
		for x := width - 2; x >= 0; x-- {
			current := input[y][x]
			if current > highest {
				highest = current
				highestX = x
				highestY = y
			} else {
				vision[coordToIndex(x, y, width)] = vision[coordToIndex(x, y, width)] + highestX - x
			}
		}
	}
	for x := 0; x < width; x++ {
		highestX := x
		highestY := 0
		highest := input[highestY][highestX]
		for y := 1; y < height; y++ {
			current := input[y][x]
			if current > highest {
				highest = current
				highestX = x
				highestY = y
			} else {
				vision[coordToIndex(x, y, width)] = vision[coordToIndex(x, y, width)] + y - highestY
			}
		}
		highestX = x
		highestY = height - 1
		highest = input[highestY][highestX]
		for y := height - 2; y >= 0; y-- {
			current := input[y][x]
			if current > highest {
				highest = current
				highestX = x
				highestY = y
			} else {
				vision[coordToIndex(x, y, width)] = vision[coordToIndex(x, y, width)] + highestY - y
			}
		}
	}
	for index, val := range vision {
		x, y := indexToCoord(index, width)
		fmt.Printf("%d %d has vision of %d\n", x, y, val)
	}
	return ""
}

func coordToIndex(x, y, width int) int {
	return y*width + x
}

func indexToCoord(index, width int) (int, int) {
	return index % width, index / width
}
