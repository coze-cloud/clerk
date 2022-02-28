package clerk

type update struct {
	collection *Collection
	filter     map[string]interface{}
	data       interface{}
	upsert     bool
}

func NewUpdate(collection *Collection, data interface{}) *update {
	return &update{
		collection: collection,
		filter:     map[string]interface{}{},
		data:       data,
		upsert:     false,
	}
}

func (c *update) Where(key string, value interface{}) *update {
	c.filter[key] = value
	return c
}

func (c *update) WithUpsert() *update {
	c.upsert = true
	return c
}

func (c *update) Execute(updater Updater) error {
	return updater.Update(c.collection, c.filter, c.data, c.upsert)
}
