package console

import (
	"fmt"
	"time"
)

type Console struct {
	Game Game
}

type Game interface {
	Live()
	Output() [][]int
}

// Draw a screen of cells.
func (c *Console) draw(g Game) {
	for _, row := range g.Output() {
		for _, v := range row {

			if v == 0 {
				fmt.Print(". ")
				continue
			}

			fmt.Print("∎ ")
		}

		fmt.Println()
	}
}

// Clear the console screen.
func (c *Console) clear() {
	fmt.Print("\033[H\033[2J")
}

// Play the game step by step.
func (c *Console) Play(g Game) {
	for i := 0; i < 50; i++ {
		c.clear()
		c.draw(g)
		g.Live()

		time.Sleep(1 * time.Second)
	}
}
