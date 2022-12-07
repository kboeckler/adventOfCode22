package solution

import (
	"math/big"
	"testing"
)

func TestExample7_1(t *testing.T) {
	input := make([]string, 0)
	input = append(input, "$ cd /")
	input = append(input, "$ ls")
	input = append(input, "dir a")
	input = append(input, "14848514 b.txt")
	input = append(input, "8504156 c.dat")
	input = append(input, "dir d")
	input = append(input, "$ cd a")
	input = append(input, "$ ls")
	input = append(input, "dir e")
	input = append(input, "29116 f")
	input = append(input, "2557 g")
	input = append(input, "62596 h.lst")
	input = append(input, "$ cd e")
	input = append(input, "$ ls")
	input = append(input, "584 i")
	input = append(input, "$ cd ..")
	input = append(input, "$ cd ..")
	input = append(input, "$ cd d")
	input = append(input, "$ ls")
	input = append(input, "4060174 j")
	input = append(input, "8033020 d.log")
	input = append(input, "5626152 d.ext")
	input = append(input, "7214296 k")
	result := day7{}.SolvePart1(input)
	resultInt := result.(*big.Int)
	if resultInt.Cmp(big.NewInt(95437)) != 0 {
		t.Errorf("Expected %d to be %d, but was not", result, 95437)
	}
}

func TestExample7_2(t *testing.T) {
	input := make([]string, 0)
	input = append(input, "$ cd /")
	input = append(input, "$ ls")
	input = append(input, "dir a")
	input = append(input, "14848514 b.txt")
	input = append(input, "8504156 c.dat")
	input = append(input, "dir d")
	input = append(input, "$ cd a")
	input = append(input, "$ ls")
	input = append(input, "dir e")
	input = append(input, "29116 f")
	input = append(input, "2557 g")
	input = append(input, "62596 h.lst")
	input = append(input, "$ cd e")
	input = append(input, "$ ls")
	input = append(input, "584 i")
	input = append(input, "$ cd ..")
	input = append(input, "$ cd ..")
	input = append(input, "$ cd d")
	input = append(input, "$ ls")
	input = append(input, "4060174 j")
	input = append(input, "8033020 d.log")
	input = append(input, "5626152 d.ext")
	input = append(input, "7214296 k")
	result := day7{}.SolvePart2(input)
	resultInt := result.(*big.Int)
	if resultInt.Cmp(big.NewInt(24933642)) != 0 {
		t.Errorf("Expected %d to be %d, but was not", result, 24933642)
	}
}

func TestExample7_SimpleTree(t *testing.T) {
	input := make([]string, 0)
	input = append(input, "$ cd /")
	input = append(input, "$ cd a")
	input = append(input, "$ cd b")
	input = append(input, "$ ls")
	input = append(input, "1 apfel")
	result := day7{}.SolvePart1(input)
	resultInt := result.(*big.Int)
	if resultInt.Cmp(big.NewInt(1)) != 1 {
		t.Errorf("Expected %d to be %d, but was not", 1, 1)
	}
}
