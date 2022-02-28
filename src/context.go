package clerk

type Creator interface {
	Create(collection *Collection, data interface{}) error
}

type Updater interface {
	Update(collection *Collection, filter map[string]interface{}, data interface{}, upsert bool) error
}

type Queryer interface {
	Query(collection *Collection, filter map[string]interface{}) ([]interface{}, error)
}
