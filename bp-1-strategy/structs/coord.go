package structs

import "fmt"

type Coord struct {
	x, y float64
}

func NewCoord(x, y float64) *Coord {
	return &Coord{
		x: x,
		y: y,
	}
}

func (c *Coord) X() float64 {
	return c.x
}

func (c *Coord) Y() float64 {
	return c.y
}

func (c *Coord) String() string {
	return fmt.Sprintf("(%v, %v)", c.x, c.y)
}
