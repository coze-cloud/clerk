package clerk

type Database struct {
	name string
}

func NewDatabase(name string) Database {
	return Database{name: name}
}

func (d Database) GetCollection(name string) Collection {
	return NewCollection(d, name)
}