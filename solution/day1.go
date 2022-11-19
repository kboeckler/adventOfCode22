package solution

import (
	"strconv"
)

func init() {
	RegisterSolution(1, day1{})
}

type day1 struct {
}

func (d day1) SolvePart1(input []string) interface{} {
	amountOfIncrements := 0
	for i := 1; i < len(input); i++ {
		previous, _ := strconv.Atoi(input[i-1])
		current, _ := strconv.Atoi(input[i])
		if previous < current {
			amountOfIncrements++
		}
	}
	return amountOfIncrements
}

func (d day1) SolvePart2(input []string) interface{} {
	amountOfIncrements := 0
	for i := 3; i < len(input); i++ {
		firstOfPreviousWindow, _ := strconv.Atoi(input[i-3])
		secondOfPreviousWindow, _ := strconv.Atoi(input[i-2])
		thirdOfPreviousWindow, _ := strconv.Atoi(input[i-1])
		firstOfCurrentWindow := secondOfPreviousWindow
		secondOfCurrentWindow := thirdOfPreviousWindow
		thirdOfCurrentWindow, _ := strconv.Atoi(input[i-0])
		current := firstOfCurrentWindow + secondOfCurrentWindow + thirdOfCurrentWindow
		previous := firstOfPreviousWindow + secondOfPreviousWindow + thirdOfPreviousWindow
		if previous < current {
			amountOfIncrements++
		}
	}
	return amountOfIncrements
}
