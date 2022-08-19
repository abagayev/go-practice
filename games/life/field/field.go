package field

import (
	"math/rand"
	"time"
)

type Field struct {
	cells [][]int
}

func (f *Field) Init(width, height uint) {
	f.cells = make([][]int, height)
	for i := range f.cells {
		f.cells[i] = make([]int, width)
	}
}

func (f *Field) Shake(density uint) {
	// initialize the generator to a deterministic state
	rand.Seed(time.Now().UnixNano())

	for i, row := range f.cells {
		for j := range row {
			r := rand.Intn(10)
			if r > (10 - int(density)) {
				f.cells[i][j] = 1
				continue
			}

			f.cells[i][j] = 0
		}
	}
}

func (f *Field) Get(i, j int) int {
	if !f.exists(i, j) {
		return 0
	}

	return f.cells[i][j]
}

func (f *Field) Set(i, j, v int) {
	if !f.exists(i, j) {
		return
	}

	f.cells[i][j] = v
}

func (f *Field) Width() int {
	return len(f.cells[0])
}

func (f *Field) Height() int {
	return len(f.cells)
}

func (f *Field) Output() [][]int {
	return f.cells
}

func (f *Field) exists(i, j int) bool {
	if i < 0 || j < 0 {
		return false
	}

	if i >= len(f.cells[0]) || j >= len(f.cells) {
		return false
	}

	return true
}
