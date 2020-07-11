package main

import (
	"Tanks/game"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	worldScale := 10

	g := game.NewGame(worldScale)
	g.Start()
}
