package object

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type KeyboardTestSuite struct {
	suite.Suite
	kb Keyboard
}

func TestKeyboardTestSuite(t *testing.T) {
	suite.Run(t, new(KeyboardTestSuite))
}

func (s *KeyboardTestSuite) SetupTest() {
	s.kb = NewKeyboard()
}

func (s *KeyboardTestSuite) TestInitialLayout() {
	layout, capLayout := s.kb.Layout()
	s.Equal(initialGrid, layout)
	s.Equal(initialCapGrid, capLayout)
}

func (s *KeyboardTestSuite) TestCursorInitialPosition() {
	point := s.kb.Cursor()
	s.Require().Equal(KeyboardHeight/2, point.Y)
	s.Require().Equal(KeyboardWidth/2, point.X)
}

func (s *KeyboardTestSuite) TestCursorMove() {
	initial := s.kb.Cursor()

	s.kb.MoveCursor(DirectionDown)

	current := s.kb.Cursor()

	s.Equal(initial.Y+1, current.Y)
}

func (s *KeyboardTestSuite) TestCursorMoveInBoundary() {
	// Move all the way to the left.
	for s.kb.Cursor().X > 0 {
		s.kb.MoveCursor(DirectionLeft)
	}

	if !s.Zero(s.kb.Cursor().X) {
		return
	}

	// Trying to move out of the boundary.
	s.kb.MoveCursor(DirectionLeft)

	// Check if it stays inside.
	s.Equal(0, s.kb.Cursor().X)
}

func (s *KeyboardTestSuite) TestPress() {
	initial := s.kb.Cursor()
	// Position where the slot is empty.
	// NOTE: This can be changed if requirement changes.
	emptySlotPos := Point{Y: 2, X: 11}

	for _, tc := range []struct {
		desc         string
		pos          Point
		capsLock     bool
		expectedChar byte
		expectedOK   bool
	}{
		{
			desc:         "initial position & caps lock off",
			pos:          initial,
			capsLock:     false,
			expectedChar: s.kb.grid[initial.Y][initial.X],
			expectedOK:   true,
		},
		{
			desc:         "initial position & caps lock on",
			pos:          initial,
			capsLock:     true,
			expectedChar: s.kb.capGrid[initial.Y][initial.X],
			expectedOK:   true,
		},
		{
			desc:         "empty slot position & caps lock off",
			pos:          emptySlotPos,
			capsLock:     false,
			expectedChar: _nullByte,
			expectedOK:   false,
		},
		{
			desc:         "empty slot position & caps lock on",
			pos:          emptySlotPos,
			capsLock:     true,
			expectedChar: _nullByte,
			expectedOK:   false,
		},
	} {
		s.Run(tc.desc, func() {
			s.kb.setCursor(tc.pos)
			ch, ok := s.kb.Press(tc.capsLock)
			s.Equal(tc.expectedOK, ok)
			s.Equal(tc.expectedChar, ch)
		})
	}
}
