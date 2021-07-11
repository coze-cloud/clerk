package clerk

type Connection interface {
	SendQuery() error
	SendCommand(command Command) error

	Close(errorHandler func(err error))
}