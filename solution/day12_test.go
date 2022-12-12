package solution

import (
	"strings"
	"testing"
)

func TestExample12_1(t *testing.T) {
	inputStr := "Sabqponm\nabcryxxl\naccszExk\nacctuvwj\nabdefghi"
	input := strings.Split(inputStr, "\n")
	result := day12{}.SolvePart1(input)
	if result != 31 {
		t.Errorf("Expected %d to be %d, but was not", result, 31)
	}
}

func TestSimple12(t *testing.T) {
	input := make([]string, 0)
	input = append(input, "SbcdefghijklmnopqrstuvwxyE")
	result := day12{}.SolvePart1(input)
	if result != 25 {
		t.Errorf("Expected %d to be %d, but was not", result, 25)
	}
}

func TestMedium12(t *testing.T) {
	input := make([]string, 0)
	input = append(input, "Sa")
	input = append(input, "bb")
	input = append(input, "cd")
	input = append(input, "fe")
	input = append(input, "gh")
	input = append(input, "ji")
	input = append(input, "kl")
	input = append(input, "nm")
	input = append(input, "op")
	input = append(input, "rq")
	input = append(input, "st")
	input = append(input, "vu")
	input = append(input, "wx")
	input = append(input, "zy")
	input = append(input, "zE")
	result := day12{}.SolvePart1(input)
	if result != 25 {
		t.Errorf("Expected %d to be %d, but was not", result, 25)
	}
}
