package solution

import (
	"math"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution(9, day9{})
}

type day9 struct {
}

func (d day9) SolvePart1(input []string) interface{} {
	moves := make([]*move, 0)
	for _, row := range input {
		split := strings.Split(row, " ")
		steps, _ := strconv.Atoi(split[1])
		for i := 0; i < steps; i++ {
			moves = append(moves, &move{split[0]})
		}
	}
	currentWidth := 1
	negativeWidth := 1
	width := currentWidth
	currentHeight := 1
	negativeHeight := 1
	height := currentHeight
	for _, move := range moves {
		switch move.direction {
		case "R":
			currentWidth++
		case "L":
			currentWidth--
		case "U":
			currentHeight++
		case "D":
			currentHeight--
		}
		if currentWidth > width {
			width = currentWidth
		}
		if currentWidth < negativeWidth {
			negativeWidth = currentWidth
		}
		if currentHeight > height {
			height = currentHeight
		}
		if currentHeight < negativeHeight {
			negativeHeight = currentHeight
		}
	}
	negativeWidthOffset := negativeWidth - 1
	negativeHeightOffset := negativeHeight - 1
	width -= negativeWidthOffset
	height -= negativeHeightOffset
	tailPositions := make(map[int]bool)
	headX := -negativeWidthOffset
	headY := -negativeHeightOffset
	tailX := -negativeWidthOffset
	tailY := -negativeHeightOffset
	tailPositions[d.coordToIndex(tailX, tailY, width)] = true
	for _, move := range moves {
		headX, headY = move.nextHead(headX, headY)
		tailX, tailY = d.udateTail(headX, headY, tailX, tailY)
		tailPositions[d.coordToIndex(tailX, tailY, width)] = true
	}
	return len(tailPositions)
}

func (d day9) SolvePart2(input []string) interface{} {
	return ""
}

func (d day9) coordToIndex(x, y, width int) int {
	return y*width + x
}

type move struct {
	direction string
}

func (m *move) nextHead(hx, hy int) (int, int) {
	switch m.direction {
	case "U":
		return hx, hy + 1
	case "D":
		return hx, hy - 1
	case "R":
		return hx + 1, hy
	case "L":
		return hx - 1, hy
	}
	panic("Invalid direction in move: " + m.direction)
}

func (d day9) udateTail(hx, hy, tx, ty int) (int, int) {
	headTailDist2 := d.euclideanDist2(hx, hy, tx, ty)
	if headTailDist2 >= 4.0 {
		newMinDist2 := 4.0
		bestX := 0
		bestY := 0
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				dist2 := d.euclideanDist2(hx, hy, tx+dx, ty+dy)
				if dist2 < newMinDist2 {
					newMinDist2 = dist2
					bestX = tx + dx
					bestY = ty + dy
				}
			}
		}
		return bestX, bestY
	}
	return tx, ty
}

func (d day9) euclideanDist2(ax, ay, bx, by int) float64 {
	return math.Pow(float64(ax-bx), 2) + math.Pow(float64(ay-by), 2)
}
