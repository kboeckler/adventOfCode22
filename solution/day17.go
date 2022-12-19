package solution

import (
	"fmt"
	"time"
)

func init() {
	RegisterSolution(17, day17{})
}

type day17 struct {
}

func (d day17) SolvePart1(input []string) interface{} {
	return d.solveWithAmount(input, 2022)
}

func (d day17) SolvePart2(input []string) interface{} {
	return d.solveWithAmount(input, 1_000_000_000_000)
}

func (d day17) solveWithAmount(input []string, amountRocks int64) interface{} {
	streams := input[0]
	forms := []*tetris{d.createMinus(), d.createPlus(), d.createL(), d.createBar(), d.createQuad()}
	streamIndex := 0
	formIndex := 0
	maxHeight := int64(-1)
	blocks := make(map[tetrisPos]bool)
	blocks2 := make(map[tetrisPos]bool)
	timeStep := time.Now()
	for i := 0; int64(i) < amountRocks; i++ {
		if i%1_000_000 == 0 {
			fmt.Print("\033[2K\r")
			fmt.Printf("Block # %d", i)
			fmt.Printf(" | Time: %dms", time.Since(timeStep).Milliseconds())
		}
		form := forms[formIndex].new(maxHeight)
		formIndex = (formIndex + 1) % len(forms)
		for {
			stream := streams[streamIndex]
			streamIndex = (streamIndex + 1) % len(streams)
			form.push(stream, &blocks)
			couldDrop := form.drop(&blocks)
			if !couldDrop {
				for _, block := range form.getBlocks() {
					if i > 100_000 {
						blocks2[*block] = true
					}
					blocks[*block] = true
					if block.y > maxHeight {
						maxHeight = block.y
					}
				}
				break
			}
		}
		if i > 200_000 && i%100_000 == 0 {
			blocks = blocks2
			blocks2 = make(map[tetrisPos]bool)
		}
	}
	fmt.Println()
	return maxHeight + 1
}

type tetris struct {
	pos             *tetrisPos
	blocks          []*tetrisPos
	moveRightBlocks []int
	moveLeftBlocks  []int
	moveDownBlocks  []int
}

func (t *tetris) new(maxHeight int64) *tetris {
	return &tetris{&tetrisPos{2, maxHeight + 4}, t.blocks, t.moveRightBlocks, t.moveLeftBlocks, t.moveDownBlocks}
}

func (t *tetris) push(stream uint8, blocks *map[tetrisPos]bool) {
	direction := int64(1)
	moveBlocks := t.moveRightBlocks
	if stream == '<' {
		direction = -1
		moveBlocks = t.moveLeftBlocks
	}
	canMove := true
	for _, blockIndex := range moveBlocks {
		block := t.blocks[blockIndex]
		pushedBlock := tetrisPos{t.pos.x + block.x + direction, t.pos.y + block.y}
		if pushedBlock.x < 0 || pushedBlock.x >= 7 || contains(blocks, &pushedBlock) {
			canMove = false
			break
		}
	}
	if canMove {
		t.pos = &tetrisPos{t.pos.x + direction, t.pos.y}
	}
}

func contains(blocks *map[tetrisPos]bool, pushedBlock *tetrisPos) bool {
	result := (*blocks)[*pushedBlock]
	return result
}

func (t *tetris) drop(blocks *map[tetrisPos]bool) bool {
	canMove := true
	for _, blockIndex := range t.moveDownBlocks {
		block := t.blocks[blockIndex]
		pushedBlock := tetrisPos{t.pos.x + block.x, t.pos.y + block.y - 1}
		if pushedBlock.y < 0 || contains(blocks, &pushedBlock) {
			canMove = false
			break
		}
	}
	if canMove {
		t.pos = &tetrisPos{t.pos.x, t.pos.y - 1}
		return true
	}
	return false
}

func (t *tetris) getBlocks() []*tetrisPos {
	blocks := make([]*tetrisPos, 0, len(t.blocks))
	for _, block := range t.blocks {
		blocks = append(blocks, &tetrisPos{t.pos.x + block.x, t.pos.y + block.y})
	}
	return blocks
}

type tetrisPos struct {
	x, y int64
}

func (d day17) createMinus() *tetris {
	blocks := make([]*tetrisPos, 4)
	blocks[0] = &tetrisPos{0, 0}
	blocks[1] = &tetrisPos{1, 0}
	blocks[2] = &tetrisPos{2, 0}
	blocks[3] = &tetrisPos{3, 0}
	return &tetris{&tetrisPos{0, 0}, blocks, []int{3}, []int{0}, []int{0, 1, 2, 3}}
}

func (d day17) createPlus() *tetris {
	blocks := make([]*tetrisPos, 5)
	blocks[0] = &tetrisPos{1, 0}
	blocks[1] = &tetrisPos{0, 1}
	blocks[2] = &tetrisPos{1, 1}
	blocks[3] = &tetrisPos{2, 1}
	blocks[4] = &tetrisPos{1, 2}
	return &tetris{&tetrisPos{0, 0}, blocks, []int{0, 3, 4}, []int{0, 1, 4}, []int{0, 1, 3}}
}

func (d day17) createL() *tetris {
	blocks := make([]*tetrisPos, 5)
	blocks[0] = &tetrisPos{0, 0}
	blocks[1] = &tetrisPos{1, 0}
	blocks[2] = &tetrisPos{2, 0}
	blocks[3] = &tetrisPos{2, 1}
	blocks[4] = &tetrisPos{2, 2}
	return &tetris{&tetrisPos{0, 0}, blocks, []int{2, 3, 4}, []int{0, 3, 4}, []int{0, 1, 2}}
}

func (d day17) createBar() *tetris {
	blocks := make([]*tetrisPos, 4)
	blocks[0] = &tetrisPos{0, 0}
	blocks[1] = &tetrisPos{0, 1}
	blocks[2] = &tetrisPos{0, 2}
	blocks[3] = &tetrisPos{0, 3}
	return &tetris{&tetrisPos{0, 0}, blocks, []int{0, 1, 2, 3}, []int{0, 1, 2, 3}, []int{0}}
}

func (d day17) createQuad() *tetris {
	blocks := make([]*tetrisPos, 4)
	blocks[0] = &tetrisPos{0, 0}
	blocks[1] = &tetrisPos{1, 0}
	blocks[2] = &tetrisPos{0, 1}
	blocks[3] = &tetrisPos{1, 1}
	return &tetris{&tetrisPos{0, 0}, blocks, []int{1, 3}, []int{0, 2}, []int{0, 1}}
}
