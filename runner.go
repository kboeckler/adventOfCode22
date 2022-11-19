package main

import (
	"fmt"
	"github.com/kboeckler/adventOfCode22/solution"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	allDays := make([]int, 0, 24)
	if len(os.Args) > 1 {
		day, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Argument is not an integer\n")
			os.Exit(-1)
		}
		if day <= 0 || day > 24 {
			fmt.Printf("Day %d is not between 1 and 24\n", day)
			os.Exit(-2)
		}
		if !solution.HasSolutionFor(day) {
			fmt.Printf("Day %d has no programmed solution\n", day)
			os.Exit(-3)
		}
		allDays = append(allDays, day)
	} else {
		for day := 1; day <= 24; day++ {
			if solution.HasSolutionFor(day) {
				allDays = append(allDays, day)
			}
		}
	}

	fmt.Println("Welcome to Advent of Code 22")
	fmt.Println("################")
	for _, day := range allDays {
		fmt.Printf("Solving day %d # ", day)
		solutionForDay := solution.GetSolutionFor(day)
		solutionName := findSimpleTypeName(solutionForDay)
		solututionInputFilename := fmt.Sprintf("input/%s.txt", solutionName)
		input, err := os.ReadFile(solututionInputFilename)
		if err != nil {
			fmt.Printf("Error reading input from file %s: %v\n", solututionInputFilename, err)
			os.Exit(-4)
		}
		inputAsRows := strings.Split(strings.ReplaceAll(string(input), "\r\n", "\n"), "\n")
		fmt.Printf("Part1:... ")
		result1 := solutionForDay.SolvePart1(inputAsRows)
		fmt.Printf("Solution Part1: %v Part2:... ", result1)
		result2 := solutionForDay.SolvePart2(inputAsRows)
		fmt.Printf("Solution Part2: %v", result2)
		fmt.Printf("\n")
	}
}

func findSimpleTypeName(solution solution.Solution) string {
	solutionName := reflect.TypeOf(solution)
	return solutionName.Name()
}
