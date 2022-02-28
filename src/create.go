package clerk

type create struct {
	collection *Collection
	data       interface{}
}

func NewCreate(collection *Collection, data interface{}) *create {
	return &create{
		collection: collection,
		data:       data,
	}
}

func (c *create) Execute(creator Creator) error {
	return creator.Create(c.collection, c.data)
}
