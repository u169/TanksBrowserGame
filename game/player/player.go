package player

import "Tanks/game/world/object/dynamic/arsenal"

type Player struct {
	transport arsenal.Arsenal
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) SetTransport(a arsenal.Arsenal) {
	p.transport = a
}

func (p *Player) Update(vector int, shoot bool) {
	p.transport.Rotate(vector)
	p.transport.SetShoot(shoot)
}
