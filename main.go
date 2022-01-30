package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
)

type Player struct {
	Obj       *resolv.Object
	Speed     float64
	IdleL     *ebiten.Image
	IdleR     *ebiten.Image
	R         []*ebiten.Image
	L         []*ebiten.Image
	Left      bool
	Moving    bool
	MSCool    int
	MoveStage int
}

type Object struct {
	Obj  *resolv.Object
	Type string
}

var (
	State   string = "menu"
	BG      *ebiten.Image
	player  Player
	Space   *resolv.Space
	Objects []Object
	Tree1   *ebiten.Image
	Tree2   *ebiten.Image
	Tree3   *ebiten.Image
	Tree4   *ebiten.Image
)

type Game struct{}

func charImports() {
	moving1, _, _ := ebitenutil.NewImageFromFile("assets/player/right/1.png")
	moving2, _, _ := ebitenutil.NewImageFromFile("assets/player/right/2.png")
	moving3, _, _ := ebitenutil.NewImageFromFile("assets/player/right/3.png")
	moving4, _, _ := ebitenutil.NewImageFromFile("assets/player/right/4.png")
	moving5, _, _ := ebitenutil.NewImageFromFile("assets/player/right/5.png")
	moving6, _, _ := ebitenutil.NewImageFromFile("assets/player/right/6.png")
	moving7, _, _ := ebitenutil.NewImageFromFile("assets/player/right/7.png")
	moving8, _, _ := ebitenutil.NewImageFromFile("assets/player/right/8.png")
	moving9, _, _ := ebitenutil.NewImageFromFile("assets/player/right/9.png")
	moving10, _, _ := ebitenutil.NewImageFromFile("assets/player/right/10.png")
	moving11, _, _ := ebitenutil.NewImageFromFile("assets/player/right/11.png")
	moving12, _, _ := ebitenutil.NewImageFromFile("assets/player/right/12.png")

	player.R = append(player.R, moving1)
	player.R = append(player.R, moving2)
	player.R = append(player.R, moving3)
	player.R = append(player.R, moving4)
	player.R = append(player.R, moving5)
	player.R = append(player.R, moving6)
	player.R = append(player.R, moving7)
	player.R = append(player.R, moving8)
	player.R = append(player.R, moving9)
	player.R = append(player.R, moving10)
	player.R = append(player.R, moving11)
	player.R = append(player.R, moving12)

	moving1, _, _ = ebitenutil.NewImageFromFile("assets/player/left/1.png")
	moving2, _, _ = ebitenutil.NewImageFromFile("assets/player/left/2.png")
	moving3, _, _ = ebitenutil.NewImageFromFile("assets/player/left/3.png")
	moving4, _, _ = ebitenutil.NewImageFromFile("assets/player/left/4.png")
	moving5, _, _ = ebitenutil.NewImageFromFile("assets/player/left/5.png")
	moving6, _, _ = ebitenutil.NewImageFromFile("assets/player/left/6.png")
	moving7, _, _ = ebitenutil.NewImageFromFile("assets/player/left/7.png")
	moving8, _, _ = ebitenutil.NewImageFromFile("assets/player/left/8.png")
	moving9, _, _ = ebitenutil.NewImageFromFile("assets/player/left/9.png")
	moving10, _, _ = ebitenutil.NewImageFromFile("assets/player/left/10.png")
	moving11, _, _ = ebitenutil.NewImageFromFile("assets/player/left/11.png")
	moving12, _, _ = ebitenutil.NewImageFromFile("assets/player/left/12.png")

	player.L = append(player.L, moving1)
	player.L = append(player.L, moving2)
	player.L = append(player.L, moving3)
	player.L = append(player.L, moving4)
	player.L = append(player.L, moving5)
	player.L = append(player.L, moving6)
	player.L = append(player.L, moving7)
	player.L = append(player.L, moving8)
	player.L = append(player.L, moving9)
	player.L = append(player.L, moving10)
	player.L = append(player.L, moving11)
	player.L = append(player.L, moving12)

	player.IdleL, _, _ = ebitenutil.NewImageFromFile("assets/player/idleL.png")
	player.IdleR, _, _ = ebitenutil.NewImageFromFile("assets/player/idleR.png")
}

func init() {
	BG, _, _ = ebitenutil.NewImageFromFile("assets/bg.png")

	Space = resolv.NewSpace(1280, 720, 1, 1)

	player.Obj = resolv.NewObject(308, 150, 23, 24, "player")
	player.Speed = 5
	player.Left = false
	player.Moving = false

	Space.Add(player.Obj)

	Tree1, _, _ = ebitenutil.NewImageFromFile("assets/tree1.png")
	Tree2, _, _ = ebitenutil.NewImageFromFile("assets/tree2.png")
	Tree3, _, _ = ebitenutil.NewImageFromFile("assets/tree3.png")
	Tree4, _, _ = ebitenutil.NewImageFromFile("assets/tree4.png")

	Objects = append(Objects, Object{resolv.NewObject(float64(60), float64(60), 2, 1, "object"), "tree1"})
	Objects = append(Objects, Object{resolv.NewObject(float64(150), float64(70), 2, 1, "object"), "tree2"})
	Objects = append(Objects, Object{resolv.NewObject(float64(210), float64(80), 2, 2, "object"), "tree3"})
	Objects = append(Objects, Object{resolv.NewObject(float64(300), float64(90), 2, 1, "object"), "tree4"})

	for _, o := range Objects {
		Space.Add(o.Obj)
	}

	charImports()
}

func drawObjects(screen *ebiten.Image) {
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

func move() {
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		if c := player.Obj.Check(0, -player.Speed, "object"); c == nil {
			for _, o := range Objects {
				o.Obj.Y += player.Speed
				o.Obj.Update()
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		if c := player.Obj.Check(0, player.Speed, "object"); c == nil {
			for _, o := range Objects {
				o.Obj.Y -= player.Speed
				o.Obj.Update()
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		player.Left = true

		if c := player.Obj.Check(-player.Speed, 0, "object"); c == nil {
			for _, o := range Objects {
				o.Obj.X += player.Speed
				o.Obj.Update()
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		player.Left = false

		if c := player.Obj.Check(player.Speed, 0, "object"); c == nil {
			for _, o := range Objects {
				o.Obj.X -= player.Speed
				o.Obj.Update()
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) ||
		ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) ||
		ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		player.Moving = true
	} else {
		player.Moving = false
	}

}

func drawPlayer(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(player.Obj.X, player.Obj.Y)

	if player.Moving {
		if player.MSCool <= 0 {
			player.MoveStage += 1
			player.MSCool = 1
		} else {
			if player.MSCool > 0 {
				player.MSCool -= 1
			}
		}

		if player.MoveStage >= 12 {
			player.MoveStage = 0
		}

		if !player.Left {
			screen.DrawImage(player.R[player.MoveStage], op)
		} else {
			screen.DrawImage(player.L[player.MoveStage], op)
		}
	} else {
		if player.Left {
			screen.DrawImage(player.IdleL, op)
		} else {
			screen.DrawImage(player.IdleR, op)
		}
	}
}

func (g *Game) Update() error {
	switch State {
	case "menu":
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			State = "game"
		}
	case "game":
		move()
	case "gameOver":
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(BG, nil)

	switch State {
	case "menu":
	case "game":
		drawPlayer(screen)
		drawObjects(screen)
	case "gameOver":
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 360
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Sunrise")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Panic("Failed to run game: ", err)
	}
}
