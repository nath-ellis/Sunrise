package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/resolv"
)

type Data struct {
	Obj           *resolv.Object
	Speed         float64
	IdleL         *ebiten.Image
	IdleR         *ebiten.Image
	R             []*ebiten.Image
	L             []*ebiten.Image
	IsLeft        bool
	Moving        bool
	MSCool        int
	MoveStage     int
	ShootCool     int
	Damage        int
	Health        int
	ImmunityTicks int
}

type Bullet struct {
	Obj  *resolv.Object
	DirX float64
	DirY float64
}

var (
	Player    Data
	Gun1      *ebiten.Image
	Gun2      *ebiten.Image
	BulletImg *ebiten.Image
	Bullets   []Bullet
)

func Init(Space *resolv.Space) {
	// Import Player's sprites
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

	Player.R = append(Player.R, moving1)
	Player.R = append(Player.R, moving2)
	Player.R = append(Player.R, moving3)
	Player.R = append(Player.R, moving4)
	Player.R = append(Player.R, moving5)
	Player.R = append(Player.R, moving6)
	Player.R = append(Player.R, moving7)
	Player.R = append(Player.R, moving8)
	Player.R = append(Player.R, moving9)
	Player.R = append(Player.R, moving10)
	Player.R = append(Player.R, moving11)
	Player.R = append(Player.R, moving12)

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

	Player.L = append(Player.L, moving1)
	Player.L = append(Player.L, moving2)
	Player.L = append(Player.L, moving3)
	Player.L = append(Player.L, moving4)
	Player.L = append(Player.L, moving5)
	Player.L = append(Player.L, moving6)
	Player.L = append(Player.L, moving7)
	Player.L = append(Player.L, moving8)
	Player.L = append(Player.L, moving9)
	Player.L = append(Player.L, moving10)
	Player.L = append(Player.L, moving11)
	Player.L = append(Player.L, moving12)

	Player.IdleL, _, _ = ebitenutil.NewImageFromFile("assets/player/idleL.png")
	Player.IdleR, _, _ = ebitenutil.NewImageFromFile("assets/player/idleR.png")

	Player.Obj = resolv.NewObject(308, 150, 23, 24, "player")
	Player.Speed = 3
	Player.IsLeft = false
	Player.Moving = false
	Player.Damage = 1
	Player.Health = 10
	Player.ImmunityTicks = 0

	Space.Add(Player.Obj)

	Gun1, _, _ = ebitenutil.NewImageFromFile("assets/gun1.png")
	Gun2, _, _ = ebitenutil.NewImageFromFile("assets/gun2.png")

	BulletImg, _, _ = ebitenutil.NewImageFromFile("assets/bullet.png")
}
