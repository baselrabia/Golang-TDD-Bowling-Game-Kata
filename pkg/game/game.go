package game

type Game struct {
	rolled      [22]int
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
	thisBall := 0

	for i := 0; i < 10; i++ {

		if g.rolled[i] == 10 { // strike 
			score += g.rolled[thisBall] + g.rolled[thisBall+1] + g.rolled[thisBall+2]
			thisBall++
		} else {

			if g.rolled[thisBall]+g.rolled[thisBall+1] == 10 { //spare
				score += g.rolled[thisBall+2]
			}

			score += g.rolled[thisBall] + g.rolled[thisBall+1]
			thisBall+=2
		}

	}

	return score
}
