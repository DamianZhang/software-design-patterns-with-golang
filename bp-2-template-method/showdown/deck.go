package showdown

import (
	"math/rand"
	"time"
)

type Deck struct {
	cards []*Card
}

func NewDeck() (*Deck, error) {
	cards := make([]*Card, 0)

	for suit := 0; suit <= 3; suit++ {
		for rank := 2; rank <= 14; rank++ {
			card, err := NewCard(Suit(suit), Rank(rank))
			if err != nil {
				return &Deck{}, err
			}

			cards = append(cards, card)
		}
	}

	return &Deck{cards: cards}, nil
}

func (d *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *Deck) DrawCard() *Card {
	var card *Card

	if len(d.cards) > 0 {
		card = d.cards[0]
		d.cards = d.cards[1:]
	}

	return card
}
