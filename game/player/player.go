package player

import "Tanks/game/world/object/dynamic/arsenal"

type Player struct {
	Id string `json:"id"`
	Transport arsenal.Arsenal `json:"transport"`
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) SetTransport(a arsenal.Arsenal) {
	p.Transport = a
}

func (p *Player) Update(vector int, shoot bool) {
	p.Transport.Rotate(vector)
	p.Transport.SetShoot(shoot)
}
