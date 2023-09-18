package showdown

import (
	"errors"
	"fmt"
)

type Card struct {
	suit Suit
	rank Rank
}

func NewCard(suit Suit, rank Rank) (*Card, error) {
	card := &Card{}

	if err := card.setSuit(suit); err != nil {
		return card, err
	}

	if err := card.setRank(rank); err != nil {
		return card, err
	}

	return card, nil
}

func (c *Card) Showdown(cc *Card) bool {
	if c.rank > cc.rank {
		return true
	} else if c.rank == cc.rank && c.suit > cc.suit {
		return true
	} else {
		return false
	}
}

func (c *Card) setSuit(suit Suit) error {
	if ShouldBeInRange(int(suit), 0, 3) {
		c.suit = suit
		return nil
	}

	return errors.New("card's suit should be in range 0-3")
}

func (c *Card) setRank(rank Rank) error {
	if ShouldBeInRange(int(rank), 2, 14) {
		c.rank = rank
		return nil
	}

	return errors.New("card's rank should be in range 2-14")
}

func (c *Card) String() string {
	return fmt.Sprintf("%s:%s", c.suit, c.rank)
}
