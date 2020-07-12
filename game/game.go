package game

import (
	"Tanks/game/world"
	"fmt"
)

type Game struct {
	w world.World
}

func NewGame(worldScale int) *Game {
	return &Game{
		w: *world.NewWorld(worldScale),
	}
}

func (g *Game) Start() {
	fmt.Println("Game Started")
	g.testProcess(3)
}

//TODO remove
func (g *Game) testProcess(tics int) {
	g.w.Draw()
	for i:=0; i < tics; i++ {
		fmt.Println()
		g.w.Tic()
		g.w.Draw()
	}
}
