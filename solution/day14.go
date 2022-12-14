package solution

import (
	"strconv"
	"strings"
)

func init() {
	RegisterSolution(14, day14{})
}

type day14 struct {
}

func (d day14) SolvePart1(input []string) interface{} {
	height, blocked := d.parseCave(input)
	return d.fillCave(blocked, height)
}

func (d day14) SolvePart2(input []string) interface{} {
	height, blocked := d.parseCave(input)
	floorFrom := &vertex{499 - height, height + 1}
	floorTo := &vertex{501 + height, height + 1}
	for _, vertex := range floorFrom.lineTo(floorTo) {
		blocked[*vertex] = true
	}
	return d.fillCave(blocked, height+2)
}

func (d day14) parseCave(input []string) (int, map[vertex]bool) {
	height := -1
	blocked := make(map[vertex]bool)
	for _, row := range input {
		vertices := strings.Split(row, "->")
		startVertex := d.parseVertex(vertices[0])
		for i := 1; i < len(vertices); i++ {
			endVertex := d.parseVertex(vertices[i])
			for _, vertex := range startVertex.lineTo(endVertex) {
				blocked[*vertex] = true
				if vertex.y+1 > height {
					height = vertex.y + 1
				}
			}
			startVertex = endVertex
		}
	}
	return height, blocked
}

func (d day14) parseVertex(raw string) *vertex {
	split := strings.Split(strings.TrimSpace(raw), ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	return &vertex{x, y}
}

func (d day14) fillCave(blocked map[vertex]bool, height int) interface{} {
	startVertex := &vertex{500, 0}
	currentVertex := startVertex
	counter := 0
	for {
		hasStopped := false
		for {
			nextVertex := currentVertex.nextVertex(blocked)
			if nextVertex == nil {
				hasStopped = true
				break
			} else if nextVertex.y == height-1 {
				break
			} else {
				currentVertex = nextVertex
			}
		}
		if hasStopped {
			counter++
			blocked[*currentVertex] = true
			if *currentVertex == *startVertex {
				break
			}
			currentVertex = startVertex
		} else {
			break
		}
	}
	return counter
}

type vertex struct {
	x, y int
}

func (v *vertex) bottom() *vertex {
	return &vertex{v.x, v.y + 1}
}

func (v *vertex) bottomLeft() *vertex {
	return &vertex{v.x - 1, v.y + 1}
}

func (v *vertex) bottomRight() *vertex {
	return &vertex{v.x + 1, v.y + 1}
}

func (v *vertex) nextVertex(blocked map[vertex]bool) *vertex {
	if !blocked[*v.bottom()] {
		return v.bottom()
	}
	if !blocked[*v.bottomLeft()] {
		return v.bottomLeft()
	}
	if !blocked[*v.bottomRight()] {
		return v.bottomRight()
	}
	return nil
}

func (v *vertex) lineTo(end *vertex) []*vertex {
	vertices := make([]*vertex, 0)
	if v.x != end.x {
		incr := 1
		if v.x > end.x {
			incr = -1
		}
		for x := v.x; x != end.x; x = x + incr {
			vertices = append(vertices, &vertex{x, v.y})
		}
		vertices = append(vertices, &vertex{end.x, v.y})
	} else {
		incr := 1
		if v.y > end.y {
			incr = -1
		}
		for y := v.y; y != end.y; y = y + incr {
			vertices = append(vertices, &vertex{v.x, y})
		}
		vertices = append(vertices, &vertex{v.x, end.y})
	}
	return vertices
}
