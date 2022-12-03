package objects

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
)

type Object struct {
	Obj  *resolv.Object
	Type string
}

var (
	Objects     []Object
	Tree1, _, _     = ebitenutil.NewImageFromFile("assets/tree1.png")
	Tree2, _, _     = ebitenutil.NewImageFromFile("assets/tree2.png")
	Tree3, _, _     = ebitenutil.NewImageFromFile("assets/tree3.png")
	Tree4, _, _     = ebitenutil.NewImageFromFile("assets/tree4.png")
	SpawnRangeX int = 3840
	SpawnRangeY int = 2160
	TreeAmount  int = 400
)
