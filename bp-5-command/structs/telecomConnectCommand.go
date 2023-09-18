package structs

type TelecomConnectCommand struct {
	telecom *Telecom
}

func NewTelecomConnectCommand(telecom *Telecom) *TelecomConnectCommand {
	return &TelecomConnectCommand{telecom: telecom}
}

func (t *TelecomConnectCommand) Execute() {
	t.telecom.Connect()
}

func (t *TelecomConnectCommand) Undo() {
	t.telecom.Disconnect()
}

func (t *TelecomConnectCommand) String() string {
	return "ConnectTelecom"
}
