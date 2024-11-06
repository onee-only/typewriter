package object

// Point represents a point in a 2D space.
type Point struct {
	Y, X int
}

var dirs = [4]Point{
	{Y: -1, X: 0}, // [DirectionUp]
	{Y: 0, X: 1},  // [DirectionRight]
	{Y: 1, X: 0},  // [DirectionDown]
	{Y: 0, X: -1}, // [DirectionLeft]
}

func (p Point) Move(dir Direction, delta int) Point {
	d := dirs[dir]
	p.Y += d.Y * delta
	p.X += d.X * delta

	return p
}
