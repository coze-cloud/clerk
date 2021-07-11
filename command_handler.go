package clerk

type CommandHandler interface {
	Create(entity interface{}) error
	Update(filter interface{}, entity interface{}) error
	Delete(filter interface{}) error
}
