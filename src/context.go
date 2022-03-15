package clerk

type Creator interface {
	Create(collection *Collection, data interface{}) error
}

type Updater interface {
	Update(collection *Collection, filter map[string]interface{}, data interface{}, upsert bool) error
}

type Deleter interface {
	Delete(collection *Collection, filter map[string]interface{}) error
}

type Querier interface {
	Query(collection *Collection, filter map[string]interface{}, results interface{}) error
}
