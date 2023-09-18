package main

import (
	"fmt"
	"sp-4-composite/game"
	"sp-4-composite/logger"
)

func main() {
	BasicAnswer()
	// AdvancedAnswer()
}

func BasicAnswer() {
	// 定義根日誌器
	root := logger.NewRootLogger()

	// 定義 app.game 日誌器，繼承根日誌器並覆寫分級門檻和輸出器
	gameLogger, err := logger.NewNormalLogger(
		"app.game",
		root,
		logger.WithLevelThreshold(logger.INFO),
		logger.WithExporter(
			logger.NewCompositeExporter(
				logger.NewConsoleExporter(),
				logger.NewCompositeExporter(
					logger.NewFileExporter("./game.log"),
					logger.NewFileExporter("./game.backup.log")))),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 定義 app.game.ai 日誌器，繼承 app.game 日誌器並覆寫分級門檻
	aiLogger, err := logger.NewNormalLogger(
		"app.game.ai",
		gameLogger,
		logger.WithLevelThreshold(logger.TRACE),
		logger.WithLayout(logger.NewStandardLayout()),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 配置剛定義好的三個日誌器
	logger.DeclareLoggers(root, gameLogger, aiLogger)

	// 創建遊戲物件，並執行遊戲
	game := game.NewGame()
	game.Start()
}

func AdvancedAnswer() {
	// 以 JSON 格式的檔案來定義和配置所有日誌器
	jsonParser := logger.NewJSONParser()
	err := jsonParser.Parse("./loggers.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 創建遊戲物件，並執行遊戲
	game := game.NewGame()
	game.Start()
}
