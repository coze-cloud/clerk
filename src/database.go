package clerk

type database struct {
	name string
}

func NewDatabase(name string) *database {
	return &database{
		name: name,
	}
}

func (d *database) Collection(name string) *Collection {
	return &Collection{
		Database: d.name,
		Name:     name,
	}
}
