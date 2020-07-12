package world

import (
	"Tanks/game/world/object"
	"Tanks/game/world/object/dynamic"
	"Tanks/game/world/object/dynamic/transport"
	"Tanks/game/world/object/static"
	"errors"
	"fmt"
	"math"
	"math/rand"
)

const decorScalePart = 2
const decorScaleDeviation = 0.05

type World struct {
	scale    int
	statics  []static.Static
	dynamics []dynamic.Dynamic
}

func NewWorld(scale int) *World {
	w := World{
		scale: scale,
	}

	_, _ = w.CreateTransport("tank")
	fmt.Println("Tank generated")
	w.genDecor()
	fmt.Println("Decor generated")
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

	w.dynamics = append(w.dynamics, p)
	return p, nil
}

func (w *World) genDecor() {
	var stones []*static.Stone

	for i := 0; i < w.getDecorRange(); i++ {
		stones = append(stones, w.createStone())
	}

	for _, stone := range stones {
		w.statics = append(w.statics, stone)
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

func (w *World) IsDotBusied(x int, y int) bool {
	for _, s := range w.getEntities() {
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

func (w *World) getEntities() []object.Entity {

	var entities []object.Entity
	for _, v := range w.statics {
		entities = append(entities, v)
	}
	for _, v := range w.dynamics {
		entities = append(entities, v)
	}
	return entities
}

func (w *World) Tic() {
	for _, d := range w.dynamics {
		d.Move(w.IsDotBusied, w.scale)
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
	for _, e := range w.getEntities() {
		x, y := e.Coordinates()
		switch e.(type) {
		case transport.Transport:
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
