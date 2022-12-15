package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/kboeckler/adventOfCode22/solution"
	"os"
	"reflect"
	"strings"
	"time"
)

func main() {
	parser := argparse.NewParser("adventOfCode22", "Prints solutions of Advent of Code 2022")
	day := *parser.Int("d", "day", &argparse.Options{Required: false, Help: "one specific day to solve"})
	inputFolder := parser.String("i", "input", &argparse.Options{Required: false, Help: "input folder of puzzle input", Default: "input"})
	shortPrint := parser.Flag("s", "short", &argparse.Options{Required: false, Help: "Prints the results in a short format"})
	timeTracking := parser.Flag("t", "time", &argparse.Options{Required: false, Help: "Tracks and prints the time needed for each day"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(-1)
	}
	fileInfo, err := os.Stat(*inputFolder)
	if err != nil || !fileInfo.IsDir() {
		if strings.EqualFold(*inputFolder, "input") {
			fmt.Printf("Cannot read <%s> as input directory. Consider using -i argument to pass the correct folder.\n", *inputFolder)
		} else {
			fmt.Printf("Cannot read <%s> as input directory.\n", *inputFolder)
		}
		os.Exit(-5)
	}
	var printSrc printing
	if *shortPrint {
		printSrc = short{*timeTracking}
	} else {
		printSrc = verbose{*timeTracking}
	}
	allDays := make([]int, 0, 24)
	if day > 0 {
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

	fmt.Printf(printSrc.welcomeMsg())
	before := time.Now()
	for _, day := range allDays {
		fmt.Printf(printSrc.startSolving(day))
		solutionForDay := solution.GetSolutionFor(day)
		solutionName := findSimpleTypeName(solutionForDay)
		solutionInputFilename := fmt.Sprintf("%s/%s.txt", *inputFolder, solutionName)
		input, err := os.ReadFile(solutionInputFilename)
		if err != nil {
			fmt.Printf("Error reading input from file %s: %v\n", solutionInputFilename, err)
			os.Exit(-4)
		}
		inputAsRows := strings.Split(strings.ReplaceAll(string(input), "\r\n", "\n"), "\n")
		fmt.Printf(printSrc.startPart1())
		before1 := time.Now()
		result1 := solutionForDay.SolvePart1(inputAsRows)
		time1 := time.Since(before1)
		fmt.Printf(printSrc.result1(result1))
		fmt.Printf(printSrc.startPart2())
		before2 := time.Now()
		result2 := solutionForDay.SolvePart2(inputAsRows)
		time2 := time.Since(before2)
		fmt.Printf(printSrc.result2(result2))
		fmt.Printf(printSrc.times(time1, time2))
		fmt.Printf("\n")
	}
	timeAll := time.Since(before)
	fmt.Printf(printSrc.timeAll(timeAll))
}

func findSimpleTypeName(solution solution.Solution) string {
	solutionName := reflect.TypeOf(solution)
	return solutionName.Name()
}

type printing interface {
	welcomeMsg() string
	startSolving(day int) string
	startPart1() string
	startPart2() string
	result1(result interface{}) string
	result2(result interface{}) string
	times(time1 time.Duration, time2 time.Duration) string
	timeAll(timeAll time.Duration) string
}

type verbose struct {
	timeTracking bool
}

func (v verbose) timeAll(timeAll time.Duration) string {
	if v.timeTracking {
		return fmt.Sprintf("Total time needed: %v\n", timeAll)
	}
	return ""
}

func (v verbose) times(time1 time.Duration, time2 time.Duration) string {
	if v.timeTracking {
		return fmt.Sprintf(" (%v+%v=%v)", time1, time2, time1+time2)
	}
	return ""
}

func (v verbose) welcomeMsg() string {
	return "Welcome to Advent of Code 22\n###############################\n"
}

func (v verbose) startSolving(day int) string {
	return fmt.Sprintf("Solving day %d #", day)
}

func (v verbose) startPart1() string {
	return " Part1:... "
}

func (v verbose) startPart2() string {
	return " Part2:... "
}

func (v verbose) result1(result interface{}) string {
	return fmt.Sprintf("%v", result)
}

func (v verbose) result2(result interface{}) string {
	return fmt.Sprintf("%v", result)
}

type short struct {
	timeTracking bool
}

func (s short) timeAll(timeAll time.Duration) string {
	if s.timeTracking {
		return fmt.Sprintf("(%v)\n", timeAll)
	}
	return ""
}

func (s short) times(time1 time.Duration, time2 time.Duration) string {
	if s.timeTracking {
		return fmt.Sprintf(" (%v)", time1+time2)
	}
	return ""
}

func (s short) welcomeMsg() string {
	return ""
}

func (s short) startSolving(day int) string {
	return fmt.Sprintf("%d", day)
}

func (s short) startPart1() string {
	return " "
}

func (s short) startPart2() string {
	return " "
}

func (s short) result1(result interface{}) string {
	return fmt.Sprintf("%v", result)
}

func (s short) result2(result interface{}) string {
	return fmt.Sprintf("%v", result)
}
