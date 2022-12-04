package solution

import "testing"

func TestExample4_1(t *testing.T) {
	input := make([]string, 0)
	input = append(input, "2-4,6-8")
	input = append(input, "2-3,4-5")
	input = append(input, "5-7,7-9")
	input = append(input, "2-8,3-7")
	input = append(input, "6-6,4-6")
	input = append(input, "2-6,4-8")
	result := day4{}.SolvePart1(input)
	if result != 2 {
		t.Errorf("Expected %d to be %d, but was not", result, 2)
	}
}

func TestExampleNoContain(t *testing.T) {
	input := make([]string, 0)
	input = append(input, "2-4,6-8")
	result := day4{}.SolvePart1(input)
	if result != 0 {
		t.Errorf("Expected %d to be %d, but was not", result, 0)
	}
}

func TestExampleOneContain(t *testing.T) {
	input := make([]string, 0)
	input = append(input, "2-8,3-7")
	result := day4{}.SolvePart1(input)
	if result != 1 {
		t.Errorf("Expected %d to be %d, but was not", result, 1)
	}
}

func TestNoContainTwoDigitNumbers(t *testing.T) {
	input := make([]string, 0)
	input = append(input, "1-1,11-12")
	result := day4{}.SolvePart1(input)
	if result != 0 {
		t.Errorf("Expected %d to be %d, but was not", result, 0)
	}
}

func TestExample4_2(t *testing.T) {
	input := make([]string, 0)
	input = append(input, "2-4,6-8")
	input = append(input, "2-3,4-5")
	input = append(input, "5-7,7-9")
	input = append(input, "2-8,3-7")
	input = append(input, "6-6,4-6")
	input = append(input, "2-6,4-8")
	result := day4{}.SolvePart2(input)
	if result != 4 {
		t.Errorf("Expected %d to be %d, but was not", result, 4)
	}
}
