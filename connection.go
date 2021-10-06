package clerk

type Connection interface {
	SendQuery(query Query) (Iterator, error)
	SendCommand(command Command) error

	Close(handler func(err error))
}