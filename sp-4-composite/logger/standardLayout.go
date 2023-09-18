package logger

import (
	"fmt"
	"time"
)

type StandardLayout struct{}

func NewStandardLayout() *StandardLayout {
	return &StandardLayout{}
}

func (s *StandardLayout) ModifyMessageFormat(levelOfMessage Level, nameOfLogger string, message Message) (formatedMessage Message) {
	return Message(fmt.Sprintf("%s\t|-%s\t%s\t-\t%s\n",
		time.Now().Format("2006-1-2 15:04:05.000"),
		levelOfMessage,
		nameOfLogger,
		message))
}
