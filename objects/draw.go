package objects

import "github.com/hajimehoshi/ebiten/v2"

// Draws the trees and scenery
func Draw(screen *ebiten.Image) {
	for _, o := range Objects {
		op := &ebiten.DrawImageOptions{}
		switch o.Type {
		case "tree1":
			op.GeoM.Scale(4, 4)
			op.GeoM.Translate(o.Obj.X-20, o.Obj.Y-52)
			screen.DrawImage(Tree1, op)
		case "tree2":
			op.GeoM.Scale(4, 4)
			op.GeoM.Translate(o.Obj.X-20, o.Obj.Y-52)
			screen.DrawImage(Tree2, op)
		case "tree3":
			op.GeoM.Scale(4, 4)
			op.GeoM.Translate(o.Obj.X-16, o.Obj.Y-50)
			screen.DrawImage(Tree3, op)
		case "tree4":
			op.GeoM.Scale(4, 4)
			op.GeoM.Translate(o.Obj.X-20, o.Obj.Y-52)
			screen.DrawImage(Tree4, op)
		}
	}
}
