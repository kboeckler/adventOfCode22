package solution

import (
	"fmt"
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
	for i := 0; i < 20; i++ {
		beginOverall := time.Now()
		for j := 0; j < len(monkeys); j++ {
			m := monkeys[j]
			for _, item := range m.items {
				begin := time.Now()
				itemNext := m.worryOp.apply(item)
				timeForWorry += time.Since(begin)
				begin = time.Now()
				itemNext.Div(itemNext, big.NewInt(3))
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

func (d day11) SolvePart2(input []string) interface{} {
	monkeys := make(map[int]monkey2)
	m := monkey2{id: -1}
	divisors := make([]int, 0)
	for _, row := range input {
		if d.containsString("Test", row) {
			divisor := d.findAllIntegers(row)[0]
			divisors = append(divisors, divisor)
		}
	}
	for _, row := range input {
		if d.containsString("Monkey", row) {
			if m.id != -1 {
				monkeys[m.id] = m
			}
			m.id = d.findAllIntegers(row)[0]
		}
		if d.containsString("Starting items", row) {
			m.items = make([]*itemContainer, 0)
			for _, item := range d.findAllIntegers(row) {
				m.items = append(m.items, d.createItem(item, divisors))
			}
		}
		if d.containsString("Operation", row) {
			m.worryOp = d.parseOperation2(row, divisors)
		}
		if d.containsString("Test", row) {
			divisor := d.findAllIntegers(row)[0]
			m.testFunc = func(worryLevel *itemContainer) bool {
				return worryLevel.isDivisibleBy(divisor)
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
	for i := 0; i < 10000; i++ {
		beginOverall := time.Now()
		for j := 0; j < len(monkeys); j++ {
			m := monkeys[j]
			for _, item := range m.items {
				begin := time.Now()
				m.worryOp.apply(item)
				timeForWorry += time.Since(begin)
				begin = time.Now()
				timeForRelief += time.Since(begin)
				begin = time.Now()
				if m.testFunc(item) {
					targetMonkey := monkeys[m.trueTarget]
					targetMonkey.items = append(targetMonkey.items, item)
					monkeys[m.trueTarget] = targetMonkey
				} else {
					targetMonkey := monkeys[m.falseTarget]
					targetMonkey.items = append(targetMonkey.items, item)
					monkeys[m.falseTarget] = targetMonkey
				}
				timeForThrowing += time.Since(begin)
				inspectionsByMonkey[m.id] = inspectionsByMonkey[m.id] + 1
			}
			m.items = make([]*itemContainer, 0)
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

func (d day11) parseOperation2(row string, divisors []int) operation2 {
	assignment := strings.ReplaceAll(row, "  Operation: new = old ", "")
	parts := strings.Split(assignment, " ")
	if strings.EqualFold(parts[1], "old") {
		return &squareOperation2{}
	} else if strings.EqualFold(parts[0], "*") {
		factor, _ := strconv.Atoi(parts[1])
		return &multOperation2{d.createItem(factor, divisors)}
	} else {
		summand, _ := strconv.Atoi(parts[1])
		return &sumOperation2{d.createItem(summand, divisors)}
	}
}

type monkey struct {
	id                      int
	items                   []*big.Int
	worryOp                 operation
	testFunc                func(*big.Int) bool
	trueTarget, falseTarget int
}

type monkey2 struct {
	id                      int
	items                   []*itemContainer
	worryOp                 operation2
	testFunc                func(container *itemContainer) bool
	trueTarget, falseTarget int
}

type operation interface {
	apply(value *big.Int) *big.Int
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

type operation2 interface {
	apply(value *itemContainer)
}

type sumOperation2 struct {
	summand *itemContainer
}

func (s *sumOperation2) apply(old *itemContainer) {
	old.add(s.summand)
}

type multOperation2 struct {
	factor *itemContainer
}

func (m *multOperation2) apply(old *itemContainer) {
	old.mult(m.factor)
}

type squareOperation2 struct {
}

func (s *squareOperation2) apply(old *itemContainer) {
	old.mult(old)
}

type itemContainer struct {
	values []*valueContainer
}

func (d day11) createItem(item int, divisors []int) *itemContainer {
	values := make([]*valueContainer, len(divisors))
	for i, div := range divisors {
		value := item % div
		values[i] = &valueContainer{value, div}
	}
	return &itemContainer{values}
}

func (ic *itemContainer) add(item *itemContainer) {
	for i, container := range ic.values {
		container.add(item.values[i])
	}
}

func (ic *itemContainer) mult(item *itemContainer) {
	for i, container := range ic.values {
		container.mult(item.values[i])
	}
}

func (ic *itemContainer) div(divisor int) {
	for _, container := range ic.values {
		container.div(divisor)
	}
}

func (ic *itemContainer) isDivisibleBy(divisor int) bool {
	for _, container := range ic.values {
		if container.divisor == divisor {
			return container.value == 0
		}
	}
	panic(fmt.Sprintf("No matching divisor found for %d.", divisor))
}

type valueContainer struct {
	value, divisor int
}

func (vc *valueContainer) add(item *valueContainer) {
	vc.value = (vc.value + item.value) % vc.divisor
}

func (vc *valueContainer) mult(item *valueContainer) {
	vc.value = (vc.value * item.value) % vc.divisor
}

func (vc *valueContainer) div(divisor int) {
	vc.value = (vc.value / divisor) % vc.divisor
}
