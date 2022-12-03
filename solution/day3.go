package solution

import (
	"fmt"
	"strings"
)

func init() {
	RegisterSolution(3, day3{})
}

type day3 struct {
}

func (d day3) SolvePart1(input []string) interface{} {
	prioritySum := 0
	for _, row := range input {
		leftBagUnresolved := make(map[uint8]bool)
		for i := 0; i < len(row)/2; i++ {
			leftBagUnresolved[row[i]] = true
		}
		for i := len(row) / 2; i < len(row); i++ {
			if leftBagUnresolved[row[i]] {
				leftBagUnresolved[row[i]] = false
				prioritySum += prioritzeType(row[i])
			}
		}
	}
	return prioritySum
}

func (d day3) SolvePart2(input []string) interface{} {
	prioritySum := 0
	for g := 0; g < len(input)/3; g++ {
		contentInAllBags := make(map[uint8]string)
		for b, row := range input[g*3 : (g+1)*3] {
			for i := 0; i < len(row); i++ {
				contentInAllBags[row[i]] = fmt.Sprintf("%s%d", contentInAllBags[row[i]], b)
			}
		}
		for ch, content := range contentInAllBags {
			if strings.Contains(content, "0") && strings.Contains(content, "1") && strings.Contains(content, "2") {
				prioritySum += prioritzeType(ch)
			}
		}
	}
	return prioritySum
}

func prioritzeType(itemType uint8) int {
	val := itemType - 38
	if itemType >= 97 {
		val = itemType - 96
	}
	return int(val)
}
