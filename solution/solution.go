package solution

var (
	solutions = make(map[int]Solution, 24)
)

func RegisterSolution(day int, solution Solution) {
	solutions[day] = solution
}

func GetSolutionFor(day int) Solution {
	return solutions[day]
}

func HasSolutionFor(day int) bool {
	_, exists := solutions[day]
	return exists
}

type Solution interface {
	SolvePart1(input []string) interface{}
	SolvePart2(input []string) interface{}
}
