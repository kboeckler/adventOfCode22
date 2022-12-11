package solution

import (
	"math"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution(10, day10{})
}

type day10 struct {
}

func (d day10) SolvePart1(input []string) interface{} {
	xValues := d.calculateXValuesPerCycle(input)
	sum := 0
	for i := 20; i <= 220; i = i + 40 {
		sum += xValues[i-1] * i
	}
	return sum
}

func (d day10) SolvePart2(input []string) interface{} {
	xValues := d.calculateXValuesPerCycle(input)
	crt := strings.Builder{}
	for i := 0; i < 240; i++ {
		column := i % 40
		if column == 0 {
			crt.WriteString("\n")
		}
		distanceSprite := int(math.Abs(float64(column - xValues[i])))
		if distanceSprite <= 1 {
			crt.WriteString("#")
		} else {
			crt.WriteString(".")
		}
	}
	return crt.String()
}

func (d day10) calculateXValuesPerCycle(input []string) []int {
	xValues := make([]int, 0)
	xValues = append(xValues, 1)
	for _, row := range input {
		if strings.EqualFold(row, "noop") {
			xValues = append(xValues, xValues[len(xValues)-1])
		} else {
			split := strings.Split(row, " ")
			summand, _ := strconv.Atoi(split[1])
			xValues = append(xValues, xValues[len(xValues)-1])
			xValues = append(xValues, xValues[len(xValues)-1]+summand)
		}
	}
	return xValues
}
