package main

import (
	"github.com/abagayev/go-practice/basics/games/life/console"
	"github.com/abagayev/go-practice/basics/games/life/game"
)

func main() {
	g := game.NewGame(20, 20, 3)
	c := &console.Console{g}

	c.Play(g)
}
