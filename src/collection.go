package clerk

type Collection struct {
	Database string
	Name     string
}

func NewCollection(name string) *Collection {
	return &Collection{
		Name: name,
	}
}

func NewCollectionWithDatabase(
	database string,
	name string,
) *Collection {
	return &Collection{
		Database: database,
		Name:     name,
	}
}
