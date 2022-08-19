package game

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jinzhu/copier"
)

func TestField_Init(t *testing.T) {
	tests := []struct {
		name   string
		width  uint
		height uint
		output [][]int
	}{
		{
			name:   "EmptyField",
			width:  0,
			height: 0,
			output: [][]int{},
		},
		{
			name:   "2x2",
			width:  2,
			height: 2,
			output: [][]int{{0, 0}, {0, 0}},
		},
		{
			name:   "3x2",
			width:  3,
			height: 2,
			output: [][]int{{0, 0, 0}, {0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Field{}
			f.Init(tt.width, tt.height)

			if !reflect.DeepEqual(tt.output, f.cells) {
				t.Fail()
			}
		})
	}
}

func TestField_Shake(t *testing.T) {
	for d := 1; d < 10; d++ {
		t.Run(fmt.Sprintf("Density%d", d), func(t *testing.T) {
			f := Field{}
			f.Init(3, 3)
			f.Shake(uint(d))

			for i := 0; i < 10; i++ {
				// Create a copy of previous state and shake field again.
				prev := [][]int{}
				copier.Copy(&f.cells, &prev)
				f.Shake(uint(d))

				if reflect.DeepEqual(prev, f.cells) {
					t.Fail()
				}
			}
		})
	}
}
