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
	return d.solveRopeWith(input, 2)
}

func (d day9) SolvePart2(input []string) interface{} {
	return d.solveRopeWith(input, 10)
}

func (d day9) solveRopeWith(input []string, amountKnots int) interface{} {
	moves := d.parseMoves(input)
	dim := d.findDimension(moves)
	tailPositions := make(map[int]bool)
	knots := make([]int, 0, amountKnots)
	for i := 0; i < amountKnots; i++ {
		knotX := -dim.negativeWidthOffset
		knotY := -dim.negativeHeightOffset
		knots = append(knots, d.coordToIndex(knotX, knotY, dim.width))
	}
	tailPositions[knots[len(knots)-1]] = true
	for _, move := range moves {
		headX, headY := d.indexToCoord(knots[0], dim.width)
		headX, headY = move.nextHead(headX, headY)
		knots[0] = d.coordToIndex(headX, headY, dim.width)
		for i := 0; i < len(knots)-1; i++ {
			leaderX, leaderY := d.indexToCoord(knots[i], dim.width)
			followerX, followerY := d.indexToCoord(knots[i+1], dim.width)
			followerX, followerY = d.updateKnotPosition(leaderX, leaderY, followerX, followerY)
			knots[i+1] = d.coordToIndex(followerX, followerY, dim.width)
		}
		tailPositions[knots[len(knots)-1]] = true
	}
	return len(tailPositions)
}

type dimension struct {
	width, height, negativeWidthOffset, negativeHeightOffset int
}

type move struct {
	direction string
}

func (d day9) findDimension(moves []*move) dimension {
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
	dim := dimension{width, height, negativeWidthOffset, negativeHeightOffset}
	return dim
}

func (d day9) parseMoves(input []string) []*move {
	moves := make([]*move, 0)
	for _, row := range input {
		split := strings.Split(row, " ")
		steps, _ := strconv.Atoi(split[1])
		for i := 0; i < steps; i++ {
			moves = append(moves, &move{split[0]})
		}
	}
	return moves
}

func (d day9) coordToIndex(x, y, width int) int {
	return y*width + x
}

func (d day9) indexToCoord(index, width int) (int, int) {
	return index % width, index / width
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

func (d day9) updateKnotPosition(leaderX, leaderY, followerX, followerY int) (int, int) {
	headTailDist2 := d.euclideanDist2(leaderX, leaderY, followerX, followerY)
	if headTailDist2 >= 4.0 {
		newMinDist2 := 4.0
		bestX := 0
		bestY := 0
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				dist2 := d.euclideanDist2(leaderX, leaderY, followerX+dx, followerY+dy)
				if dist2 < newMinDist2 {
					newMinDist2 = dist2
					bestX = followerX + dx
					bestY = followerY + dy
				}
			}
		}
		return bestX, bestY
	}
	return followerX, followerY
}

func (d day9) euclideanDist2(ax, ay, bx, by int) float64 {
	return math.Pow(float64(ax-bx), 2) + math.Pow(float64(ay-by), 2)
}
