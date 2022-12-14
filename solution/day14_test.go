package solution

import (
	"strings"
	"testing"
)

func TestExample14_1(t *testing.T) {
	inputStr := "498,4 -> 498,6 -> 496,6\n503,4 -> 502,4 -> 502,9 -> 494,9"
	input := strings.Split(inputStr, "\n")
	result := day14{}.SolvePart1(input)
	if result != 24 {
		t.Errorf("Expected %d to be %d, but was not", result, 24)
	}
}
