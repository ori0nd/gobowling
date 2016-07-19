package gobowling

import "math"

// Game structure represents a single bowling game
type Game struct {
	CurrentFrame      int
	firstThrowInFrame bool
	scorer            *scorer
}

// CreateGame creates a fresh game
func CreateGame() (game *Game) {
	game = new(Game)
	game.scorer = new(scorer)
	game.firstThrowInFrame = true
	game.CurrentFrame = 1
	return
}

// GetScore returns current game score
func (g *Game) GetScore() int {
	return g.scoreForFrame(g.CurrentFrame)
}

// AddThrow simulates in-game throw.
// pins is # of pins taken down by the throw.
func (g *Game) AddThrow(pins int) {
	g.scorer.addThrow(pins)
	g.adjustCurrentFrame(pins)
}

func (g *Game) scoreForFrame(theFrame int) int {
	return g.scorer.scoreForFrame(theFrame)
}

func (g *Game) adjustCurrentFrame(pins int) {
	if g.firstThrowInFrame == true {
		if g.adjustFrameForStrike(pins) == false {
			g.firstThrowInFrame = false
		}
	} else {
		g.firstThrowInFrame = true
		g.advanceFrame()
	}
}

func (g *Game) adjustFrameForStrike(pins int) bool {
	if pins == 10 {
		g.advanceFrame()
		return true
	}
	return false
}

func (g *Game) advanceFrame() {
	g.CurrentFrame++
	g.CurrentFrame = int(math.Min(10, float64(g.CurrentFrame)))
}
