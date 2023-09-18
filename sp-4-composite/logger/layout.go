package logger

type Layout interface {
	ModifyMessageFormat(levelOfMessage Level, nameOfLogger string, message Message) (formatedMessage Message)
}
