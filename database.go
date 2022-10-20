package clerk

type Database struct {
	Name string
}

func NewDatabase(name string) *Database {
	return &Database{
		Name: name,
	}
}

func (d *Database) Collection(name string) *Collection {
	return NewCollectionWithDatabase(d.Name, name)
}
