package console

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestConsole_draw(t *testing.T) {
	tests := []struct {
		name   string
		cells  [][]int
		output []string
	}{
		{
			"4x4",
			[][]int{
				{0, 1, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			[]string{
				". ∎ . . ",
				". ∎ . . ",
				". . . . ",
				". . . . ",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := GameMock{}
			g.cells = tt.cells

			c := Console{&g}

			// Capture stdout of a function into a string.
			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			c.draw(&g)

			w.Close()
			out, _ := io.ReadAll(r)
			os.Stdout = rescueStdout

			// Remove last element from captured output (it's always a break line).
			if string(out[:len(out)-1]) != strings.Join(tt.output, "\n") {
				t.Fail()
			}
		})
	}
}

type GameMock struct {
	cells [][]int
}

func (g *GameMock) Live() {}

func (g *GameMock) Output() [][]int {
	return g.cells
}
