package enemies

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
)

type Enemy struct {
	Obj    *resolv.Object
	Type   string
	Speed  int
	Health int
}

type Particle struct {
	Obj   *resolv.Object
	Timer int
}

var (
	Enemies           []Enemy
	Particles         []Particle
	Wave              int = 0
	WaveCounter       int = 0 // change later
	Zombie, _, _          = ebitenutil.NewImageFromFile("assets/enemies/zombie.png")
	ParticleImg, _, _     = ebitenutil.NewImageFromFile("assets/particle.png")
)
