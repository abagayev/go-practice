package main

import (
	"fmt"
	"math/rand"
	"time"
)

func field(width, height uint) [][]int {
	screen := make([][]int, height)
	for i := range screen {
		screen[i] = make([]int, width)
	}

	return screen
}

func shake(screen [][]int) [][]int {
	// initialize the generator to a deterministic state
	rand.Seed(time.Now().UnixNano())

	for i, row := range screen {
		for j := range row {
			r := rand.Intn(10)
			if r > (10 - 3) {
				screen[i][j] = 1
				continue
			}

			screen[i][j] = 0
		}
	}

	return screen
}

// check if cell is alive
func alive(screen [][]int, i, j int) bool {
	if i < 0 || j < 0 {
		return false
	}

	if i >= len(screen[0]) || j >= len(screen) {
		return false
	}

	return screen[i][j] > 0
}

// count cell neighbours
func neighbours(screen [][]int, i, j int) (count int) {
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
		if alive(screen, i+id, j+jd) {
			count += 1
		}
	}

	return count
}

func live(screen [][]int) [][]int {
	for i, row := range screen {
		for j, _ := range row {
			al := alive(screen, i, j)
			ns := neighbours(screen, i, j)

			// Any live cell with fewer than two live neighbours dies, as if by underpopulation
			if al && ns < 2 {
				screen[i][j] = 0
			}

			// Any live cell with two or three live neighbours lives on to the next generation
			if al && ns >= 2 && ns <= 3 {
				screen[i][j] = 1
			}

			// Any live cell with more than three live neighbours dies, as if by overpopulation
			if al && ns > 3 {
				screen[i][j] = 0
			}

			// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction
			if !al && ns == 3 {
				screen[i][j] = 1
			}
		}
	}

	return screen
}

func draw(screen [][]int) {
	for _, row := range screen {
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
	screen := shake(field(20, 20))

	for range [50]int{} {
		// clear the screen
		fmt.Print("\033[H\033[2J")

		draw(screen)
		screen = live(screen)

		time.Sleep(1 * time.Second)
	}
}
