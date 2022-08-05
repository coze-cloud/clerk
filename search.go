package clerk

import "context"

type search[T any] struct {
	collection  *Collection
	query       string
	highlight   []string
	filterable  []string
	filterQuery string
	skip        int
	take        int
}

func NewSearch[T any](collection *Collection, query string) *search[T] {
	return &search[T]{
		collection: collection,
		query:      query,
		skip:       -1,
		take:       -1,
	}
}

func (s *search[T]) Filter(expr string) *search[T] {
	s.filterQuery = expr
	return s
}

func (s *search[T]) By(key string) *search[T] {
	if s.filterable == nil {
		s.filterable = []string{}
	}

	s.filterable = append(s.filterable, key)
	return s
}

func (s *search[T]) Highlight(key string) *search[T] {
	if s.highlight == nil {
		s.highlight = []string{}
	}

	s.highlight = append(s.highlight, key)
	return s
}

func (s *search[T]) Skip(skip int) *search[T] {
	s.skip = skip
	return s
}

func (s *search[T]) Take(take int) *search[T] {
	s.take = take
	return s
}

func (s *search[T]) Execute(ctx context.Context, searcher Searcher[T]) (<-chan T, error) {
	return searcher.Search(
		ctx,
		s.collection,
		s.query,
		s.highlight,
		s.filterable,
		s.filterQuery,
		s.skip,
		s.take,
	)
}
