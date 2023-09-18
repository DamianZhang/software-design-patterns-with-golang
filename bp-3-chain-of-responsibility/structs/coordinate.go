package structs

import "errors"

type Coordinate struct {
	x int
}

func NewCoordinate(x int) (*Coordinate, error) {
	c := &Coordinate{}

	if !ShouldBeInRange(x, 0, 29) {
		return c, errors.New("coordinate's x should be in range 0-29")
	}

	c.x = x
	return c, nil
}

func (c *Coordinate) X() int {
	return c.x
}
