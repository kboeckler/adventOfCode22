package solution

import (
	"regexp"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution(5, day5{})
}

type day5 struct {
}

func (d day5) SolvePart1(input []string) interface{} {
	stacks, instructionsReversed := parseInput(input)
	for i := len(instructionsReversed) - 1; i >= 0; i-- {
		instr := instructionsReversed[i]
		for j := 1; j <= instr.amount; j++ {
			stacks[instr.to] = append(stacks[instr.to], stacks[instr.from][len(stacks[instr.from])-j])
		}
		stacks[instr.from] = stacks[instr.from][0 : len(stacks[instr.from])-instr.amount]
	}
	result := strings.Builder{}
	for _, stack := range stacks {
		result.WriteString(string(stack[len(stack)-1]))
	}
	return result.String()
}

func (d day5) SolvePart2(input []string) interface{} {
	stacks, instructionsReversed := parseInput(input)
	for i := len(instructionsReversed) - 1; i >= 0; i-- {
		instr := instructionsReversed[i]
		for j := instr.amount; j > 0; j-- {
			stacks[instr.to] = append(stacks[instr.to], stacks[instr.from][len(stacks[instr.from])-j])
		}
		stacks[instr.from] = stacks[instr.from][0 : len(stacks[instr.from])-instr.amount]
	}
	result := strings.Builder{}
	for _, stack := range stacks {
		result.WriteString(string(stack[len(stack)-1]))
	}
	return result.String()
}

type instruction struct {
	amount, from, to int
}

func parseInput(input []string) ([][]uint8, []instruction) {
	stacks := make([][]uint8, 0)
	instructionsReversed := make([]instruction, 0)
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
			if strings.EqualFold(row, "") {
				continue
			}
			instructionsReversed = append(instructionsReversed, parseInstruction(row))
		} else {
			for i, _ := range stacks {
				index := 1 + i*4
				if len(row) > index {
					letter := row[index]
					if letter >= 65 && letter <= 90 {
						stacks[i] = append(stacks[i], letter)
					}
				}
			}
		}
	}
	return stacks, instructionsReversed
}

func parseInstruction(verbose string) instruction {
	reg, _ := regexp.Compile("move (\\d+) from (\\d+) to (\\d+)")
	parts := reg.FindAllStringSubmatch(verbose, -1)
	amount, _ := strconv.Atoi(parts[0][1])
	from, _ := strconv.Atoi(parts[0][2])
	to, _ := strconv.Atoi(parts[0][3])
	return instruction{amount, from - 1, to - 1}
}
