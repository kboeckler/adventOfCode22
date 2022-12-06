package solution

func init() {
	RegisterSolution(6, day6{})
}

type day6 struct {
}

func (d day6) SolvePart1(input []string) interface{} {
	return findIndexAfterSequence(input, 4)
}

func (d day6) SolvePart2(input []string) interface{} {
	return findIndexAfterSequence(input, 14)
}

func findIndexAfterSequence(input []string, length int) interface{} {
	for idx := range input[0] {
		letters := make(map[int32]bool)
		for _, letter := range input[0][idx : idx+length] {
			exists := letters[letter]
			if exists {
				break
			}
			letters[letter] = true
		}
		if len(letters) == length {
			return idx + length
		}
	}
	panic("No solution found")
}
