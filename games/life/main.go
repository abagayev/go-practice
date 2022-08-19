package main

import (
	"fmt"
	"time"

	"github.com/abagayev/go-practice/basics/games/life/field"
	"github.com/abagayev/go-practice/basics/games/life/game"
)

func draw(f *field.Field) {
	for _, row := range f.Output() {
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

func main() {
	f := &field.Field{}
	f.Init(20, 20)
	f.Shake(3)

	g := &game.Game{f}

	for range [50]int{} {
		// clear the screen
		fmt.Print("\033[H\033[2J")

		draw(f)
		g.Live()

		time.Sleep(1 * time.Second)
	}
}
