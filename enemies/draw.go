package enemies

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

// Draws the enemies
func Draw(screen *ebiten.Image) {
	for _, e := range Enemies {
		op := &ebiten.DrawImageOptions{}
		switch e.Type {
		case "zombie":
			op.GeoM.Scale(2, 2)
			op.GeoM.Translate(e.Obj.X, e.Obj.Y)
			screen.DrawImage(Zombie, op)
		case "mini-zombie":
			op.GeoM.Translate(e.Obj.X, e.Obj.Y)
			screen.DrawImage(Zombie, op)
		}
	}
}

// For drawing the particles when an enemy dies
func DrawParticles(screen *ebiten.Image, Space *resolv.Space) {
	for i, p := range Particles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(p.Obj.X, p.Obj.Y)

		screen.DrawImage(ParticleImg, op)
		Particles[i].Timer -= 1

		if p.Timer <= 0 {
			Space.Remove(p.Obj)
			tmp := []Particle{}

			for _, P := range Particles {
				if p.Obj.X == P.Obj.X && p.Obj.Y == P.Obj.Y {
					continue
				}

				tmp = append(tmp, P)
			}
			Particles = []Particle{}
			Particles = tmp
			break
		}
	}
}
