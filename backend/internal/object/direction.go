package object

import (
	"github.com/onee-only/typewriter/backend/internal/util/enum"
	"github.com/pkg/errors"
)

type Direction uint8

const (
	DirectionUp Direction = iota
	DirectionRight
	DirectionDown
	DirectionLeft
)

var _ enum.Validator = Direction(0)

func (d Direction) String() string {
	switch d {
	case DirectionUp:
		return "up"
	case DirectionRight:
		return "right"
	case DirectionDown:
		return "down"
	case DirectionLeft:
		return "left"
	}

	return "invalid"
}

var ErrDirectionInvalid = errors.Errorf("direction should be in range of %d ~ %d",
	DirectionUp, DirectionLeft,
)

func (d Direction) Valid() error {
	if d < DirectionUp || d > DirectionLeft {
		return ErrDirectionInvalid
	}

	return nil
}
