package clerk

type query struct {
	collection *Collection
	filter     map[string]interface{}
}

func NewQuery(collection *Collection) *query {
	return &query{
		collection: collection,
		filter:     map[string]interface{}{},
	}
}

func (q *query) Where(key string, value interface{}) *query {
	q.filter[key] = value
	return q
}

func (q *query) Execute(querier Querier) ([]interface{}, error) {
	return querier.Query(q.collection, q.filter)
}
