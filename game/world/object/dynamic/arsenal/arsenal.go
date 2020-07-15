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

func NewTransport(tType int, x int, y int) (Arsenal, error) {
	var arsenal Arsenal
	var err error
	switch tType {
	case 0:
		arsenal = transport.NewTank(x, y)
	default:
		errMsg := fmt.Sprintf("Invalid arsenal \"%s\"", tType)
		err = errors.New(errMsg)
	}
	return arsenal, err
}



