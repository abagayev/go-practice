package field

import (
	"math/rand"
	"time"
)

type Field struct {
	cells [][]int
}

func (f *Field) Start(width, height uint) {
	f.cells = make([][]int, height)
	for i := range f.cells {
		f.cells[i] = make([]int, width)
	}
}

func (f *Field) Shake(density int) {
	// initialize the generator to a deterministic state
	rand.Seed(time.Now().UnixNano())

	for i, row := range f.cells {
		for j := range row {
			r := rand.Intn(10)
			if r > (10 - density) {
				f.cells[i][j] = 1
				continue
			}

			f.cells[i][j] = 0
		}
	}
}

// check if cell is alive
func (f *Field) alive(i, j int) bool {
	if i < 0 || j < 0 {
		return false
	}

	if i >= len(f.cells[0]) || j >= len(f.cells) {
		return false
	}

	return f.cells[i][j] > 0
}

// count cell neighbours
func (f *Field) neighbours(i, j int) (count int) {
	matrix := [][]int{
		{-1, -1},
		{-1, +0},
		{-1, +1},
		{+1, -1},
		{+1, +0},
		{+1, +1},
		{+0, -1},
		{+0, +1},
	}

	for _, values := range matrix {
		id, jd := values[0], values[1]
		if f.alive(i+id, j+jd) {
			count += 1
		}
	}

	return count
}

func (f *Field) Live() {
	for i, row := range f.cells {
		for j, _ := range row {
			al := f.alive(i, j)
			ns := f.neighbours(i, j)

			// Any Live cell with fewer than two Live neighbours dies, as if by underpopulation
			if al && ns < 2 {
				f.cells[i][j] = 0
			}

			// Any Live cell with two or three Live neighbours lives on to the next generation
			if al && ns >= 2 && ns <= 3 {
				f.cells[i][j] = 1
			}

			// Any Live cell with more than three Live neighbours dies, as if by overpopulation
			if al && ns > 3 {
				f.cells[i][j] = 0
			}

			// Any dead cell with exactly three Live neighbours becomes a Live cell, as if by reproduction
			if !al && ns == 3 {
				f.cells[i][j] = 1
			}
		}
	}
}

func (f *Field) Output() [][]int {
	return f.cells
}
