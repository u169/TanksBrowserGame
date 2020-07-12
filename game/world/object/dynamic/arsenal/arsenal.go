package arsenal

import (
	"Tanks/game/world/object/dynamic"
	"Tanks/game/world/object/dynamic/arsenal/transport"
	"errors"
	"fmt"
)

type Arsenal interface {
	dynamic.Dynamic
	dynamic.Alive
	dynamic.Shooting
}

func NewTransport(tType string, x int, y int) (Arsenal, error) {
	switch tType {
	case tType:
		return transport.NewTank(x, y), nil
	default:
		errMsg := fmt.Sprintf("Invalid arsenal \"%s\"", tType)
		return nil, errors.New(errMsg)
	}
}



