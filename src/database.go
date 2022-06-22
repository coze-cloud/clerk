package clerk

type Database struct {
	name string
}

func NewDatabase(name string) *Database {
	return &Database{
		name: name,
	}
}

func (d *Database) Collection(name string) *Collection {
	return NewCollectionWithDatabase(d.name, name)
}
