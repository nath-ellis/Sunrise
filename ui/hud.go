package ui

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nath-ellis/Sunrise/enemies"
	"github.com/nath-ellis/Sunrise/player"
)

var (
	Heart, _, _ = ebitenutil.NewImageFromFile("assets/heart.png")
	Zero, _, _  = ebitenutil.NewImageFromFile("assets/numbers/0.png")
	One, _, _   = ebitenutil.NewImageFromFile("assets/numbers/1.png")
	Two, _, _   = ebitenutil.NewImageFromFile("assets/numbers/2.png")
	Three, _, _ = ebitenutil.NewImageFromFile("assets/numbers/3.png")
	Four, _, _  = ebitenutil.NewImageFromFile("assets/numbers/4.png")
	Five, _, _  = ebitenutil.NewImageFromFile("assets/numbers/5.png")
	Six, _, _   = ebitenutil.NewImageFromFile("assets/numbers/6.png")
	Seven, _, _ = ebitenutil.NewImageFromFile("assets/numbers/7.png")
	Eight, _, _ = ebitenutil.NewImageFromFile("assets/numbers/8.png")
	Nine, _, _  = ebitenutil.NewImageFromFile("assets/numbers/9.png")
)

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

func DrawWaveNumber(screen *ebiten.Image) {
	WaveNumber := strconv.Itoa(enemies.Wave)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(5, 42)

	for _, w := range WaveNumber {
		switch string(w) {
		case "0":
			screen.DrawImage(Zero, op)
			op.GeoM.Translate(28, 0)
		case "1":
			screen.DrawImage(One, op)
			op.GeoM.Translate(16, 0)
		case "2":
			screen.DrawImage(Two, op)
			op.GeoM.Translate(28, 0)
		case "3":
			screen.DrawImage(Three, op)
			op.GeoM.Translate(28, 0)
		case "4":
			screen.DrawImage(Four, op)
			op.GeoM.Translate(28, 0)
		case "5":
			screen.DrawImage(Five, op)
			op.GeoM.Translate(28, 0)
		case "6":
			screen.DrawImage(Six, op)
			op.GeoM.Translate(28, 0)
		case "7":
			screen.DrawImage(Seven, op)
			op.GeoM.Translate(28, 0)
		case "8":
			screen.DrawImage(Eight, op)
			op.GeoM.Translate(28, 0)
		case "9":
			screen.DrawImage(Nine, op)
			op.GeoM.Translate(28, 0)
		}
	}
}
