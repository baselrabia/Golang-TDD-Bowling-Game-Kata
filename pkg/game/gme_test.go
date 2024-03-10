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

 }
