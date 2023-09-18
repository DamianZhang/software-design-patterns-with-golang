package structs

type TelecomDisconnectCommand struct {
	telecom *Telecom
}

func NewTelecomDisconnectCommand(telecom *Telecom) *TelecomDisconnectCommand {
	return &TelecomDisconnectCommand{telecom: telecom}
}

func (t *TelecomDisconnectCommand) Execute() {
	t.telecom.Disconnect()
}

func (t *TelecomDisconnectCommand) Undo() {
	t.telecom.Connect()
}

func (t *TelecomDisconnectCommand) String() string {
	return "DisconnectTelecom"
}
