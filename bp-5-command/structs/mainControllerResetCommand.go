package structs

import "fmt"

type MainControllerResetCommand struct {
	mainController   *MainController
	previousCommands *[NUM_OF_KEYS]Command
}

func NewMainControllerResetCommand(mainController *MainController) *MainControllerResetCommand {
	return &MainControllerResetCommand{
		mainController:   mainController,
		previousCommands: new([NUM_OF_KEYS]Command),
	}
}

func (m *MainControllerResetCommand) Execute() {
	m.previousCommands = m.mainController.Commands()
	m.mainController.SetCommands(new([NUM_OF_KEYS]Command))

	fmt.Println("重置 Commands")
}

func (m *MainControllerResetCommand) Undo() {
	m.mainController.SetCommands(m.previousCommands)
	m.previousCommands = new([NUM_OF_KEYS]Command)

	fmt.Println("取消重置 Commands")
}

func (m *MainControllerResetCommand) String() string {
	return "ResetMainControlKeyboard"
}
