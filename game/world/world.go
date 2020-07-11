package world

import (
	"Tanks/game/world/object"
	"Tanks/game/world/object/dynamic/transport"
	"Tanks/game/world/object/static"
	"errors"
	"fmt"
	"math"
	"math/rand"
)

const decorScalePart = 3
const decorScaleDeviation = 0.05

type World struct {
	Tic      int
	Scale    int
	Entities []object.Entity
}

func NewWorld(scale int) *World {
	w := World{
		Scale:    scale,
	}

	_, _ = w.CreateTransport("tank")
	w.genDecor()
	return &w
}

func (w *World) CreateTransport(t string) (transport.Transport, error) {
	var p transport.Transport

	x, y := w.getFreeCoordinates()

	switch t {
	case "tank":
		p = transport.NewTank("0", x, y)
	default:
		msg := fmt.Sprintf("Unknown type \"%s\"", t)
		return nil, errors.New(msg)
	}

	w.Entities = append(w.Entities, p)
	return p, nil
}

func (w *World) genDecor() {
	var stones []*static.Stone

	for i := 0; i < w.getDecorRange(); i++ {
		stones = append(stones, w.createStone())
	}

	for _, stone := range stones {
		w.Entities = append(w.Entities, stone)
	}
}

func (w *World) getDecorRange() int {
	preVolume := math.Round(decorScalePart * float64(w.Scale))
	deviation := math.Round(decorScaleDeviation * float64(w.Scale))
	max, min := preVolume + deviation, preVolume - deviation
	volume := rand.Intn(int(max) - int(min)) + int(min)
	return volume
}

func (w *World) createStone() *static.Stone {
	x, y := w.getFreeCoordinates()
	return static.NewStone(x, y)
}

func (w *World) isDotBusied(x int, y int) bool {
	for _, e := range w.Entities {
		ex, ey := e.Coordinates()
		if ex == x && ey == y {
			return true
		}
	}
	return false
}

func (w *World) getFreeCoordinates() (int, int) {
	var x, y int
	min, max := 1, w.Scale - 1
	for {
		x = rand.Intn(max-min) + min
		y = rand.Intn(max-min) + min
		if !w.isDotBusied(x, y) {
			break
		}
	}
	return x, y
}

func (w *World) Draw() {
	var entityIndex int
	var symbols = []string{"□", "▦", "■"}
	area := make([][]int, w.Scale)
	for i := 0; i < w.Scale; i++ {
		area[i] = make([]int, w.Scale)
	}

	for _, e := range w.Entities {
		x, y := e.Coordinates()
		switch e.(type) {
		case transport.Transport:
			entityIndex = 2
		case static.Static:
			entityIndex = 1
		default:
			entityIndex = 0
		}
		area[x][y] = entityIndex
	}

	for _, i := range area {
		for _, j := range i {
			fmt.Print(symbols[j], " ")
		}
		fmt.Println()
	}
}