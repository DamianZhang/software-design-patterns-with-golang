package showdown

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
	fmt.Println("showdown is starting...")

	// 玩家依序取名字
	for _, player := range g.players {
		player.NameHimself()
	}

	// 洗牌
	g.deck.Shuffle()

	// 玩家輪流抽牌，直到每個玩家手牌有 13 張為止
	for i := 1; i <= 13; i++ {
		for _, player := range g.players {
			player.AddHandCard(g.deck.DrawCard())
		}
	}

	// 遊戲出牌階段
	cardsOfTheLatestRound := [4]*Card{}
	for round := 1; round <= 13; round++ {
		// 玩家輪流出牌
		for playerIndex, player := range g.players {
			card := player.TakeTurn()
			cardsOfTheLatestRound[playerIndex] = card
		}

		// 顯示 P1~P4 各出的牌的內容
		fmt.Printf("【ROUND %d】\n", round)
		for playerIndex, card := range cardsOfTheLatestRound {
			player := g.players[playerIndex]
			fmt.Printf("玩家 %s 打出 %s\n", player.Name(), card)
		}
		fmt.Println()

		// 將 P1~P4 出的牌進行比大小決勝負，將最勝者的分數加一
		winnerOfTheLatestRound := 0
		maxCard := cardsOfTheLatestRound[0]
		for playerIndex, card := range cardsOfTheLatestRound {
			if !maxCard.Showdown(card) {
				maxCard = card
				winnerOfTheLatestRound = playerIndex
			}
		}
		g.players[winnerOfTheLatestRound].GainPoint()
	}

	// 結算階段，分數最高者為贏家
	winnerOfGame := g.players[0]
	for _, player := range g.players {
		fmt.Printf("玩家: %s, 分數: %d\n", player.Name(), player.Point())
		if winnerOfGame.Point() < player.Point() {
			winnerOfGame = player
		}
	}
	fmt.Printf("贏家為:【%s】\n", winnerOfGame.Name())
}
