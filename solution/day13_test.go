package solution

import (
	"strings"
	"testing"
)

func TestExample13_1(t *testing.T) {
	inputStr := "[1,1,3,1,1]\n[1,1,5,1,1]\n\n[[1],[2,3,4]]\n[[1],4]\n\n[9]\n[[8,7,6]]\n\n[[4,4],4,4]\n[[4,4],4,4,4]\n\n[7,7,7,7]\n[7,7,7]\n\n[]\n[3]\n\n[[[]]]\n[[]]\n\n[1,[2,[3,[4,[5,6,7]]]],8,9]\n[1,[2,[3,[4,[5,6,0]]]],8,9]"
	input := strings.Split(inputStr, "\n")
	result := day13{}.SolvePart1(input)
	if result != 13 {
		t.Errorf("Expected %d to be %d, but was not", result, 13)
	}
}

func TestExample13_2(t *testing.T) {
	inputStr := "[1,1,3,1,1]\n[1,1,5,1,1]\n\n[[1],[2,3,4]]\n[[1],4]\n\n[9]\n[[8,7,6]]\n\n[[4,4],4,4]\n[[4,4],4,4,4]\n\n[7,7,7,7]\n[7,7,7]\n\n[]\n[3]\n\n[[[]]]\n[[]]\n\n[1,[2,[3,[4,[5,6,7]]]],8,9]\n[1,[2,[3,[4,[5,6,0]]]],8,9]"
	input := strings.Split(inputStr, "\n")
	result := day13{}.SolvePart2(input)
	if result != 140 {
		t.Errorf("Expected %d to be %d, but was not", result, 140)
	}
}

func TestCompareTo1(t *testing.T) {
	leftElements := make([]*element, 0)
	leftElements = append(leftElements, &element{1, nil, false})
	leftElements = append(leftElements, &element{1, nil, false})
	leftElements = append(leftElements, &element{3, nil, false})
	leftElements = append(leftElements, &element{1, nil, false})
	leftElements = append(leftElements, &element{1, nil, false})
	left := &element{-1, leftElements, false}
	rightElements := make([]*element, 0)
	rightElements = append(rightElements, &element{1, nil, false})
	rightElements = append(rightElements, &element{1, nil, false})
	rightElements = append(rightElements, &element{5, nil, false})
	rightElements = append(rightElements, &element{1, nil, false})
	rightElements = append(rightElements, &element{1, nil, false})
	right := &element{-1, rightElements, false}
	result := left.compareTo(right)
	if result >= 0 {
		t.Errorf("Expected left to be smaller than right, but was not")
	}
}

func TestCompareTo2(t *testing.T) {
	leftElements := make([]*element, 0)
	leftInnerElement := make([]*element, 0)
	leftInnerElement = append(leftInnerElement, &element{1, nil, false})
	leftElements = append(leftElements, &element{-1, leftInnerElement, false})
	secondLeftInnerElement := make([]*element, 0)
	secondLeftInnerElement = append(secondLeftInnerElement, &element{2, nil, false})
	secondLeftInnerElement = append(secondLeftInnerElement, &element{3, nil, false})
	secondLeftInnerElement = append(secondLeftInnerElement, &element{4, nil, false})
	leftElements = append(leftElements, &element{-1, secondLeftInnerElement, false})
	left := &element{-1, leftElements, false}
	rightElements := make([]*element, 0)
	rightInnerElement := make([]*element, 0)
	rightInnerElement = append(rightInnerElement, &element{1, nil, false})
	rightElements = append(rightElements, &element{-1, rightInnerElement, false})
	rightElements = append(rightElements, &element{4, nil, false})
	right := &element{-1, rightElements, false}
	result := left.compareTo(right)
	if result >= 0 {
		t.Errorf("Expected left to be smaller than right, but was not")
	}
}

func TestCompareTo5(t *testing.T) {
	leftElements := make([]*element, 0)
	leftElements = append(leftElements, &element{7, nil, false})
	leftElements = append(leftElements, &element{7, nil, false})
	leftElements = append(leftElements, &element{7, nil, false})
	leftElements = append(leftElements, &element{7, nil, false})
	left := &element{-1, leftElements, false}
	rightElements := make([]*element, 0)
	rightElements = append(rightElements, &element{7, nil, false})
	rightElements = append(rightElements, &element{7, nil, false})
	rightElements = append(rightElements, &element{7, nil, false})
	right := &element{-1, rightElements, false}
	result := left.compareTo(right)
	if result < 0 {
		t.Errorf("Expected right to be smaller than left, but was not")
	}
}
