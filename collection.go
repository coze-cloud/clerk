package clerk

type Collection struct {
	database Database
	name     string
}

func NewCollection(database Database, name string) Collection {
	return Collection{database: database, name: name}
}
