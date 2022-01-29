package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
)

type Player struct {
	Obj  *resolv.Object
	Icon *ebiten.Image
}

var (
	State  string = "menu"
	BG     *ebiten.Image
	player Player
	Space  *resolv.Space
)

type Game struct{}

func init() {
	BG, _, _ = ebitenutil.NewImageFromFile("assets/bg.png")

	Space = resolv.NewSpace(1280, 720, 16, 16)

	player.Icon, _, _ = ebitenutil.NewImageFromFile("assets/character1.png")
	player.Obj = resolv.NewObject(622/4, 300/4, 36, 56, "player") // divided by 4 as the image is scaled up by 4

	Space.Add(player.Obj)
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
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(player.Obj.X, player.Obj.Y)
		op.GeoM.Scale(4, 4)
		screen.DrawImage(player.Icon, op)
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
