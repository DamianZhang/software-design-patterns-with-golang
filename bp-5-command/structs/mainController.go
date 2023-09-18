package structs

import "fmt"

type MainController struct {
	keyboard         *Keyboard
	commands         *[NUM_OF_KEYS]Command
	executedCommands []Command
	undoCommands     []Command
}

func NewMainController() *MainController {
	return &MainController{
		keyboard:         NewKeyboard(),
		commands:         new([NUM_OF_KEYS]Command),
		executedCommands: make([]Command, 0),
		undoCommands:     make([]Command, 0),
	}
}

func (m *MainController) Press(key Key) {
	switch key {
	case Z:
		m.undo()
	case Y:
		m.redo()
	default:
		k := m.keyboard.Keys()[key]
		command := m.commands[k]

		if command != nil {
			command.Execute()

			m.executedCommands = append(m.executedCommands, command)
			m.undoCommands = make([]Command, 0)
		} else {
			fmt.Println("【該按鍵沒有指令】")
		}
	}
}

func (m *MainController) undo() {
	lenOfExecutedCommands := len(m.executedCommands)

	if lenOfExecutedCommands > 0 {
		previousExecutedCommand := m.executedCommands[lenOfExecutedCommands-1]

		previousExecutedCommand.Undo()

		m.executedCommands = m.executedCommands[0 : lenOfExecutedCommands-1]
		m.undoCommands = append(m.undoCommands, previousExecutedCommand)
	} else {
		fmt.Println("沒有指令可以 Undo")
	}
}

func (m *MainController) redo() {
	lenOfUndoCommands := len(m.undoCommands)

	if lenOfUndoCommands > 0 {
		previousUndoCommand := m.undoCommands[lenOfUndoCommands-1]

		previousUndoCommand.Execute()

		m.undoCommands = m.undoCommands[0 : lenOfUndoCommands-1]
		m.executedCommands = append(m.executedCommands, previousUndoCommand)
	} else {
		fmt.Println("沒有指令可以 Redo")
	}
}

func (m *MainController) AddCommand(key Key, command Command) {
	k := m.keyboard.Keys()[key]
	m.commands[k] = command
}

func (m *MainController) RemoveCommand(key Key) {
	k := m.keyboard.Keys()[key]
	m.commands[k] = nil
}

func (m *MainController) SetCommands(commands *[NUM_OF_KEYS]Command) {
	m.commands = commands
}

func (m *MainController) Commands() *[NUM_OF_KEYS]Command {
	return m.commands
}
