package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nath-ellis/Sunrise/player"
)

var Heart, _, _ = ebitenutil.NewImageFromFile("assets/heart.png")

// For drawing hearts
func DrawHealth(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(5, 5)
	op.GeoM.Scale(0.8, 0.8)

	for i := 0; i < player.Player.Health; i++ {
		screen.DrawImage(Heart, op)
		op.GeoM.Translate(33, 0)
	}
}
