package player

import "Tanks/game/world/object/dynamic/arsenal"

type Player struct {
	Id        string
	transport arsenal.Arsenal
}

func NewPlayer(transportType string, id string, x int, y int) *Player {
	t, _ := arsenal.NewTransport(transportType, x, y)
	return &Player{
		Id:        id,
		transport: t,
	}
}

func (p *Player) GetTransport() arsenal.Arsenal {
	return p.transport
}
