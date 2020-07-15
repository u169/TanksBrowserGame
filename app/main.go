package main

import (
	"Tanks/controller"
	"fmt"
	"time"
)

const testGameScale = 10
const testGamePlayersNumber = 1
const testPlayerTransportType = 0

func main() {
	fmt.Println("Started main")

	mController := controller.NewController()
	testGame(mController)

	time.Sleep(time.Second) //TODO remove
}

func testGame(c *controller.Controller) {
	gameId := c.CreateGame(testGameScale, testGamePlayersNumber)
	pId := c.AddPlayer(gameId, testPlayerTransportType)
	fmt.Println(pId)
	go c.StartGame(gameId)
}
