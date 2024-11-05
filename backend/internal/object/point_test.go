package object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointMove(t *testing.T) {
	point := Point{Y: 1, X: 1}

	for _, tc := range []struct {
		desc   string
		dir    Direction
		delta  int
		expect Point
	}{
		{
			desc:   "down",
			dir:    DirectionDown,
			delta:  1,
			expect: Point{Y: 2, X: 1},
		},
		{
			desc:   "up",
			dir:    DirectionUp,
			delta:  1,
			expect: Point{Y: 0, X: 1},
		},
		{
			desc:   "left",
			dir:    DirectionLeft,
			delta:  1,
			expect: Point{Y: 1, X: 0},
		},
		{
			desc:   "right",
			dir:    DirectionRight,
			delta:  3,
			expect: Point{Y: 1, X: 4},
		},
	} {
		t.Run("move "+tc.desc, func(t *testing.T) {
			p := point.Move(tc.dir, tc.delta)
			assert.Equal(t, tc.expect, p)
		})
	}
}
