package controller

import (
	"Tanks/game"
	"math/rand"
	"time"
)

type PlayerInfo struct {
	PlayerId string `json:"player_id"`
	Vector int `json:"vector"`
	Shot bool `json:"shot"`
}

type Controller struct {
	games map[string]*game.Game
}

func NewController() *Controller {
	rand.Seed(time.Now().UTC().UnixNano())
	cGames := map[string]*game.Game{}
	return &Controller{games: cGames}
}

func (c *Controller) CreateGame(
	gameId string,
	playersInfo map[string]string,
	worldScale int) {
	g := game.NewGame(
		worldScale,
		playersInfo)
	c.games[gameId] = g
}

func (c *Controller) StartGame(gameId string) {
	c.games[gameId].Start()
}

func (c *Controller) Update(gameId string, pInfo PlayerInfo) {
	g := c.games[gameId]
	g.Update(
		pInfo.PlayerId,
		pInfo.Vector,
		pInfo.Shot)
}
