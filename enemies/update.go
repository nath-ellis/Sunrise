package enemies

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nath-ellis/Sunrise/player"
	"github.com/solarlune/resolv"
)

// Updates and moves the enemies
func Update(Space *resolv.Space) {
	for _, e := range Enemies {
		// Left Collisions
		if c := e.Obj.Check(e.Speed, 0, "object"); c != nil {
			e.Obj.Y -= e.Speed
			e.Obj.Update()
			continue
		}
		// Right collisisons
		if c := e.Obj.Check(-e.Speed, 0, "object"); c != nil {
			e.Obj.Y -= e.Speed
			e.Obj.Update()
			continue
		}
		// Above Collisions
		if c := e.Obj.Check(0, e.Speed, "object"); c != nil {
			e.Obj.X -= e.Speed
			e.Obj.Update()
			continue
		}
		// Below Collison
		if c := e.Obj.Check(0, -e.Speed, "object"); c != nil {
			e.Obj.X -= e.Speed
			e.Obj.Update()
			continue
		}

		// Left of player
		if e.Obj.X < player.Player.Obj.X {
			e.Obj.X += e.Speed
		}

		// Right of player
		if e.Obj.X >= player.Player.Obj.X {
			e.Obj.X -= e.Speed
		}

		// Above player
		if e.Obj.Y < player.Player.Obj.Y {
			e.Obj.Y += e.Speed
		}

		// Below player
		if e.Obj.Y >= player.Player.Obj.Y {
			e.Obj.Y -= e.Speed
		}

		// If health is 0
		if e.Health <= 0 {
			tmp := []Enemy{}
			for _, e := range Enemies {
				if e.Health <= 0 {
					Space.Remove(e.Obj)
					continue
				}
				tmp = append(tmp, e)
			}

			Particles = append(Particles, Particle{resolv.NewObject(e.Obj.X+10, e.Obj.Y+6, 14, 12, "particle"), 60})

			Enemies = []Enemy{}
			Enemies = tmp
			continue
		}

		e.Obj.Update()
	}

	// If the player moves
	// moves the particles and enemies to prevent an import cycle with the package player
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		if c := player.Player.Obj.Check(0, -player.Player.Speed, "object"); c == nil {
			for _, e := range Enemies {
				e.Obj.Y += player.Player.Speed
				e.Obj.Update()
			}

			for _, p := range Particles {
				p.Obj.Y += player.Player.Speed
				p.Obj.Update()
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		if c := player.Player.Obj.Check(0, player.Player.Speed, "object"); c == nil {
			for _, e := range Enemies {
				e.Obj.Y -= player.Player.Speed
				e.Obj.Update()
			}

			for _, p := range Particles {
				p.Obj.Y -= player.Player.Speed
				p.Obj.Update()
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if c := player.Player.Obj.Check(-player.Player.Speed, 0, "object"); c == nil {
			for _, e := range Enemies {
				e.Obj.X += player.Player.Speed
				e.Obj.Update()
			}

			for _, p := range Particles {
				p.Obj.X += player.Player.Speed
				p.Obj.Update()
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		if c := player.Player.Obj.Check(player.Player.Speed, 0, "object"); c == nil {
			for _, e := range Enemies {
				e.Obj.X -= player.Player.Speed
				e.Obj.Update()
			}

			for _, p := range Particles {
				p.Obj.X -= player.Player.Speed
				p.Obj.Update()
			}
		}
	}

	// Check if a bullet has collided
	for _, b := range player.Bullets {
		xSpeed := b.DirX * 5
		ySpeed := b.DirY * 5

		if c := b.Obj.Check(xSpeed, ySpeed, "object", "zombie", "mini-zombie"); c != nil {
			// If the bullet hits an enemy
			if c.HasTags("zombie", "mini-zombie") {
				for i, e := range Enemies {
					if c.Objects[0].X == e.Obj.X && c.Objects[0].Y == e.Obj.Y {
						Enemies[i].Health -= player.Player.Damage
					}
				}
			}

			tmp := []player.Bullet{}

			for _, B := range player.Bullets {
				if b.Obj.X == B.Obj.X && b.Obj.Y == B.Obj.Y {
					continue
				}

				tmp = append(tmp, B)
			}

			player.Bullets = []player.Bullet{}
			player.Bullets = tmp

			Space.Remove(b.Obj)
			break
		}

		b.Obj.X += xSpeed
		b.Obj.Y += ySpeed

		b.Obj.Update()
	}
}
