package solution

import (
	"strconv"
	"strings"
)

func init() {
	RegisterSolution(4, day4{})
}

type day4 struct {
}

func (d day4) SolvePart1(input []string) interface{} {
	return d.solve(input, func(a *assignment) bool {
		return a.matchesFully()
	})
}

func (d day4) SolvePart2(input []string) interface{} {
	return d.solve(input, func(a *assignment) bool {
		return a.matchesPartially()
	})
}

func (d day4) solve(input []string, matchCondition func(a *assignment) bool) interface{} {
	assignments := make([]*assignment, 0, len(input))
	for _, row := range input {
		bothElves := strings.Split(row, ",")
		assignments = append(assignments, createAssignment(bothElves[0], bothElves[1]))
	}
	matchingAssignments := 0
	for _, pair := range assignments {
		if matchCondition(pair) {
			matchingAssignments++
		}
	}
	return matchingAssignments
}

func createAssignment(elf, otherElf string) *assignment {
	elfInterval := strings.Split(elf, "-")
	elfFrom, _ := strconv.Atoi(elfInterval[0])
	elfTo, _ := strconv.Atoi(elfInterval[1])
	otherElfInterval := strings.Split(otherElf, "-")
	otherElfFrom, _ := strconv.Atoi(otherElfInterval[0])
	otherElfTo, _ := strconv.Atoi(otherElfInterval[1])
	return &assignment{elfFrom, elfTo, otherElfFrom, otherElfTo}
}

type assignment struct {
	elfFrom, elfTo, otherElfFrom, otherElfTo int
}

func (a *assignment) matchesFully() bool {
	elfContainsOtherElf := a.elfFrom <= a.otherElfFrom && a.elfTo >= a.otherElfTo
	otherElfContainsElf := a.otherElfFrom <= a.elfFrom && a.otherElfTo >= a.elfTo
	return elfContainsOtherElf || otherElfContainsElf
}

func (a *assignment) matchesPartially() bool {
	elfContainsOtherElfFrom := a.otherElfFrom >= a.elfFrom && a.otherElfFrom <= a.elfTo
	elfContainsOtherElfTo := a.otherElfTo >= a.elfFrom && a.otherElfTo <= a.elfTo
	otherElfContainsElfFrom := a.elfFrom >= a.otherElfFrom && a.elfFrom <= a.otherElfTo
	otherElfContainsElfTo := a.elfTo >= a.otherElfFrom && a.elfTo <= a.otherElfTo
	elfOverlapsOtherElf := elfContainsOtherElfFrom || elfContainsOtherElfTo
	otherElfOverlapsElf := otherElfContainsElfFrom || otherElfContainsElfTo
	return elfOverlapsOtherElf || otherElfOverlapsElf
}
