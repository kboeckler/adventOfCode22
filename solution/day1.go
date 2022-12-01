package solution

import (
	"sort"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution(1, day1{})
}

type day1 struct {
}

func (d day1) SolvePart1(input []string) interface{} {
	maxCalories := -1
	currentCalories := 0
	for _, row := range input {
		if strings.EqualFold(row, "") {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
			continue
		}
		itemCalories, err := strconv.Atoi(row)
		if err != nil {
			panic(err)
		}
		currentCalories += itemCalories
	}
	return maxCalories
}

func (d day1) SolvePart2(input []string) interface{} {
	caloriesByElf := make([]int, 0)
	currentCalories := 0
	for _, row := range input {
		if strings.EqualFold(row, "") {
			if currentCalories > 0 {
				caloriesByElf = append(caloriesByElf, currentCalories)
			}
			currentCalories = 0
			continue
		}
		itemCalories, err := strconv.Atoi(row)
		if err != nil {
			panic(err)
		}
		currentCalories += itemCalories
	}
	sort.Ints(caloriesByElf)
	amountElfs := len(caloriesByElf)
	return caloriesByElf[amountElfs-1] + caloriesByElf[amountElfs-2] + caloriesByElf[amountElfs-3]
}
