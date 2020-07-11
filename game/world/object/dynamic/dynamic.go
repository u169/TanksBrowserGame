package dynamic

import "Tanks/game/world/object"

type Dynamic interface {
	object.Entity
	Move(available func(int, int) bool, scale int)
}

type Alive interface {
	GetHP() int
	EditHP(int)
}

type Shooting interface {
	GetAmmo() int
	EditAmmo(int)
}

type Player interface {
	GetId() string
}