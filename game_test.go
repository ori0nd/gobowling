package gobowling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoThrowsNoMarks(t *testing.T) {
	game := CreateGame()
	game.AddThrow(5)
	game.AddThrow(4)
	assert.Equal(t, 9, game.GetScore())
	assert.Equal(t, 2, game.CurrentFrame)
}

func TestFourThrowsNoMarks(t *testing.T) {
	game := CreateGame()
	game.AddThrow(5)
	game.AddThrow(4)
	game.AddThrow(7)
	game.AddThrow(2)
	assert.Equal(t, 18, game.GetScore())
	assert.Equal(t, 9, game.scoreForFrame(1))
	assert.Equal(t, 18, game.scoreForFrame(2))
	assert.Equal(t, 3, game.CurrentFrame)
}

func TestSimpleSpare(t *testing.T) {
	game := CreateGame()
	game.AddThrow(3)
	game.AddThrow(7)
	game.AddThrow(3)
	assert.Equal(t, 13, game.scoreForFrame(1))
	assert.Equal(t, 2, game.CurrentFrame)
}

func TestSimpleFrameAfterSpare(t *testing.T) {
	game := CreateGame()
	game.AddThrow(3)
	game.AddThrow(7)
	game.AddThrow(3)
	game.AddThrow(2)
	assert.Equal(t, 13, game.scoreForFrame(1))
	assert.Equal(t, 18, game.scoreForFrame(2))
	assert.Equal(t, 18, game.GetScore())
	assert.Equal(t, 3, game.CurrentFrame)
}

func TestSimpleStrike(t *testing.T) {
	game := CreateGame()
	game.AddThrow(10)
	game.AddThrow(3)
	game.AddThrow(6)
	assert.Equal(t, 19, game.scoreForFrame(1))
	assert.Equal(t, 28, game.GetScore())
	assert.Equal(t, 3, game.CurrentFrame)
}

func TestPerfectGame(t *testing.T) {
	game := CreateGame()
	for i := 0; i < 12; i++ {
		game.AddThrow(10)
	}
	assert.Equal(t, 300, game.GetScore())
	assert.Equal(t, 10, game.CurrentFrame)
}

func TestEndOfArray(t *testing.T) {
	game := CreateGame()
	for i := 0; i < 9; i++ {
		game.AddThrow(0)
		game.AddThrow(0)
	}
	game.AddThrow(2)
	game.AddThrow(8)
	game.AddThrow(10)
	assert.Equal(t, 20, game.GetScore())
}

func TestSampleGame(t *testing.T) {
	game := CreateGame()
	game.AddThrow(1)
	game.AddThrow(4)
	game.AddThrow(4)
	game.AddThrow(5)
	game.AddThrow(6)
	game.AddThrow(4)
	game.AddThrow(5)
	game.AddThrow(5)
	game.AddThrow(10)
	game.AddThrow(0)
	game.AddThrow(1)
	game.AddThrow(7)
	game.AddThrow(3)
	game.AddThrow(6)
	game.AddThrow(4)
	game.AddThrow(10)
	game.AddThrow(2)
	game.AddThrow(8)
	game.AddThrow(6)
	assert.Equal(t, 133, game.GetScore())
}

func TestHeartBreak(t *testing.T) {
	game := CreateGame()
	for i := 0; i < 11; i++ {
		game.AddThrow(10)
	}
	game.AddThrow(9)
	assert.Equal(t, 299, game.GetScore())
}

func TestTenthFrameSpare(t *testing.T) {
	game := CreateGame()
	for i := 0; i < 9; i++ {
		game.AddThrow(10)
	}
	game.AddThrow(9)
	game.AddThrow(1)
	game.AddThrow(1)
	assert.Equal(t, 270, game.GetScore())
}
