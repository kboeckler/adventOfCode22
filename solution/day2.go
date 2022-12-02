package solution

import (
	"fmt"
	"strings"
)

func init() {
	RegisterSolution(2, day2{})
}

type day2 struct {
}

func (d day2) SolvePart1(input []string) interface{} {
	myShapes := make([]shape, 0, len(input))
	oppShapes := make([]shape, 0, len(input))
	for _, row := range input {
		symbols := strings.Split(row, " ")
		oppShapes = append(oppShapes, createShape(symbols[0]))
		myShapes = append(myShapes, createShape(symbols[1]))
	}
	score := 0
	for i := 0; i < len(myShapes); i++ {
		score += myShapes[i].score()
		score += myShapes[i].battleResultAgainst(oppShapes[i])
	}
	return score
}

func (d day2) SolvePart2(input []string) interface{} {
	myShapes := make([]shape, 0, len(input))
	oppShapes := make([]shape, 0, len(input))
	resultIndicators := make([]int, 0, len(input))
	for _, row := range input {
		symbols := strings.Split(row, " ")
		oppShapes = append(oppShapes, createShape(symbols[0]))
		resultIndicators = append(resultIndicators, createIndicator(symbols[1]))
	}
	for i := 0; i < len(oppShapes); i++ {
		myNeededShape := oppShapes[i].getShapeSoTheirResultIs(resultIndicators[i])
		myShapes = append(myShapes, myNeededShape)
	}
	score := 0
	for i := 0; i < len(myShapes); i++ {
		score += myShapes[i].score()
		score += resultIndicators[i]
	}
	return score
}

const (
	SHAPE_ROCK    = 1
	SHAPE_PAPER   = 2
	SHAPE_SCISSOR = 3
	RESULT_LOSE   = 0
	RESULT_DRAW   = 3
	RESULT_WIN    = 6
)

func createIndicator(symbol string) int {
	if strings.EqualFold("X", symbol) {
		return RESULT_LOSE
	}
	if strings.EqualFold("Y", symbol) {
		return RESULT_DRAW
	}
	if strings.EqualFold("Z", symbol) {
		return RESULT_WIN
	}
	panic("Illegal Indicator: " + symbol)
}

type shape struct {
	shapeType int
}

func createShape(symbol string) shape {
	shapeType := -1
	switch symbol {
	case "A":
		fallthrough
	case "X":
		shapeType = SHAPE_ROCK
	case "B":
		fallthrough
	case "Y":
		shapeType = SHAPE_PAPER
	case "C":
		fallthrough
	case "Z":
		shapeType = SHAPE_SCISSOR
	default:
		panic("Illegal Shape: " + symbol)
	}
	return shape{shapeType}
}

func (s shape) score() int {
	return s.shapeType
}

func (s shape) battleResultAgainst(other shape) int {
	switch s.shapeType {
	case SHAPE_ROCK:
		if other.shapeType == SHAPE_SCISSOR {
			return RESULT_WIN
		}
		if other.shapeType == SHAPE_ROCK {
			return RESULT_DRAW
		}
		return RESULT_LOSE
	case SHAPE_PAPER:
		if other.shapeType == SHAPE_ROCK {
			return RESULT_WIN
		}
		if other.shapeType == SHAPE_PAPER {
			return RESULT_DRAW
		}
		return RESULT_LOSE
	case SHAPE_SCISSOR:
		if other.shapeType == SHAPE_PAPER {
			return RESULT_WIN
		}
		if other.shapeType == SHAPE_SCISSOR {
			return RESULT_DRAW
		}
		return RESULT_LOSE
	default:
		panic(fmt.Sprintf("Unsupported shape type %d", s.shapeType))
	}
}

func (s shape) getShapeSoTheirResultIs(result int) shape {
	theirShapeType := -1
	switch result {
	case RESULT_WIN:
		if s.shapeType == SHAPE_ROCK {
			theirShapeType = SHAPE_PAPER
		} else if s.shapeType == SHAPE_PAPER {
			theirShapeType = SHAPE_SCISSOR
		} else {
			theirShapeType = SHAPE_ROCK
		}
	case RESULT_LOSE:
		if s.shapeType == SHAPE_PAPER {
			theirShapeType = SHAPE_ROCK
		} else if s.shapeType == SHAPE_SCISSOR {
			theirShapeType = SHAPE_PAPER
		} else {
			theirShapeType = SHAPE_SCISSOR
		}
	case RESULT_DRAW:
		theirShapeType = s.shapeType
	default:
		panic(fmt.Sprintf("Invalid result type: %d", result))
	}
	return shape{theirShapeType}
}
