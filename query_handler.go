package clerk

type QueryHandler interface {
	RetrieveAll() (Iterator, error)
	Retrieve(filter interface{}) (Iterator, error)
}