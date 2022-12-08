package solution

import (
	"sort"
)

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
	visionToLeft := make(map[int]int, 0)
	visionToRight := make(map[int]int, 0)
	visionToTop := make(map[int]int, 0)
	visionToBottom := make(map[int]int, 0)
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
				visionToLeft[coordToIndex(x, y, width)] = visionToLeft[coordToIndex(x, y, width)] + x
			} else {
				sight := 1
				for sight = 1; sight <= width-x; sight++ {
					if input[y][x-sight] >= current {
						break
					}
				}
				visionToLeft[coordToIndex(x, y, width)] = visionToLeft[coordToIndex(x, y, width)] + sight
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
				visionToRight[coordToIndex(x, y, width)] = visionToRight[coordToIndex(x, y, width)] + width - 1 - x
			} else {
				sight := 1
				for sight = 1; sight <= x; sight++ {
					if input[y][x+sight] >= current {
						break
					}
				}
				visionToRight[coordToIndex(x, y, width)] = visionToRight[coordToIndex(x, y, width)] + sight
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
				visionToTop[coordToIndex(x, y, width)] = visionToTop[coordToIndex(x, y, width)] + y
			} else {
				sight := 1
				for sight = 1; sight <= height-y; sight++ {
					if input[y-sight][x] >= current {
						break
					}
				}
				visionToTop[coordToIndex(x, y, width)] = visionToTop[coordToIndex(x, y, width)] + sight
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
				visionToBottom[coordToIndex(x, y, width)] = visionToBottom[coordToIndex(x, y, width)] + height - 1 - y
			} else {
				sight := 1
				for sight = 1; sight <= y; sight++ {
					if input[y+sight][x] >= current {
						break
					}
				}
				visionToBottom[coordToIndex(x, y, width)] = visionToBottom[coordToIndex(x, y, width)] + sight
			}
		}
	}
	scenicScores := make([]int, 0, width*height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			index := coordToIndex(x, y, width)
			score := visionToLeft[index] * visionToRight[index] * visionToTop[index] * visionToBottom[index]
			scenicScores = append(scenicScores, score)
		}
	}
	sort.Ints(scenicScores)
	return scenicScores[len(scenicScores)-1]
}

func coordToIndex(x, y, width int) int {
	return y*width + x
}
