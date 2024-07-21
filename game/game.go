package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *Player
}

func NewGame() *Game {
	player := NewPlayer()
	return &Game{
		player: player,
	}
}

// Update game logic
func (g *Game) Update() error {
	return nil
}

// Draw object on screen
func (g *Game) Draw(screen *ebiten.Image)  {
	g.player.Draw(screen)
}

// Return screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}