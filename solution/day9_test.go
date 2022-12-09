package solution

import "testing"

func TestExample9_1(t *testing.T) {
	input := make([]string, 0)
	input = append(input, "R 4")
	input = append(input, "U 4")
	input = append(input, "L 3")
	input = append(input, "D 1")
	input = append(input, "R 4")
	input = append(input, "D 1")
	input = append(input, "L 5")
	input = append(input, "R 2")
	result := day9{}.SolvePart1(input)
	if result != 13 {
		t.Errorf("Expected %d to be %d, but was not", result, 13)
	}
}

func TestBackAndForth9_1(t *testing.T) {
	input := make([]string, 0)
	input = append(input, "R 4")
	input = append(input, "L 4")
	result := day9{}.SolvePart1(input)
	if result != 4 {
		t.Errorf("Expected %d to be %d, but was not", result, 4)
	}
}

func TestDiagonally9_1(t *testing.T) {
	input := make([]string, 0)
	input = append(input, "R 1")
	input = append(input, "U 1")
	input = append(input, "R 1")
	input = append(input, "U 1")
	result := day9{}.SolvePart1(input)
	if result != 2 {
		t.Errorf("Expected %d to be %d, but was not", result, 2)
	}
}
