package structs

type Command interface {
	Execute()
	Undo()
}
