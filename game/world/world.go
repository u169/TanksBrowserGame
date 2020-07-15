package world

import (
	"Tanks/game/player"
	"Tanks/game/world/object"
	"Tanks/game/world/object/dynamic"
	"Tanks/game/world/object/dynamic/arsenal"
	"Tanks/game/world/object/static"
	"fmt"
	"math"
	"math/rand"
)

const decorScalePart = 2
const decorScaleDeviation = 0.05

type World struct {
	scale    int
	entities []object.Entity
}

func NewWorld(scale int) *World {
	return &World{scale: scale}
}

func (w *World) GenerateObjects(arsenalPlayers map[int][]*player.Player) {
	w.genDynamic(arsenalPlayers)
	w.genDecor()
}

func (w *World) genDynamic(arsenalPlayers map[int][]*player.Player) {
	w.genTransport(arsenalPlayers)
	fmt.Println("Players generated")
}

func (w *World) genTransport(arsenalPlayers map[int][]*player.Player) {
	for tType, tPlayers := range arsenalPlayers {
		for _, tP := range tPlayers {
			x, y := w.getFreeCoordinates()
			transport, _ := arsenal.NewTransport(tType, x, y)
			tP.SetTransport(transport)
			w.entities = append(w.entities, transport)
		}
	}
}

func (w *World) genDecor() {
	for i := 0; i < w.getDecorRange(); i++ {
		w.entities = append(w.entities, w.createStone())
	}
}

func (w *World) getDecorRange() int {
	var max, min float64

	maxLimit := math.Pow(float64(w.scale - 2), 2)
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
	return static.NewStone(w.getFreeCoordinates())
}

func (w *World) getFreeCoordinates() (int, int) {
	var x, y int
	min, max := 0, w.scale
	for {
		x = rand.Intn(max-min) + min
		y = rand.Intn(max-min) + min
		if !w.isDotBusied(x, y) {
			break
		}
	}
	return x, y
}

func (w *World) isDotBusied(x int, y int) bool {
	for _, s := range w.entities {
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

func (w *World) Tic() {
	for _, entity := range w.entities {
		switch entity.(type) {
		case dynamic.Dynamic:
			d := entity.(dynamic.Dynamic)
			d.Move(w.isDotBusied, w.scale)
		}
	}
}

//TODO remove
func (w *World) Draw() {
	var entityIndex int
	var symbols = []string{"□", "▦", "■"}
	area := make([][]int, w.scale)
	for i := 0; i < w.scale; i++ {
		area[i] = make([]int, w.scale)
	}
	for _, e := range w.entities {
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
