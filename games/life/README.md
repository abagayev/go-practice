# Conway's Game of Life

A classical zero-player game written in Golang. The main purpose of this project is to lean about main conceptions and
principles: condition, cycles, data structures, packages, documentation and tests.

![](./docs/life.gif)

To start the game, run a script without or with command-line flags:

```shell
# start a basic game
go run main.go

# change screen size
go run main.go -width=30 -height=30

# change starting population density
go run main.go -density=9
```

### For curious minds

Still curios what is going on here? Please refer to Wikipedia:   
https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
