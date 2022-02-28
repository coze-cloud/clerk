package clerk

type Collection struct {
	Database string
	Name     string
}

func NewCollection(database string, name string) *Collection {
	return &Collection{
		Database: database,
		Name:     name,
	}
}
