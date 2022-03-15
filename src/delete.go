package clerk

type delete struct {
	collection *Collection
	filter     map[string]interface{}
}

func NewDelete(collection *Collection) *delete {
	return &delete{
		collection: collection,
		filter:     map[string]interface{}{},
	}
}

func (d *delete) Where(key string, value interface{}) *delete {
	d.filter[key] = value
	return d
}

func (d *delete) Execute(deleter Deleter) error {
	return deleter.Delete(d.collection, d.filter)
}
