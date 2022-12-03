package enemies

import (
	"math/rand"

	"github.com/solarlune/resolv"
)

// Creates a new wave
func NewWave(Space *resolv.Space) {
	for i := 0; i < Wave*3; i++ {
		c := rand.Intn(5)
		r := rand.Intn(4)

		if c == 1 {
			if r == 1 || r == 2 {
				Enemies = append(Enemies, Enemy{resolv.NewObject(float64(rand.Intn(640)+640), float64(rand.Intn(360)+360), 28, 32, "zombie"), "zombie", 1, 4})
			} else {
				Enemies = append(Enemies, Enemy{resolv.NewObject(float64(rand.Intn(640)+640), float64(rand.Intn(360)+360), 14, 16, "mini-zombie"), "mini-zombie", 2, 2})
			}
		} else if c == 2 {
			if r == 1 || r == 2 {
				Enemies = append(Enemies, Enemy{resolv.NewObject(-float64(rand.Intn(640)), float64(rand.Intn(360)+360), 28, 32, "zombie"), "zombie", 1, 4})
			} else {
				Enemies = append(Enemies, Enemy{resolv.NewObject(-float64(rand.Intn(640)), float64(rand.Intn(360)+360), 14, 16, "mini-zombie"), "mini-zombie", 2, 2})
			}
		} else if c == 3 {
			if r == 1 || r == 2 {
				Enemies = append(Enemies, Enemy{resolv.NewObject(float64(rand.Intn(640)+640), -float64(rand.Intn(360)), 28, 32, "zombie"), "zombie", 1, 4})

			} else {
				Enemies = append(Enemies, Enemy{resolv.NewObject(float64(rand.Intn(640)+640), -float64(rand.Intn(360)), 14, 16, "mini-zombie"), "mini-zombie", 2, 2})
			}
		} else {
			if r == 1 || r == 2 {
				Enemies = append(Enemies, Enemy{resolv.NewObject(-float64(rand.Intn(640)), -float64(rand.Intn(360)), 28, 32, "zombie"), "zombie", 1, 4})
			} else {
				Enemies = append(Enemies, Enemy{resolv.NewObject(-float64(rand.Intn(640)), -float64(rand.Intn(360)), 14, 16, "mini-zombie"), "mini-zombie", 2, 2})
			}
		}
	}

	for _, e := range Enemies {
		Space.Add(e.Obj)
	}

	WaveCounter = 120
	Wave += 1
}
