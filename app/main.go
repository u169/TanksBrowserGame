package main

import (
	"Tanks/controller"
	"fmt"
	"time"
)

const testGameId = "testGameId"
const testGameScale = 10
const testPlayerId = "testPlayerId"
const testPlayerTransportType = "tank"

func main() {
	fmt.Println("Started main")

	mController := controller.NewController()
	testGame(mController)

	time.Sleep(time.Second) //TODO remove
}

func testGame(c *controller.Controller) {
	c.CreateGame(
		testGameId,
		map[string]string{
			testPlayerId: testPlayerTransportType,
		},
		testGameScale)
	c.StartGame(testGameId)
}
