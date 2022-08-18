package main

import (
	"fmt"
	"time"

	"github.com/abagayev/go-practice/basics/games/life/field"
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
	f.Start(20, 20)
	f.Shake(3)

	for range [50]int{} {
		// clear the screen
		fmt.Print("\033[H\033[2J")

		draw(f)
		f.Live()

		time.Sleep(1 * time.Second)
	}
}
