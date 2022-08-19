package main

import (
	"github.com/abagayev/go-practice/basics/games/life/console"
	"github.com/abagayev/go-practice/basics/games/life/field"
	"github.com/abagayev/go-practice/basics/games/life/game"
)

func main() {
	f := &field.Field{}
	f.Init(20, 20)
	f.Shake(3)

	g := &game.Game{f}
	c := &console.Console{g}

	c.Play(g)
}
