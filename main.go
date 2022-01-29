package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	State string = "menu"
	BG    *ebiten.Image
)

type Game struct{}

func init() {
	BG, _, _ = ebitenutil.NewImageFromFile("assets/bg.png")
}

func (g *Game) Update() error {
	switch State {
	case "menu":
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			State = "game"
		}
	case "game":
	case "gameOver":
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(BG, nil)

	switch State {
	case "menu":
	case "game":
	case "gameOver":
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Sunrise")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Panic("Failed to run game: ", err)
	}
}
