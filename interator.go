package clerk

type Iterator interface {
	Next() bool
	Decode(entity interface{}) error
	Close()

	Single(entity interface{}) error
}