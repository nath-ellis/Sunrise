package objects

import (
	"math/rand"

	"github.com/solarlune/resolv"
)

func AddTrees(Space *resolv.Space) {
	// Adds trees (right and down)
	for i := 0; i < TreeAmount/4; i++ {
		Objects = append(Objects, Object{resolv.NewObject(float64(rand.Intn(SpawnRangeX)), float64(rand.Intn(SpawnRangeY)), 2, 1, "object"), "tree1"})
		Objects = append(Objects, Object{resolv.NewObject(float64(rand.Intn(SpawnRangeX)), float64(rand.Intn(SpawnRangeY)), 2, 1, "object"), "tree2"})
		Objects = append(Objects, Object{resolv.NewObject(float64(rand.Intn(SpawnRangeX)), float64(rand.Intn(SpawnRangeY)), 2, 2, "object"), "tree3"})
		Objects = append(Objects, Object{resolv.NewObject(float64(rand.Intn(SpawnRangeX)), float64(rand.Intn(SpawnRangeY)), 2, 1, "object"), "tree4"})
	}
	// Left and Down
	for i := 0; i < TreeAmount/4; i++ {
		Objects = append(Objects, Object{resolv.NewObject(-float64(rand.Intn(SpawnRangeX)), float64(rand.Intn(SpawnRangeY)), 2, 1, "object"), "tree1"})
		Objects = append(Objects, Object{resolv.NewObject(-float64(rand.Intn(SpawnRangeX)), float64(rand.Intn(SpawnRangeY)), 2, 1, "object"), "tree2"})
		Objects = append(Objects, Object{resolv.NewObject(-float64(rand.Intn(SpawnRangeX)), float64(rand.Intn(SpawnRangeY)), 2, 2, "object"), "tree3"})
		Objects = append(Objects, Object{resolv.NewObject(-float64(rand.Intn(SpawnRangeX)), float64(rand.Intn(SpawnRangeY)), 2, 1, "object"), "tree4"})
	}
	// Right and Up
	for i := 0; i < TreeAmount/4; i++ {
		Objects = append(Objects, Object{resolv.NewObject(float64(rand.Intn(SpawnRangeX)), -float64(rand.Intn(SpawnRangeY)), 2, 1, "object"), "tree1"})
		Objects = append(Objects, Object{resolv.NewObject(float64(rand.Intn(SpawnRangeX)), -float64(rand.Intn(SpawnRangeY)), 2, 1, "object"), "tree2"})
		Objects = append(Objects, Object{resolv.NewObject(float64(rand.Intn(SpawnRangeX)), -float64(rand.Intn(SpawnRangeY)), 2, 2, "object"), "tree3"})
		Objects = append(Objects, Object{resolv.NewObject(float64(rand.Intn(SpawnRangeX)), -float64(rand.Intn(SpawnRangeY)), 2, 1, "object"), "tree4"})
	}
	// Left and Up
	for i := 0; i < TreeAmount/4; i++ {
		Objects = append(Objects, Object{resolv.NewObject(-float64(rand.Intn(SpawnRangeX)), -float64(rand.Intn(SpawnRangeY)), 2, 1, "object"), "tree1"})
		Objects = append(Objects, Object{resolv.NewObject(-float64(rand.Intn(SpawnRangeX)), -float64(rand.Intn(SpawnRangeY)), 2, 1, "object"), "tree2"})
		Objects = append(Objects, Object{resolv.NewObject(-float64(rand.Intn(SpawnRangeX)), -float64(rand.Intn(SpawnRangeY)), 2, 2, "object"), "tree3"})
		Objects = append(Objects, Object{resolv.NewObject(-float64(rand.Intn(SpawnRangeX)), -float64(rand.Intn(SpawnRangeY)), 2, 1, "object"), "tree4"})
	}

	// Adds hitboxes
	for _, o := range Objects {
		Space.Add(o.Obj)
	}
}