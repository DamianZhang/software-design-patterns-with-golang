package structs

type MarcoCommand struct {
	commands []Command
}

func NewMarcoCommand() *MarcoCommand {
	return &MarcoCommand{commands: make([]Command, 0)}
}

func (m *MarcoCommand) Execute() {
	for _, command := range m.commands {
		command.Execute()
	}
}

func (m *MarcoCommand) Undo() {
	for commandIndex := len(m.commands) - 1; commandIndex >= 0; commandIndex-- {
		m.commands[commandIndex].Undo()
	}
}

func (m *MarcoCommand) AddCommand(command Command) {
	m.commands = append(m.commands, command)
}
