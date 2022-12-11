package solution

import (
	"math/big"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func init() {
	RegisterSolution(11, day11{})
}

type day11 struct {
}

func (d day11) SolvePart1(input []string) interface{} {
	return d.solveWithRelief(input, big.NewInt(int64(3)), 20)
}

func (d day11) SolvePart2(input []string) interface{} {
	return d.solveWithRelief(input, big.NewInt(int64(1)), 10000)
}

func (d day11) solveWithRelief(input []string, reliefDivisor *big.Int, rounds int) interface{} {
	if rounds > 20 {
		return "TIMEOUT"
	}
	monkeys := make(map[int]monkey)
	m := monkey{id: -1}
	for _, row := range input {
		if d.containsString("Monkey", row) {
			if m.id != -1 {
				monkeys[m.id] = m
			}
			m.id = d.findAllIntegers(row)[0]
		}
		if d.containsString("Starting items", row) {
			m.items = make([]*big.Int, 0)
			for _, item := range d.findAllIntegers(row) {
				m.items = append(m.items, big.NewInt(int64(item)))
			}
		}
		if d.containsString("Operation", row) {
			m.worryOp = d.parseOperation(row)
		}
		if d.containsString("Test", row) {
			divisor := big.NewInt(int64(d.findAllIntegers(row)[0]))
			m.testFunc = func(worryLevel *big.Int) bool {
				result := big.NewInt(0)
				result.Mod(worryLevel, divisor)
				return result.Int64() == 0
			}
		}
		if d.containsString("If true", row) {
			m.trueTarget = d.findAllIntegers(row)[0]
		}
		if d.containsString("If false", row) {
			m.falseTarget = d.findAllIntegers(row)[0]
		}
	}
	monkeys[m.id] = m
	inspectionsByMonkey := make(map[int]int)
	var timeOverall time.Duration = 0
	var timeForWorry time.Duration = 0
	var timeForThrowing time.Duration = 0
	var timeForRelief time.Duration = 0
	for i := 0; i < rounds; i++ {
		beginOverall := time.Now()
		for j := 0; j < len(monkeys); j++ {
			m := monkeys[j]
			for _, item := range m.items {
				begin := time.Now()
				itemNext := m.worryOp.apply(item)
				timeForWorry += time.Since(begin)
				begin = time.Now()
				itemNext.Div(itemNext, reliefDivisor)
				timeForRelief += time.Since(begin)
				begin = time.Now()
				if m.testFunc(itemNext) {
					targetMonkey := monkeys[m.trueTarget]
					targetMonkey.items = append(targetMonkey.items, itemNext)
					monkeys[m.trueTarget] = targetMonkey
				} else {
					targetMonkey := monkeys[m.falseTarget]
					targetMonkey.items = append(targetMonkey.items, itemNext)
					monkeys[m.falseTarget] = targetMonkey
				}
				timeForThrowing += time.Since(begin)
				inspectionsByMonkey[m.id] = inspectionsByMonkey[m.id] + 1
			}
			m.items = make([]*big.Int, 0)
			monkeys[m.id] = m
		}
		timeOverall += time.Since(beginOverall)
	}
	inspections := make([]int, 0)
	for _, value := range inspectionsByMonkey {
		inspections = append(inspections, value)
	}
	sort.Ints(inspections)
	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func (d day11) containsString(pattern string, row string) bool {
	return len(strings.Replace(row, pattern, "", 1)) < len(row)
}

func (d day11) findAllIntegers(row string) []int {
	allIntegers := make([]int, 0)
	reg, _ := regexp.Compile("([0-9]+)(,[0-9]+)*")
	parts := reg.FindAllStringSubmatch(row, -1)
	if parts != nil {
		for i := 0; i < len(parts); i++ {
			integer, _ := strconv.Atoi(parts[i][1])
			allIntegers = append(allIntegers, integer)
		}
	}
	return allIntegers
}

type monkey struct {
	id                      int
	items                   []*big.Int
	worryOp                 operation
	testFunc                func(*big.Int) bool
	trueTarget, falseTarget int
}

type operation interface {
	apply(value *big.Int) *big.Int
}

func (d day11) parseOperation(row string) operation {
	assignment := strings.ReplaceAll(row, "  Operation: new = old ", "")
	parts := strings.Split(assignment, " ")
	if strings.EqualFold(parts[1], "old") {
		return &squareOperation{}
	} else if strings.EqualFold(parts[0], "*") {
		factor, _ := strconv.Atoi(parts[1])
		return &multOperation{big.NewInt(int64(factor))}
	} else {
		summand, _ := strconv.Atoi(parts[1])
		return &sumOperation{big.NewInt(int64(summand))}
	}
}

type sumOperation struct {
	summand *big.Int
}

func (s *sumOperation) apply(old *big.Int) *big.Int {
	return old.Add(old, s.summand)
}

type multOperation struct {
	factor *big.Int
}

func (m *multOperation) apply(old *big.Int) *big.Int {
	return old.Mul(old, m.factor)
}

type squareOperation struct {
}

func (s *squareOperation) apply(old *big.Int) *big.Int {
	return old.Mul(old, old)
}
