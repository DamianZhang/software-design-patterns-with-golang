package main

import (
	"bp-5-command/structs"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("keyboard shortcut setter is starting...")

	mainController := structs.NewMainController()

	tank := structs.NewTank()
	telecom := structs.NewTelecom()

	tankMFC := structs.NewTankMoveForwardCommand(tank)
	tankMBC := structs.NewTankMoveBackwardCommand(tank)
	telecomCC := structs.NewTelecomConnectCommand(telecom)
	telecomDC := structs.NewTelecomDisconnectCommand(telecom)
	mainControllerRC := structs.NewMainControllerResetCommand(mainController)
	commands := []structs.Command{tankMFC, tankMBC, telecomCC, telecomDC, mainControllerRC}

	mainController.AddCommand(structs.F, tankMFC)
	mainController.AddCommand(structs.B, tankMBC)
	mainController.AddCommand(structs.C, telecomCC)
	mainController.AddCommand(structs.D, telecomDC)
	mainController.AddCommand(structs.R, mainControllerRC)

	var (
		input string
		exit  bool = false
	)

	for !exit {
		fmt.Println("(1) 快捷鍵設置 (2) Undo (3) Redo (4) 按下按鍵 (5) 退出程序:")
		fmt.Scanf("%s", &input)

		switch input {
		case "1":
			setCommand(mainController, commands)
		case "2":
			mainController.Press(structs.Z)
		case "3":
			mainController.Press(structs.Y)
		case "4":
			handleClickKey(mainController)
		case "5":
			exit = true
		default:
			fmt.Println("請輸入數字 1-5 以執行程序")
		}
	}
}

func scanKeyOfCommand(promptWord string) (structs.Key, error) {
	var alphabet string
	fmt.Println(promptWord)
	fmt.Scanf("%s", &alphabet)

	keyOfCommand, err := structs.ConvertAlphabetToKey(alphabet)
	if err != nil {
		return keyOfCommand, err
	}

	return keyOfCommand, nil
}

func printCommands(promptWord string, commands []structs.Command) {
	fmt.Println(promptWord)

	for commandIndex, command := range commands {
		fmt.Printf("(%d) %s\n", commandIndex, command)
	}
}

func setCommand(mainController *structs.MainController, commands []structs.Command) {
	var input string
	fmt.Println("設定巨集指令(y/n):")
	fmt.Scanf("%s", &input)

	switch input {
	case "y":
		setMarcoCommand(mainController, commands)
	case "n":
		setNormalCommand(mainController, commands)
	default:
		fmt.Println("指令設置失敗，請輸入(y/n):")
	}
}

func setMarcoCommand(mainController *structs.MainController, commands []structs.Command) {
	keyOfCommand, err := scanKeyOfCommand("請輸入欲設置巨集的快捷鍵(A-Z):")
	if err != nil {
		fmt.Println(err)
		return
	}

	promptWord := fmt.Sprintf("要將哪些指令設置成快捷鍵 %s 的巨集(輸入多個數字，以 , 隔開):", keyOfCommand)
	printCommands(promptWord, commands)

	marcoCommand := structs.NewMarcoCommand()

	var strOfCommandIndexes string
	fmt.Scanf("%s", &strOfCommandIndexes)

	strSliceOfCommandIndexes := strings.Split(strOfCommandIndexes, ",")
	for _, strOfCommandIndex := range strSliceOfCommandIndexes {
		commandIndex, err := strconv.Atoi(strOfCommandIndex)
		if err != nil || commandIndex >= len(commands) {
			fmt.Println("巨集指令設置失敗，請輸入合理的選項")
			return
		}

		marcoCommand.AddCommand(commands[commandIndex])
	}

	mainController.AddCommand(keyOfCommand, marcoCommand)
	fmt.Println("巨集指令設置成功")
}

func setNormalCommand(mainController *structs.MainController, commands []structs.Command) {
	keyOfCommand, err := scanKeyOfCommand("請輸入欲設置的快捷鍵(A-Z):")
	if err != nil {
		fmt.Println(err)
		return
	}

	promptWord := fmt.Sprintf("要將哪一道指令設置到快捷鍵 %s 上:\n", keyOfCommand)
	printCommands(promptWord, commands)

	var strOfCommandIndex string
	fmt.Scanf("%s", &strOfCommandIndex)

	commandIndex, err := strconv.Atoi(strOfCommandIndex)
	if err != nil || commandIndex >= len(commands) {
		fmt.Println("指令設置失敗，請輸入合理的選項")
		return
	}

	mainController.AddCommand(keyOfCommand, commands[commandIndex])
	fmt.Println("指令設置成功")
}

func handleClickKey(mainController *structs.MainController) {
	keyOfCommand, err := scanKeyOfCommand("請輸入按鍵(A-Z):")
	if err != nil {
		fmt.Println(err)
		return
	}

	mainController.Press(keyOfCommand)
}
