package solution

import (
	"fmt"
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
	profile := d.createProfile()
	profiles := make(map[string]caveProfileEntry)
	extrapolatedHeight := int64(0)
	blocks := make(map[tetrisPos]bool)
	for i := int64(0); i < amountRocks; i++ {
		form := forms[formIndex].new(maxHeight)
		formIndex = (formIndex + 1) % len(forms)
		for {
			stream := streams[streamIndex]
			streamIndex = (streamIndex + 1) % len(streams)
			form.push(stream, &blocks)
			couldDrop := form.drop(&blocks)
			if !couldDrop {
				newMaxHeight := maxHeight
				for _, block := range form.getBlocks() {
					blocks[*block] = true
					profile.addBlock(block.x, block.y, maxHeight)
					if block.y > newMaxHeight {
						newMaxHeight = block.y
					}
				}
				if newMaxHeight > maxHeight {
					profile.registerHeightIncrease(newMaxHeight - maxHeight)
					maxHeight = newMaxHeight
				}
				break
			}
		}
		if extrapolatedHeight == 0 {
			profileAsString := fmt.Sprintf("%d %d %s", (len(forms)+formIndex-1)%len(forms), (len(streams)+streamIndex-1)%len(streams), profile.String())
			entry, hasProfile := profiles[profileAsString]
			if hasProfile {
				repeatingRoundOffset := i - entry.round
				repeatingHeightOffset := maxHeight - entry.maxHeight
				extrapolatedRepeats := (amountRocks - i) / repeatingRoundOffset
				extrapolatedHeight = extrapolatedRepeats * repeatingHeightOffset
				i = i + extrapolatedRepeats*repeatingRoundOffset
			}
			profiles[profileAsString] = caveProfileEntry{i, maxHeight}
		}
	}
	return extrapolatedHeight + maxHeight + 1
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
		if pushedBlock.x < 0 || pushedBlock.x >= 7 || (*blocks)[pushedBlock] {
			canMove = false
			break
		}
	}
	if canMove {
		t.pos = &tetrisPos{t.pos.x + direction, t.pos.y}
	}
}

func (t *tetris) drop(blocks *map[tetrisPos]bool) bool {
	canMove := true
	for _, blockIndex := range t.moveDownBlocks {
		block := t.blocks[blockIndex]
		pushedBlock := tetrisPos{t.pos.x + block.x, t.pos.y + block.y - 1}
		if pushedBlock.y < 0 || (*blocks)[pushedBlock] {
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

type caveProfile struct {
	lastBlocksPerColumn [][]int64
}

func (d day17) createProfile() *caveProfile {
	lastBlocks := make([][]int64, 7)
	for i := 0; i < 7; i++ {
		lastBlocks[i] = make([]int64, 0)
	}
	return &caveProfile{lastBlocks}
}

func (c *caveProfile) addBlock(x, y, maxHeight int64) {
	c.lastBlocksPerColumn[x] = append(c.lastBlocksPerColumn[x], maxHeight-y)
	if len(c.lastBlocksPerColumn[x]) > 5 {
		c.lastBlocksPerColumn[x] = c.lastBlocksPerColumn[x][len(c.lastBlocksPerColumn[x])-5 : len(c.lastBlocksPerColumn[x])-1]
	}
}

func (c *caveProfile) registerHeightIncrease(increase int64) {
	for x := range c.lastBlocksPerColumn {
		for y := range c.lastBlocksPerColumn[x] {
			c.lastBlocksPerColumn[x][y] = c.lastBlocksPerColumn[x][y] + increase
		}
	}
}

func (c *caveProfile) String() string {
	return fmt.Sprintf("%v", c.lastBlocksPerColumn)
}

type caveProfileEntry struct {
	round     int64
	maxHeight int64
}
