package solution

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution(5, day5{})
}

type day5 struct {
}

func (d day5) SolvePart1(input []string) interface{} {
	instructionsReversed := make([]string, 0)
	stacks := make([][]uint8, 0)
	parsingInstructions := true
	for i := len(input) - 1; i >= 0; i-- {
		row := input[i]
		if !strings.EqualFold(strings.TrimLeft(row, " 1"), row) {
			parsingInstructions = false
			lastStackId, _ := strconv.Atoi(string(row[len(row)-1]))
			for i := 0; i < lastStackId; i++ {
				stacks = append(stacks, make([]uint8, 0))
			}
			continue
		}
		if parsingInstructions {
			instructionsReversed = append(instructionsReversed, row)
		} else {
			stacksRaw := strings.Split(strings.Replace(row, "[", "", 1), "[")
			for i, raw := range stacksRaw {
				letter := raw[0]
				if letter >= 65 && letter <= 90 {
					stacks[i] = append(stacks[i], letter)
				}
			}
		}
	}
	for _, ins := range stacks {
		for _, crate := range ins {
			fmt.Print(string(crate))
		}
		fmt.Println("-")
	}
	return ""
}

func (d day5) SolvePart2(input []string) interface{} {
	return ""
}
