package game

type Game struct {
	score int
}

func NewGame() Game {
	return Game{}
}

func (g *Game) Score() int {
	return g.score
}

func (g *Game) Roll(rolls int) {
	g.score += rolls
}
