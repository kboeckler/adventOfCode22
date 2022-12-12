package solution

import "errors"

func init() {
	RegisterSolution(12, day12{})
}

type day12 struct {
}

func (d day12) SolvePart1(input []string) interface{} {
	height := len(input)
	width := len(input[0])
	cells := make([]int32, 0, width*height)
	source, sink := -1, -1
	for y, row := range input {
		for x, cell := range row {
			if cell == int32('S') {
				source = d.coordToIndex(x, y, width)
				cell = int32('a')
			} else if cell == int32('E') {
				sink = d.coordToIndex(x, y, width)
				cell = int32('z')
			}
			cells = append(cells, cell)
		}
	}
	dist, _ := d.bfs(source, sink, width, height, cells)
	return dist
}

func (d day12) SolvePart2(input []string) interface{} {
	height := len(input)
	width := len(input[0])
	cells := make([]int32, 0, width*height)
	sink := -1
	for y, row := range input {
		for x, cell := range row {
			if cell == int32('S') {
				cell = int32('a')
			} else if cell == int32('E') {
				sink = d.coordToIndex(x, y, width)
				cell = int32('z')
			}
			cells = append(cells, cell)
		}
	}
	minDist := 9999
	for i, cell := range cells {
		if cell == int32('a') {
			dist, err := d.bfs(i, sink, width, height, cells)
			if err == nil && dist < minDist {
				minDist = dist
			}
		}
	}
	return minDist
}

func (d day12) bfs(source int, sink int, width int, height int, cells []int32) (int, error) {
	visited := make(map[int]bool)
	openCells := make([]int, 0)
	openCells = append(openCells, source)
	visited[source] = true
	iterations := -1
	done := false
	for len(openCells) > 0 {
		if done {
			break
		}
		iterations++
		cellsInThisIteration := len(openCells)
		for i := 0; i < cellsInThisIteration; i++ {
			current := openCells[0]
			openCells = openCells[1:]
			if current == sink {
				done = true
				break
			}
			for _, neighbor := range d.neighbors(current, width, height) {
				if !visited[neighbor] && cells[neighbor]-cells[current] <= 1 {
					visited[neighbor] = true
					openCells = append(openCells, neighbor)
				}
			}
		}
	}
	if !done {
		return -1, errors.New("no solution")
	}
	return iterations, nil
}

func (d day12) coordToIndex(x, y, width int) int {
	return y*width + x
}

func (d day12) indexToCoord(index, width int) (int, int) {
	return index % width, index / width
}

func (d day12) neighbors(cell int, width, height int) []int {
	neighbors := make([]int, 0)
	x, y := d.indexToCoord(cell, width)
	if x > 0 {
		neighbors = append(neighbors, d.coordToIndex(x-1, y, width))
	}
	if x < width-1 {
		neighbors = append(neighbors, d.coordToIndex(x+1, y, width))
	}
	if y > 0 {
		neighbors = append(neighbors, d.coordToIndex(x, y-1, width))
	}
	if y < height-1 {
		neighbors = append(neighbors, d.coordToIndex(x, y+1, width))
	}
	return neighbors
}
