package solution

import "testing"

func TestExample8_1(t *testing.T) {
	input := make([]string, 0)
	input = append(input, "30373")
	input = append(input, "25512")
	input = append(input, "65332")
	input = append(input, "33549")
	input = append(input, "35390")
	result := day8{}.SolvePart1(input)
	if result != 21 {
		t.Errorf("Expected %d to be %d, but was not", result, 21)
	}
}

func TestExample8_2(t *testing.T) {
	input := make([]string, 0)
	input = append(input, "30373")
	input = append(input, "25512")
	input = append(input, "65332")
	input = append(input, "33549")
	input = append(input, "35390")
	result := day8{}.SolvePart2(input)
	if result != 8 {
		t.Errorf("Expected %d to be %d, but was not", result, 8)
	}
}
