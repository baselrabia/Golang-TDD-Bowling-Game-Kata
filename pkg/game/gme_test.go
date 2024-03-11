package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScore(t *testing.T) {
 	t.Run("roll zero score is zero", func(t *testing.T) {
		game := NewGame()
		game.Roll(0)

		assert.Equal(t, 0, game.Score())
	})


	t.Run("open frame adds pins", func(t *testing.T) {
		game := NewGame()
		game.Roll(4)
		game.Roll(2)

		assert.Equal(t, 6, game.Score())
	})


	t.Run("spare adds next ball", func(t *testing.T) {
		game := NewGame()
		game.Roll(4)
		game.Roll(6)
		game.Roll(3)
		game.Roll(0)

		assert.Equal(t, 16, game.Score())
	})


	t.Run("a ten in two frames is not a spare", func(t *testing.T) {
		game := NewGame()
		game.Roll(3)
		game.Roll(6)
		game.Roll(4)
		game.Roll(2)

		assert.Equal(t, 15, game.Score())
	})


	t.Run("a strike add next two balls", func(t *testing.T) {
		game := NewGame()
 		game.Roll(10)
		game.Roll(3)
		game.Roll(2)

		assert.Equal(t, 20, game.Score())
	})


	t.Run("the perfect game score is 300", func(t *testing.T) {
		game := NewGame()

		for i := 0; i < 12; i++ {
			game.Roll(10)
	
		}

		assert.Equal(t, 300, game.Score())
	})
 }
