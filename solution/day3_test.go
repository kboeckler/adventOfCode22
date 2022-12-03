package solution

import (
	"testing"
)

func TestExample1(t *testing.T) {
	input := make([]string, 0)
	input = append(input, "vJrwpWtwJgWrhcsFMMfFFhFp")
	input = append(input, "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")
	input = append(input, "PmmdzqPrVvPwwTWBwg")
	input = append(input, "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn")
	input = append(input, "ttgJtRGJQctTZtZT")
	input = append(input, "CrZsJsPPZsGzwwsLwLmpwMDw")
	result := day3{}.SolvePart1(input)
	if result != 157 {
		t.Errorf("Expected %d to be %d, but was not", result, 157)
	}
}

func TestExample2(t *testing.T) {
	input := make([]string, 0)
	input = append(input, "vJrwpWtwJgWrhcsFMMfFFhFp")
	input = append(input, "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL")
	input = append(input, "PmmdzqPrVvPwwTWBwg")
	input = append(input, "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn")
	input = append(input, "ttgJtRGJQctTZtZT")
	input = append(input, "CrZsJsPPZsGzwwsLwLmpwMDw")
	result := day3{}.SolvePart2(input)
	if result != 70 {
		t.Errorf("Expected %d to be %d, but was not", result, 70)
	}
}
