package dynamic

import "Tanks/game/world/object"

type Dynamic interface {
	object.Entity
	Move(isBusied func(int, int) bool, scale int)
	Rotate(vector int)
}

type Alive interface {
	GetHP() int
	EditHP(int)
}

type Shooting interface {
	GetAmmo() int
	EditAmmo(int)
	SetShoot(bool)
	GetShoot() bool
}

type Identified interface {
	GetId() string
}