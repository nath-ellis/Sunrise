package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	LeftClick, _, _   = ebitenutil.NewImageFromFile("assets/left-click.png")
	Sunrise, _, _     = ebitenutil.NewImageFromFile("assets/sunrise.png")
	GameOverImg, _, _ = ebitenutil.NewImageFromFile("assets/game-over.png")
)

func DrawMenu(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(140, 50)
	screen.DrawImage(Sunrise, op)
	op.GeoM.Reset()
	op.GeoM.Scale(3, 3)
	op.GeoM.Translate(300, 250)
	screen.DrawImage(LeftClick, op)
}

func DrawGameOver(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(215, 50)
	screen.DrawImage(GameOverImg, op)
	op.GeoM.Reset()
	op.GeoM.Scale(3, 3)
	op.GeoM.Translate(300, 250)
	screen.DrawImage(LeftClick, op)
}
