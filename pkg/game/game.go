package game

type Game struct {
    rolled [22]int
	currentBall int 
}

func NewGame() Game {
	return Game{}
}


func (g *Game) Roll(pins int) {
	g.rolled[g.currentBall] = pins
	g.currentBall++
}

func (g *Game) Score() int {
	score := 0

	for i := 0; i < len(g.rolled)-2; i++ {

		if (g.rolled[i] + g.rolled[i+1] == 10) { //spare 
			score += g.rolled[i+2]
		}

		score += g.rolled[i]
	}

	return score
}