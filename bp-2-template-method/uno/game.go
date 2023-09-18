package uno

import "fmt"

type Game struct {
	deck    *Deck
	players []*Player
}

func NewGame(players []*Player) (*Game, error) {
	game := &Game{}

	deck, err := NewDeck()
	if err != nil {
		return game, err
	}

	game.deck = deck
	game.players = players

	return game, nil
}

func (g *Game) Start() {
	fmt.Println("UNO is starting...")

	// 玩家依序取名字
	for _, player := range g.players {
		player.NameHimself()
	}

	// 洗牌
	g.deck.Shuffle()

	// 玩家輪流抽牌，直到每個玩家手牌有 5 張為止
	for i := 1; i <= 5; i++ {
		for _, player := range g.players {
			player.AddHandCard(g.deck.DrawCard())
		}
	}

	fmt.Printf("g.deck.cards: %v\n", g.deck.cards)
	// 遊戲階段
	// 1. 從牌堆中翻出第一張牌到檯面上。
	cardOnTheTable := g.deck.DrawCard()
	cardsOfMuck := make([]*Card, 0)
	var winnerOfGame *Player = nil

	for winnerOfGame == nil {
		for _, player := range g.players {
			fmt.Printf("玩家【%s】手牌: %v\n", player.Name(), player.hand.Cards())
		}

		// 2. 由 P1 開始，出牌順序為 P1 → P2 → P3 → P4 → P1 以此類推。
		for _, player := range g.players {
			fmt.Printf("檯面上的牌: 【%v】\n", cardOnTheTable)

			// 3. 玩家出的牌必須與檯面上最新的牌的顏色一樣，或是數字一樣。出完的牌就會成為檯面上最新的牌。
			if player.HasLegalHandCard(cardOnTheTable) {
				cardsOfMuck = append(cardsOfMuck, cardOnTheTable)
				cardOnTheTable = player.TakeTurn(cardOnTheTable)
				fmt.Printf("玩家【%s】打出 %s\n", player.Name(), cardOnTheTable)

				// 4. 最快出完手中牌的人為遊戲的贏家。
				if player.IsEmptyHandCard() {
					winnerOfGame = player
					break
				}
			} else {
				// 5. 如果玩家沒有任何可出的牌，玩家就必須從牌堆中抽一張牌，
				// 如果此時牌堆空了，則會先把檯面上除了最新的牌以外的牌放回牌堆中進行洗牌。
				if g.deck.IsEmptyCards() {
					fmt.Println("牌堆空了，將棄牌區的牌加回牌堆後進行洗牌")
					g.deck.SetCards(cardsOfMuck)
					cardsOfMuck = make([]*Card, 0)
					fmt.Println("新的牌堆:", g.deck.cards)
					fmt.Println("新的棄牌區:", cardsOfMuck)
					g.deck.Shuffle()
				}

				fmt.Printf("玩家【%s】沒牌可出，抽一張牌\n", player.Name())
				player.AddHandCard(g.deck.DrawCard())
			}
		}

		fmt.Println()
	}

	fmt.Printf("贏家為:【%s】\n", winnerOfGame.Name())
}
