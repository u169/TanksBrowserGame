package transport

import "Tanks/game/world/object/dynamic"

type Transport interface {
	dynamic.Player
	dynamic.Dynamic
	dynamic.Alive
	dynamic.Shooting
}
