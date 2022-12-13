package solution

import (
	"sort"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution(13, day13{})
}

type day13 struct {
}

func (d day13) SolvePart1(input []string) interface{} {
	elements := d.parseInput(input)
	rightOrders := make([]int, 0)
	for i := 0; i < len(elements)-1; i = i + 2 {
		leftElement := elements[i]
		rightElement := elements[i+1]
		if leftElement.compareTo(rightElement) <= 0 {
			rightOrders = append(rightOrders, i/2+1)
		}
	}
	result := 0
	for _, index := range rightOrders {
		result += index
	}
	return result
}

func (d day13) SolvePart2(input []string) interface{} {
	elements := d.parseInput(input)
	divider1 := d.parseElement("[[2]]")
	divider1.special = true
	divider2 := d.parseElement("[[6]]")
	divider2.special = true
	elements = append(elements, divider1)
	elements = append(elements, divider2)
	list := elementList{elements}
	sort.Sort(list)
	result := 1
	for i := range list.elements {
		if list.elements[i].special {
			result *= i + 1
		}
	}
	return result
}

func (d day13) parseInput(input []string) []*element {
	elements := make([]*element, 0)
	for _, row := range input {
		if len(row) == 0 {
			continue
		}
		element := d.parseElement(row)
		elements = append(elements, element)
	}
	return elements
}

func (d day13) parseElement(raw string) *element {
	if len(raw) == 0 {
		return &element{-1, make([]*element, 0), false}
	}
	value, err := strconv.Atoi(raw)
	if err == nil {
		return &element{value, nil, false}
	}
	elements := make([]*element, 0)
	innerRaw := raw[1 : len(raw)-1]
	innerElements := make([]string, 0)
	depth := 0
	token := strings.Builder{}
	for _, c := range innerRaw {
		if c == '[' {
			token.WriteRune(c)
			depth++
		} else if c == ']' {
			token.WriteRune(c)
			depth--
		} else if c == ',' {
			if depth == 0 {
				innerElements = append(innerElements, token.String())
				token = strings.Builder{}
			} else {
				token.WriteRune(c)
			}
		} else {
			token.WriteRune(c)
		}
	}
	innerElements = append(innerElements, token.String())
	for _, inner := range innerElements {
		elements = append(elements, d.parseElement(inner))
	}
	return &element{-1, elements, false}
}

type element struct {
	value    int
	elements []*element
	special  bool
}

func (ele *element) compareTo(other *element) int {
	if ele.value > -1 && other.value > -1 {
		return ele.value - other.value
	}
	if ele.value > -1 && other.value == -1 {
		helperElements := make([]*element, 1)
		helperElements[0] = ele
		helperElement := &element{-1, helperElements, false}
		return helperElement.compareTo(other)
	}
	if ele.value == -1 && other.value > -1 {
		helperElements := make([]*element, 1)
		helperElements[0] = other
		helperElement := &element{-1, helperElements, false}
		return ele.compareTo(helperElement)
	}
	size := len(ele.elements)
	if len(other.elements) > size {
		size = len(other.elements)
	}
	for i := 0; i < size; i++ {
		if len(other.elements)-1 < i {
			return 1
		}
		if len(ele.elements)-1 < i {
			return -1
		}
		result := ele.elements[i].compareTo(other.elements[i])
		if result == 0 {
			continue
		}
		return result
	}
	return 0
}

type elementList struct {
	elements []*element
}

func (e elementList) Len() int {
	return len(e.elements)
}

func (e elementList) Less(i, j int) bool {
	return e.elements[i].compareTo(e.elements[j]) <= 0
}

func (e elementList) Swap(i, j int) {
	temp := e.elements[i]
	e.elements[i] = e.elements[j]
	e.elements[j] = temp
}
