package clerk

type Connection interface {
	SendQuery(query Query) (Iterator, error)
	SendCommand(command Command) error

	Close(errorHandler func(err error))
}