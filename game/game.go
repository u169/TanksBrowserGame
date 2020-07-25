package game

import (
	"Tanks/game/player"
	"Tanks/game/world"
	"Tanks/game/world/object/static"
	"Tanks/utils"
	"time"
)

const (
	gameStateCreated = 0
	gameStateProcess = 1
	gameStateEnded   = 2
)
const gameTickPeriod = time.Millisecond * 500

type Game struct {
	state int
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
		state:          gameStateCreated,
		w:              *world.NewWorld(worldScale),
		playersNumber:  playersNumber,
		players:        map[string]*player.Player{},
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
	g.state = gameStateProcess
	g.w.GenerateObjects(g.arsenalPlayers)
	go g.process()
}

func (g *Game) process() {
	for range time.Tick(gameTickPeriod) {
		if g.isEnded() {
			g.state = gameStateEnded
			return
		} else {
			g.w.Tick()
		}
	}
}

func (g *Game) isEnded() bool {
	var counter int
	for _, p := range g.players {
		if p != nil {
			counter++
		}
	}
	return counter <= 1
}

func (g *Game) Update(playerId string, vector int, shoot bool) {
	g.players[playerId].Update(vector, shoot)
}

func (g *Game) Get() {
	type GameInfo struct {
		Players []*player.Player
		Stones []*static.Stone
		Info struct{}
	}
}
