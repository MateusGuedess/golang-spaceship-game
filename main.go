package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {

}

func main() {
	g := &Game{}

	err := ebiten.RunGame(g)

	if err != nil {
		panic(err)
	}
}

// Update game logic
func (g *Game) Update() error {
	return nil
}

// Draw object on screen
func (g *Game) Draw(screen *ebiten.Image)  {
	
}

// Return screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}