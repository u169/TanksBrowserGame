package main

import (
	"Tanks/game"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Started")

	rand.Seed(time.Now().UTC().UnixNano())

	worldScale := 10

	g := game.NewGame(worldScale)
	g.Start()
}
