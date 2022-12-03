package player

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// For drawing the player and its animations
func Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(Player.Obj.X, Player.Obj.Y)

	if Player.Moving {
		// Changes the frame of the animation
		if Player.MSCool <= 0 {
			Player.MoveStage += 1
			Player.MSCool = 1
		} else {
			if Player.MSCool > 0 {
				Player.MSCool -= 1
			}
		}

		// Resets the frame back to the first
		if Player.MoveStage >= 12 {
			Player.MoveStage = 0
		}

		if Player.IsLeft {
			screen.DrawImage(Player.L[Player.MoveStage], op)
		} else {
			screen.DrawImage(Player.R[Player.MoveStage], op)
		}
	} else {
		if Player.IsLeft {
			screen.DrawImage(Player.IdleL, op)
		} else {
			screen.DrawImage(Player.IdleR, op)
		}
	}
}

// For drawing the weapon
func DrawWeapon(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	mouseX, mouseY := ebiten.CursorPosition()

	dirX := float64(mouseX) - Player.Obj.X
	dirY := float64(mouseY) - Player.Obj.Y

	// The square root of the sum of dirX (Pythagoras a) and dirY (Pythagoras b) to find length of hypotenuse
	length := math.Hypot(dirX, dirY)

	if length == 0.0 {
		dirX = 0
		dirY = -1
	} else {
		dirX /= length
		dirY /= length
	}

	// math.Atan2 returns the inverse tan of dirY/dirX
	angle := math.Atan2(dirY, dirX)
	// Rotates the weapon
	op.GeoM.Translate(20, 0)
	op.GeoM.Rotate(angle)
	// Moves it near to the player
	op.GeoM.Translate(Player.Obj.X+10, Player.Obj.Y+10)

	if math.Signbit(dirX) { // Mouse is left of the player
		Player.IsLeft = true
		screen.DrawImage(Gun2, op)
	} else { // Mouse is right of the player
		Player.IsLeft = false
		screen.DrawImage(Gun1, op)
	}
}

// For drawing bullets
func DrawBullets(screen *ebiten.Image) {
	for _, b := range Bullets {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(b.Obj.X, b.Obj.Y)
		screen.DrawImage(BulletImg, op)
	}
}
