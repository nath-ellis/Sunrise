package main

import (
	_ "image/png"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/nath-ellis/Sunrise/enemies"
	"github.com/nath-ellis/Sunrise/objects"
	"github.com/nath-ellis/Sunrise/player"
	"github.com/nath-ellis/Sunrise/ui"
	"github.com/solarlune/resolv"
)

var (
	State string = "menu"
	BG    *ebiten.Image
	Space *resolv.Space
	Ticks int = 0
)

type Game struct{}

// Imports assets and prepares the game
func init() {
	BG, _, _ = ebitenutil.NewImageFromFile("assets/bg.png")

	Space = resolv.NewSpace(1280, 720, 1, 1)

	player.Init(Space)

	rand.Seed(time.Now().Unix())
}

func (g *Game) Update() error {
	switch State {
	case "menu":
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			objects.AddTrees(Space)
			State = "game"
		}
	case "game":
		Ticks += 1

		// Triggers every second
		if (Ticks % 60) == 0 {
			enemies.WaveCounter -= 1
		}

		// New wave
		if enemies.WaveCounter <= 0 {
			enemies.NewWave(Space)
		}

		// Starts a wave sooner if all enemies are dead
		if len(enemies.Enemies) <= 0 {
			enemies.WaveCounter -= int(enemies.WaveCounter / 5)
		}

		// When health is at 0
		if player.Player.Health <= 0 {
			State = "gameOver"
		}

		enemies.Update(Space)
		player.Controls()
		player.Shoot(Space)
	case "gameOver":
		// Resets
		Ticks = 0
		enemies.Wave = 1
		enemies.WaveCounter = 0

		// Removes hitboxes
		for _, e := range enemies.Enemies {
			Space.Remove(e.Obj)
		}
		for _, b := range player.Bullets {
			Space.Remove(b.Obj)
		}
		for _, p := range enemies.Particles {
			Space.Remove(p.Obj)
		}
		for _, o := range objects.Objects {
			Space.Remove(o.Obj)
		}

		// Removes images
		enemies.Enemies = []enemies.Enemy{}
		player.Bullets = []player.Bullet{}
		enemies.Particles = []enemies.Particle{}
		objects.Objects = []objects.Object{}

		// Starts the game again
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			objects.AddTrees(Space)

			player.Player.Health = 10
			player.Player.ImmunityFrames = 0

			State = "game"
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(BG, nil)

	switch State {
	case "menu":
		ui.DrawMenu(screen)
	case "game":
		enemies.DrawParticles(screen, Space)
		player.DrawBullets(screen)
		player.DrawWeapon(screen)
		player.Draw(screen)
		enemies.Draw(screen)
		objects.Draw(screen)
		ui.DrawHealth(screen)
	case "gameOver":
		ui.DrawGameOver(screen)
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
