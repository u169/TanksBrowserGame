package game

import (
	"Tanks/game/world"
	"fmt"
)

type Game struct {
	w world.World
}

func NewGame(worldScale int, playersInfo map[string]string) *Game {
	return &Game{
		w: *world.NewWorld(worldScale, playersInfo),
	}
}

func (g *Game) Start() {
	fmt.Println("Game Started")
	g.testProcess(3)
}

func (g *Game) Update(playerId string, vector int, shoot bool) {
	g.w.Update(playerId, vector, shoot)
}

//TODO remove
func (g *Game) testProcess(tics int) {
	g.w.Draw()

	var i int

	for i=0; i < tics; i++ {
		fmt.Println()
		g.w.Tic()
		g.w.Draw()
	}
}
