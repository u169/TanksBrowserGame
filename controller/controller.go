package controller

import (
	"Tanks/game"
	"Tanks/utils"
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

func (c *Controller) CreateGame(worldScale int, playersNumber int) string {
	g := game.NewGame(worldScale, playersNumber)
	gameId := utils.ObjectId(func() chan string {
		stringKeys := make(chan string, len(c.games))
		defer close(stringKeys)
		for key := range c.games {
			stringKeys <- key
		}
		return stringKeys
	}())
	c.games[gameId] = g
	return gameId
}

func (c *Controller) AddPlayer(gameId string, arsenalType int) string {
	return c.games[gameId].AddPlayer(arsenalType)
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
