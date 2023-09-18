package uno

import (
	"fmt"
	"math/rand"
)

type Player struct {
	name string
	hand *Hand
}

func NewPlayer() *Player {
	return &Player{hand: NewHand()}
}

// TODO: 抽象方法
func (p *Player) NameHimself() {
	var name string

	fmt.Printf("請輸入玩家的暱稱:\n")
	fmt.Scanln(&name)

	p.name = name
}

func (p *Player) AddHandCard(c *Card) {
	p.hand.AddHandCard(c)
}

func (p *Player) TakeTurn(cardOnTheTable *Card) *Card {
	return p.ShowCard(cardOnTheTable)
}

// TODO: 抽象方法
func (p *Player) ShowCard(cardOnTheTable *Card) *Card {
	var card *Card
	lenOfHand := p.hand.Size()
	cardsOfHand := p.hand.Cards()

	if p.HasLegalHandCard(cardOnTheTable) {
		for {
			cardIndex := rand.Intn(lenOfHand)
			card = cardsOfHand[cardIndex]
			if card.IsLegalCard(cardOnTheTable) {
				p.hand.RemoveHandCard(cardIndex)
				break
			}
		}
	}

	return card
}

func (p *Player) HasLegalHandCard(cardOnTheTable *Card) bool {
	return p.hand.HasLegalHandCard(cardOnTheTable)
}

func (p *Player) IsEmptyHandCard() bool {
	return p.hand.Size() == 0
}

func (p *Player) Name() string {
	return p.name
}
