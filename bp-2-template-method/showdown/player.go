package showdown

import (
	"fmt"
	"math/rand"
)

type Player struct {
	name  string
	point int
	hand  *Hand
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

func (p *Player) TakeTurn() *Card {
	return p.ShowCard()
}

// TODO: 抽象方法
func (p *Player) ShowCard() *Card {
	var card *Card
	lenOfHand := p.hand.Size()
	cardsOfHand := p.hand.Cards()

	if lenOfHand > 0 {
		cardIndex := rand.Intn(lenOfHand)
		card = cardsOfHand[cardIndex]
		p.hand.RemoveHandCard(cardIndex)
	}

	return card
}

func (p *Player) GainPoint() {
	p.point++
}

func (p *Player) Point() int {
	return p.point
}

func (p *Player) Name() string {
	return p.name
}
