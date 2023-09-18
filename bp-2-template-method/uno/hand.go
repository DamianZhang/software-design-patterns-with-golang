package uno

type Hand struct {
	cards []*Card
}

func NewHand() *Hand {
	return &Hand{}
}

func (h *Hand) Size() int {
	return len(h.cards)
}

func (h *Hand) Cards() []*Card {
	return h.cards
}

func (h *Hand) AddHandCard(c *Card) {
	h.cards = append(h.cards, c)
}

func (h *Hand) RemoveHandCard(cardIndex int) {
	lenOfHands := h.Size()

	if lenOfHands > 1 {
		h.cards = append(h.cards[:cardIndex], h.cards[cardIndex+1:]...)
	} else if lenOfHands == 1 {
		h.cards = make([]*Card, 0)
	}
}

func (h *Hand) HasLegalHandCard(cardOnTheTable *Card) bool {
	for _, card := range h.Cards() {
		if card.IsLegalCard(cardOnTheTable) {
			return true
		}
	}

	return false
}
