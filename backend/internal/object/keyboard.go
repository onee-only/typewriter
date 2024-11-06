package object

const (
	KeyboardHeight = 5
	KeyboardWidth  = 13
)

/*
This is the initial state of the keyboard.
[KeyboardHeight] X [KeyboardWidth] sized.

# Caps Lock OFF
┌───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┐
| ` │ 1 │ 2 │ 3 │ 4 │ 5 │ 6 │ 7 │ 8 │ 9 │ 0 │ - │ = │
├───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┤
│ Q │ W │ E │ R │ T │ Y │ U │ I │ O │ P │ [ │ ] │ \ │
├───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┤
│ A │ S │ D │ F │ G │ H │ J │ K │ L │ ; │ ' │   │   │
├───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┤
│ Z │ X │ C │ V │ B │ N │ M │ , │ . │ / │   │   |   |
├───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┤
│                    SPACE BAR                      │
└───────────────────────────────────────────────────┘

# Caps Lock ON
┌───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┐
| ~ │ ! │ @ │ # │ $ │ % │ ^ │ & │ * │ ( │ ) │ _ │ + │
├───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┤
│ Q │ W │ E │ R │ T │ Y │ U │ I │ O │ P │ { │ } │ | │
├───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┤
│ A │ S │ D │ F │ G │ H │ J │ K │ L │ : │ " │   │   │
├───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┤
│ Z │ X │ C │ V │ B │ N │ M │ < │ > │ ? │   │   |   |
├───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┤
│                    SPACE BAR                      │
└───────────────────────────────────────────────────┘
*/

const _nullByte byte = 0

var initialGrid KeyLayout = [KeyboardHeight][KeyboardWidth]byte{
	{'`', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '-', '='},
	{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', '[', ']', '\\'},
	{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', ';', '\'', _nullByte, _nullByte},
	{'z', 'x', 'c', 'v', 'b', 'n', 'm', ',', '.', '/', _nullByte, _nullByte, _nullByte},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
}

var initialCapGrid KeyLayout = [KeyboardHeight][KeyboardWidth]byte{
	{'~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+'},
	{'Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P', '{', '}', '|'},
	{'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', ':', '"', _nullByte, _nullByte},
	{'Z', 'X', 'C', 'V', 'B', 'N', 'M', '<', '>', '?', _nullByte, _nullByte, _nullByte},
	{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
}

type KeyLayout [KeyboardHeight][KeyboardWidth]byte

type Keyboard struct {
	cursor        Point
	grid, capGrid KeyLayout
}

func NewKeyboard() Keyboard {
	const (
		halfHeight = KeyboardHeight / 2
		halfWidth  = KeyboardWidth / 2
	)

	return Keyboard{
		cursor:  Point{Y: halfHeight, X: halfWidth},
		grid:    initialGrid,
		capGrid: initialCapGrid,
	}
}

func (kb *Keyboard) Layout() (grid, capGrid KeyLayout) {
	return kb.grid, kb.capGrid
}

// Cursor returns the current position of cursor.
func (kb *Keyboard) Cursor() Point {
	return kb.cursor
}

// MoveCursor moves cursor 1 block to the dir.
// If it tries to leave the boundary, it stays at the edge of the keyboard.
func (kb *Keyboard) MoveCursor(dir Direction) {
	cursor := kb.cursor.Move(dir, 1)
	if kb.inBounds(cursor) {
		kb.setCursor(cursor)
	}
}

// Press presses the key where the cursor is currently at.
// When capsLock is true, the key in alternative layout will be pressed.
// If the cursor is at the empty slot, ok will be false.
func (kb *Keyboard) Press(capsLock bool) (ch byte, ok bool) {
	if capsLock {
		ch = kb.capGrid[kb.cursor.Y][kb.cursor.X]
	} else {
		ch = kb.grid[kb.cursor.Y][kb.cursor.X]
	}

	return ch, ch != _nullByte
}

func (kb *Keyboard) setCursor(cursor Point) {
	kb.cursor = cursor
}

func (kb *Keyboard) inBounds(cursor Point) (ok bool) {
	y := cursor.Y >= 0 && cursor.Y < KeyboardHeight
	x := cursor.X >= 0 && cursor.X < KeyboardWidth

	return y && x
}
