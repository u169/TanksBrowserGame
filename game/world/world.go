package world

import (
	"Tanks/game/world/object"
	"Tanks/game/world/object/dynamic/arsenal"
	"Tanks/game/world/object/dynamic/arsenal/player"
	"Tanks/game/world/object/static"
	"fmt"
	"math"
	"math/rand"
)

const decorScalePart = 2
const decorScaleDeviation = 0.05

type World struct {
	scale    int
	players  map[string]*player.Player
	statics  []static.Static
}

func NewWorld(scale int, playersInfo map[string]string) *World {
	w := World{
		scale: scale,
	}

	w.genDecor()
	fmt.Println("Decor generated")
	w.genDynamic(playersInfo)

	return &w
}

func (w *World) genDynamic(playersInfo map[string]string) {
	w.genPlayers(playersInfo)
	fmt.Println("Players generated")
}

func (w *World) genPlayers(playersInfo map[string]string) {
	if w.players == nil {
		w.players = map[string]*player.Player{}
	}
	for playerId, transportType := range playersInfo{
		x, y := w.getFreeCoordinates()
		p := player.NewPlayer(transportType, playerId, x, y)
		w.players[playerId] = p
	}
}

func (w *World) genDecor() {
	for i := 0; i < w.getDecorRange(); i++ {
		w.statics = append(w.statics, w.createStone())
	}
}

func (w *World) getDecorRange() int {
	var max, min float64

	maxLimit := math.Pow(float64(w.scale- 2), 2)
	minLimit := 0

	preVolume := math.Round(decorScalePart * float64(w.scale))
	if preVolume > maxLimit {
		preVolume = maxLimit
	}
	deviation := math.Round(decorScaleDeviation * float64(w.scale))
	if deviation == 0 {
		return 1
	}

	max, min = preVolume + deviation, preVolume - deviation
	if max > maxLimit {
		max = maxLimit
	}
	if min < 0 {
		min = float64(minLimit)
	}

	volume := rand.Intn(int(max) - int(min)) + int(min)
	return volume
}

func (w *World) createStone() *static.Stone {
	x, y := w.getFreeCoordinates()
	return static.NewStone(x, y)
}

func (w *World) getFreeCoordinates() (int, int) {
	var x, y int
	min, max := 0, w.scale
	for {
		x = rand.Intn(max-min) + min
		y = rand.Intn(max-min) + min
		if !w.IsDotBusied(x, y) {
			break
		}
	}
	return x, y
}

func (w *World) IsDotBusied(x int, y int) bool {
	for s := range w.getEntities() {
		if isDotBusied(x, y, s) {
			return true
		}
	}
	return false
}

func isDotBusied(x int, y int, e object.Entity) bool {
	ex, ey := e.Coordinates()
	if ex == x && ey == y {
		return true
	}
	return false
}

func (w *World) getEntities() chan object.Entity {
	transports := w.getTransports()
	entities := make(chan object.Entity, len(transports) + len(w.statics))
	defer close(entities)
	for _, v := range w.statics {
		entities <- v
	}
	for t := range transports {
		entities <- t
	}
	return entities
}

func (w *World) getTransports() chan arsenal.Arsenal {
	transports := make(chan arsenal.Arsenal, len(w.players))
	defer close(transports)
	for _, p := range w.players {
		transports <- p.GetTransport()
	}
	return transports
}

func (w *World) Tic() {
	for t := range w.getTransports() {
		t.Move(w.IsDotBusied, w.scale)
	}
}

func (w *World) Update(playerId string, vector int, shoot bool) {
	p := w.players[playerId]
	p.GetTransport().Rotate(vector)
}

//TODO remove
func (w *World) Draw() {
	var entityIndex int
	var symbols = []string{"□", "▦", "■"}
	area := make([][]int, w.scale)
	for i := 0; i < w.scale; i++ {
		area[i] = make([]int, w.scale)
	}
	for e := range w.getEntities() {
		x, y := e.Coordinates()
		switch e.(type) {
		case arsenal.Arsenal:
			entityIndex = 2
		case static.Static:
			entityIndex = 1
		default:
			entityIndex = 0
		}
		area[y][x] = entityIndex
	}

	for _, i := range area {
		for _, j := range i {
			fmt.Print(symbols[j], " ")
		}
		fmt.Println()
	}
}
