package enemies

import (
	"math/rand"

	"github.com/solarlune/resolv"
)

// Creates a new wave
func NewWave(Space *resolv.Space) {
	for i := 0; i < Wave*3; i++ {
		location := rand.Intn(5)
		enemyType := rand.Intn(4)

		switch location {
		case 1: // Enemies spawn below and to the right of the player
			if enemyType == 1 {
				Enemies = append(Enemies, Enemy{resolv.NewObject(float64(rand.Intn(640)+640), float64(rand.Intn(360)+360), 14, 16, "mini-zombie"), "mini-zombie", 2, 2})
			} else {
				Enemies = append(Enemies, Enemy{resolv.NewObject(float64(rand.Intn(640)+640), float64(rand.Intn(360)+360), 28, 32, "zombie"), "zombie", 1, 4})
			}
		case 2: // Enemies spawn above and to the right of the player
			if enemyType == 1 {
				Enemies = append(Enemies, Enemy{resolv.NewObject(-float64(rand.Intn(640)), float64(rand.Intn(360)+360), 14, 16, "mini-zombie"), "mini-zombie", 2, 2})
			} else {
				Enemies = append(Enemies, Enemy{resolv.NewObject(-float64(rand.Intn(640)), float64(rand.Intn(360)+360), 28, 32, "zombie"), "zombie", 1, 4})
			}
		case 3: // Enemies spawn below and to the left of the player
			if enemyType == 1 {
				Enemies = append(Enemies, Enemy{resolv.NewObject(float64(rand.Intn(640)+640), -float64(rand.Intn(360)), 14, 16, "mini-zombie"), "mini-zombie", 2, 2})
			} else {
				Enemies = append(Enemies, Enemy{resolv.NewObject(float64(rand.Intn(640)+640), -float64(rand.Intn(360)), 28, 32, "zombie"), "zombie", 1, 4})
			}
		case 4: // Enemies spawn above and to the left of the player
			if enemyType == 1 {
				Enemies = append(Enemies, Enemy{resolv.NewObject(-float64(rand.Intn(640)), -float64(rand.Intn(360)), 14, 16, "mini-zombie"), "mini-zombie", 2, 2})
			} else {
				Enemies = append(Enemies, Enemy{resolv.NewObject(-float64(rand.Intn(640)), -float64(rand.Intn(360)), 28, 32, "zombie"), "zombie", 1, 4})
			}
		}
	}

	for _, e := range Enemies {
		Space.Add(e.Obj)
	}

	WaveCounter = 120
	Wave += 1
}
