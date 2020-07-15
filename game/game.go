package game

import (
	"Tanks/game/player"
	"Tanks/game/world"
	"Tanks/utils"
	"fmt"
)

type Game struct {
	w world.World
	playersNumber int
	players map[string]*player.Player
	arsenalPlayers map[int][]*player.Player
}

func NewGame(worldScale int, playersNumber int) *Game {
	arsenalPlayers := map[int][]*player.Player{
		0: nil,
	}
	return &Game{
		w: *world.NewWorld(worldScale),
		playersNumber: playersNumber,
		players: map[string]*player.Player{},
		arsenalPlayers: arsenalPlayers,
	}
}

func (g *Game) AddPlayer(arsenalType int) string {
	pId := utils.ObjectId(func() chan string {
		stringKeys := make(chan string, len(g.players))
		defer close(stringKeys)
		for key := range g.players {
			stringKeys <- key
		}
		return stringKeys
	}())
	p := player.NewPlayer()
	g.players[pId] = p
	g.arsenalPlayers[arsenalType] = append(g.arsenalPlayers[arsenalType], p)

	return pId
}

func (g *Game) Start() {
	g.w.GenerateObjects(g.arsenalPlayers)

	// TODO
	fmt.Println("Game Started")
	g.w.Draw()
	fmt.Println()
	g.w.Tic()
	g.w.Draw()
	fmt.Println()
	g.w.Tic()
	g.w.Draw()
}

func (g *Game) Update(playerId string, vector int, shoot bool) {
	g.players[playerId].Update(vector, shoot)
}
