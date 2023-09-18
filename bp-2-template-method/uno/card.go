package uno

import (
	"errors"
	"fmt"
)

type Card struct {
	color  Color
	number Number
}

func NewCard(color Color, number Number) (*Card, error) {
	card := &Card{}

	if err := card.setColor(color); err != nil {
		return card, err
	}

	if err := card.setNumber(number); err != nil {
		return card, err
	}

	return card, nil
}

func (c *Card) IsLegalCard(cardOnTheTable *Card) bool {
	return c.color == cardOnTheTable.color || c.number == cardOnTheTable.number
}

func (c *Card) setColor(color Color) error {
	if ShouldBeInRange(int(color), 0, 3) {
		c.color = color
		return nil
	}

	return errors.New("card's color should be in range 0-3")
}

func (c *Card) setNumber(number Number) error {
	if ShouldBeInRange(int(number), 0, 9) {
		c.number = number
		return nil
	}

	return errors.New("card's number should be in range 0-9")
}

func (c *Card) String() string {
	return fmt.Sprintf("%s:%d", c.color, c.number)
}
