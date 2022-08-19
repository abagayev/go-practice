package main

import (
	"flag"

	"github.com/abagayev/go-practice/basics/games/life/console"
	"github.com/abagayev/go-practice/basics/games/life/game"
)

func main() {
	width := flag.Uint("width", 20, "screen width")
	height := flag.Uint("height", 20, "screen height")
	density := flag.Uint("density", 3, "population density [1..10]")

	flag.Parse()

	g := game.NewGame(*width, *height, *density)
	c := &console.Console{g}

	c.Play(g)
}
