package game

import (
	"fmt"
	"sp-4-composite/logger"
)

type AI struct {
	name string
}

func NewAI(name string) *AI {
	return &AI{
		name: name,
	}
}

func (ai *AI) Name() string {
	return ai.name
}

// 模擬 AI 決策，請日誌器撰寫日誌訊息，並做適當的訊息分級
func (ai *AI) MakeDecision() {
	var (
		// 取得 "app.game.ai" 日誌器，名為 log 屬性
		log, err = logger.GetLogger("app.game.ai")
		name     = ai.name
	)
	if err != nil {
		fmt.Printf("%s get logger failed: %s\n", name, err)
		return
	}

	log.Trace(logger.Message(fmt.Sprintf("%s starts making decisions...", name)))

	log.Warn(logger.Message(fmt.Sprintf("%s decides to give up.", name)))
	log.Error(logger.Message("Something goes wrong when AI gives up."))

	log.Trace(logger.Message(fmt.Sprintf("%s completes its decision.", name)))
}
