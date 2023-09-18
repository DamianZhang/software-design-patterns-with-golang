package structs

import (
	"fmt"
)

type Coordinate struct {
	row int
	col int
}

func NewCoordinate(row, col int) (*Coordinate, error) {
	c := &Coordinate{}

	if !ShouldBeInRange(row, 0, RowOfGameMap-1) || !ShouldBeInRange(col, 0, ColOfGameMap-1) {
		return c, fmt.Errorf("row and col should be in range 0-%d", ColOfGameMap-1)
	}

	c.row = row
	c.col = col
	return c, nil
}

func (c *Coordinate) Row() int {
	return c.row
}

func (c *Coordinate) Col() int {
	return c.col
}
