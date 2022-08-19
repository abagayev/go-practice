package game

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGame_neighbours(t *testing.T) {
	tests := []struct {
		name       string
		cells      [][]int
		neighbours [][]int
	}{
		{
			"4x4",
			[][]int{
				{0, 1, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			[][]int{
				{2, 1, 2, 0},
				{2, 1, 2, 0},
				{1, 1, 1, 0},
				{0, 0, 0, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Field{tt.cells}
			g := Game{f}

			for i := 0; i < len(tt.cells); i++ {
				for j := 0; j < len(tt.cells[i]); j++ {
					if g.neighbours(i, j) != tt.neighbours[i][j] {
						t.Fail()
					}
				}
			}
		})
	}
}

func TestGame_Live(t *testing.T) {
	tests := []struct {
		name   string
		frames [][][]int
	}{
		{
			"Block",
			[][][]int{
				{
					{0, 0, 0, 0},
					{0, 1, 1, 0},
					{0, 1, 1, 0},
					{0, 0, 0, 0},
				},
				{
					{0, 0, 0, 0},
					{0, 1, 1, 0},
					{0, 1, 1, 0},
					{0, 0, 0, 0},
				},
				{
					{0, 0, 0, 0},
					{0, 1, 1, 0},
					{0, 1, 1, 0},
					{0, 0, 0, 0},
				},
			},
		},
		{
			"Blinker",
			[][][]int{
				{
					{0, 0, 0},
					{1, 1, 1},
					{0, 0, 0},
				},
				{
					{0, 1, 0},
					{0, 1, 0},
					{0, 1, 0},
				},
				{
					{0, 0, 0},
					{1, 1, 1},
					{0, 0, 0},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Field{tt.frames[0]}
			g := Game{f}

			for i := 1; i < len(tt.frames); i++ {
				g.Live()

				if !reflect.DeepEqual(g.Field.cells, tt.frames[i]) {
					fmt.Println(i, g.Field.cells)
					t.Fail()
				}
			}
		})
	}
}
