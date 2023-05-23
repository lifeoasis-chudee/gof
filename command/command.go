package command

type Command interface {
	Execute()
	Undo()
}

type CommandFunc func()

func (c CommandFunc) Execute() {
	c()
}

func (c CommandFunc) Undo() {
	c()
}
