package solution

import "testing"

func TestExample17_1(t *testing.T) {
	input := []string{">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"}
	result := day17{}.SolvePart1(input)
	if result != int64(3068) {
		t.Errorf("Exptected %d to be %d, but was not", result, 3068)
	}
}

func TestExample17_2(t *testing.T) {
	input := []string{">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"}
	result := day17{}.SolvePart2(input)
	if result != int64(1514285714288) {
		t.Errorf("Exptected %d to be %d, but was not", result, 1_514_285_714_288)
	}
}
