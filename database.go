package clerk

type Database struct {
	Name string
}

func NewDatabase(name string) Database {
	database := new(Database)

	database.Name = name

	return *database
}

func (database Database) GetCollection(name string) Collection {
	return NewCollection(database, name)
}