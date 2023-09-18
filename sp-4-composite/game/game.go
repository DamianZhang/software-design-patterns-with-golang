package game

import (
	"fmt"
	"sp-4-composite/logger"
)

type Game struct {
	players []*AI
}

func NewGame() *Game {
	return &Game{
		players: make([]*AI, 0),
	}
}

func (g *Game) SetPlayers(players []*AI) {
	g.players = players
}

func (g *Game) Start() {
	// 取得 "app.game" 日誌器，名為 log 屬性
	log, err := logger.GetLogger("app.game")
	if err != nil {
		fmt.Println("game get logger failed:", err)
		return
	}

	// 四個 AI 玩家，依序命名為 AI 1~4
	players := []*AI{NewAI("AI 1"), NewAI("AI 2"), NewAI("AI 3"), NewAI("AI 4")}
	g.SetPlayers(players)

	// 模擬遊戲執行，請日誌器撰寫日誌訊息，並且做適當的訊息分級
	log.Info(logger.Message("The game begins."))

	// 每個 AI 玩家輪流做決策
	for _, ai := range g.players {
		log.Trace(logger.Message(fmt.Sprintf("The player %s begins his turn.", ai.Name())))
		ai.MakeDecision()
		log.Trace(logger.Message(fmt.Sprintf("The player %s finishes his turn.", ai.Name())))
	}

	log.Debug(logger.Message("Game ends."))
}
