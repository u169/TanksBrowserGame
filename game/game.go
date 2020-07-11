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
	iterations := 3
	fmt.Println("Game Started")
	g.w.Draw()

	for i:=0; i < iterations; i++ {
		fmt.Println()
		g.doIteration()
		g.w.Draw()
	}
}

func (g *Game) doIteration() {
	for _, d := range g.w.Dynamics {
		d.Move(g.w.IsDotBusied, g.w.Scale)
	}
}
