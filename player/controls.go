package player

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nath-ellis/Sunrise/objects"
	"github.com/solarlune/resolv"
)

// For moving the Player (moves objects instead of Player)
// Only moves values stored in player package to prevent an import cycle with the enemies package
// Enemies and Particles are moved in the enemies Update function
func Controls() {
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		if c := Player.Obj.Check(0, -Player.Speed, "object"); c == nil {
			for _, o := range objects.Objects {
				o.Obj.Y += Player.Speed
				o.Obj.Update()
			}

			for _, b := range Bullets {
				b.Obj.Y += Player.Speed
				b.Obj.Update()
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		if c := Player.Obj.Check(0, Player.Speed, "object"); c == nil {
			for _, o := range objects.Objects {
				o.Obj.Y -= Player.Speed
				o.Obj.Update()
			}

			for _, b := range Bullets {
				b.Obj.Y -= Player.Speed
				b.Obj.Update()
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if c := Player.Obj.Check(-Player.Speed, 0, "object"); c == nil {
			for _, o := range objects.Objects {
				o.Obj.X += Player.Speed
				o.Obj.Update()
			}

			for _, b := range Bullets {
				b.Obj.X += Player.Speed
				b.Obj.Update()
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		if c := Player.Obj.Check(Player.Speed, 0, "object"); c == nil {
			for _, o := range objects.Objects {
				o.Obj.X -= Player.Speed
				o.Obj.Update()
			}

			for _, b := range Bullets {
				b.Obj.X -= Player.Speed
				b.Obj.Update()
			}
		}
	}

	// For drawing the Player
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) ||
		ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) ||
		ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		Player.Moving = true
	} else {
		Player.Moving = false
	}

	// Checks if Player is being attacked by an enemy
	if Player.ImmunityTicks <= 0 {
		if c := Player.Obj.Check(0, 0, "zombie", "mini-zombie"); c != nil {
			if c.HasTags("zombie") {
				Player.Health -= 1
			} else if c.HasTags("mini-zombie") {
				Player.Health -= 2
			}
			Player.ImmunityTicks = 10
		}
	} else {
		Player.ImmunityTicks -= 1
	}

}

// For shooting
func Shoot(Space *resolv.Space) {
	mouseX, mouseY := ebiten.CursorPosition()

	if Player.ShootCool > 0 {
		Player.ShootCool -= 1
	}

	if Player.ShootCool <= 0 {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			dirX := float64(mouseX) - Player.Obj.X
			dirY := float64(mouseY) - Player.Obj.Y

			// The directions
			length := math.Hypot(dirX, dirY)

			if length == 0.0 {
				dirX = 0
				dirY = -1
			} else {
				dirX = dirX / length
				dirY = dirY / length
			}

			// Adds a bullet
			if math.Signbit(dirX) {
				Bullets = append(Bullets, Bullet{resolv.NewObject(Player.Obj.X+10, Player.Obj.Y+6, 5, 5, "bullet"), dirX, dirY})
			} else {
				Bullets = append(Bullets, Bullet{resolv.NewObject(Player.Obj.X+12, Player.Obj.Y+12, 5, 5, "bullet"), dirX, dirY})
			}

			// Adds the hitboxes
			for _, b := range Bullets {
				Space.Add(b.Obj)
			}

			Player.ShootCool += 25
		}
	}
}
