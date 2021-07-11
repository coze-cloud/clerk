package clerk

type Collection struct {
	Database Database
	Name string
}

func NewCollection(database Database, name string) Collection {
	collection := new(Collection)

	collection.Database = database
	collection.Name = name

	return *collection
}