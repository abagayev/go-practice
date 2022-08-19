package game

type Field interface {
	Get(i, j int) int
	Set(i, j, v int)
	Width() int
	Height() int
}

type Game struct {
	Field Field
}

// count cell neighbours
func (g *Game) neighbours(i, j int) (count int) {
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
		if g.Field.Get(i+id, j+jd) > 0 {
			count += 1
		}
	}

	return count
}

func (g *Game) Live() {
	for i := 0; i < g.Field.Height(); i++ {
		for j := 0; j < g.Field.Width(); j++ {
			al := g.Field.Get(i, j) > 0
			ns := g.neighbours(i, j)

			// Any Live cell with fewer than two Live neighbours dies, as if by underpopulation
			if al && ns < 2 {
				g.Field.Set(i, j, 0)
			}

			// Any Live cell with two or three Live neighbours lives on to the next generation
			if al && ns >= 2 && ns <= 3 {
				g.Field.Set(i, j, 1)
			}

			// Any Live cell with more than three Live neighbours dies, as if by overpopulation
			if al && ns > 3 {
				g.Field.Set(i, j, 0)
			}

			// Any dead cell with exactly three Live neighbours becomes a Live cell, as if by reproduction
			if !al && ns == 3 {
				g.Field.Set(i, j, 1)
			}
		}
	}
}
