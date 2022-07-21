package meilisearch

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	clerk "github.com/Becklyn/clerk/src"
	"github.com/meilisearch/meilisearch-go"
)

var (
	ErrIdFilterIsRequired = errors.New("filter id is required")
)

type MeilisearchOperator[T any] struct {
	client *meilisearch.Client
}

func NewMeillisearchOperator[T any](connection *MeilisearchConnection) *MeilisearchOperator[T] {
	return &MeilisearchOperator[T]{
		client: connection.client,
	}
}

func (c *MeilisearchOperator[T]) Create(
	ctx context.Context,
	collection *clerk.Collection,
	data T,
) error {
	index := c.client.Index(collection.Name)

	dataDocuments := []T{data}
	task, err := index.AddDocuments(dataDocuments)
	if err != nil {
		return err
	}

	waitTask, err := c.client.WaitForTask(task.TaskUID, meilisearch.WaitParams{
		Context: ctx,
	})
	if waitTask.Status == "failed" {
		return errors.New(waitTask.Error.Message)
	}
	return err
}

func (c *MeilisearchOperator[T]) Update(
	ctx context.Context,
	collection *clerk.Collection,
	filter map[string]any,
	data T,
	upsert bool,
) error {
	index := c.client.Index(collection.Name)

	filterable := []string{}
	for key := range filter {
		filterable = append(filterable, key)
	}
	if len(filterable) > 0 {
		if _, err := index.UpdateFilterableAttributes(&filterable); err != nil {
			return err
		}
	}

	if err := c.Create(ctx, collection, data); err != nil {
		return err
	}
	return nil
}

func (c *MeilisearchOperator[T]) Delete(
	ctx context.Context,
	collection *clerk.Collection,
	filter map[string]any,
) error {
	index := c.client.Index(collection.Name)

	id := ""
	for key, value := range filter {
		if strings.Contains(strings.ToLower(key), "id") {
			id = value.(string)
			break
		}
	}
	if len(id) == 0 {
		return ErrIdFilterIsRequired
	}

	task, err := index.DeleteDocument(id)
	if err != nil {
		return err
	}

	waitTask, err := c.client.WaitForTask(task.TaskUID, meilisearch.WaitParams{
		Context: ctx,
	})
	if waitTask.Status == "failed" {
		return errors.New(waitTask.Error.Message)
	}
	return err
}

func (c *MeilisearchOperator[T]) Query(
	ctx context.Context,
	collection *clerk.Collection,
	filter map[string]any,
	sorting map[string]bool,
	skip int,
	take int,
) (<-chan T, error) {
	// @todo: use sorting
	index := c.client.Index(collection.Name)

	request := &meilisearch.DocumentsQuery{}

	if skip > 0 {
		request.Offset = int64(skip)
	}
	if take > 0 {
		request.Limit = int64(take)
	}

	responded := make(chan struct{})
	responseChan := make(chan []T)

	go func() {
		defer close(responded)
		defer close(responseChan)

		var documents meilisearch.DocumentsResult

		if err := index.GetDocuments(request, &documents); err != nil {
			return
		}

		filteredDocuments := []map[string]any{}
		if len(filter) > 0 {
		loop:
			// @todo: implement pagination and let the search do the filtering
			for _, document := range documents.Results {
				for filterKey, filterValue := range filter {
					value, ok := document[filterKey]
					if !ok {
						continue loop
					}
					if value != filterValue {
						continue loop
					}
				}
				filteredDocuments = append(filteredDocuments, document)
			}
		} else {
			filteredDocuments = documents.Results
		}

		documentsAsJson, err := json.Marshal(filteredDocuments)
		if err != nil {
			return
		}

		var documentsAsT []T
		if err := json.Unmarshal(documentsAsJson, &documentsAsT); err != nil {
			return
		}

		responseChan <- documentsAsT
	}()

	out := make(chan T)

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return
			case data, ok := <-responseChan:
				if !ok {
					return
				}
				for _, d := range data {
					out <- d
				}
			}
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-responded:
		return out, nil
	}
}

func (c *MeilisearchOperator[T]) Search(
	ctx context.Context,
	collection *clerk.Collection,
	query string,
	highlight []string,
	filterable []string,
	filterQuery string,
	skip int,
	take int,
) (<-chan T, error) {
	index := c.client.Index(collection.Name)

	request := meilisearch.SearchRequest{}

	if len(highlight) > 0 {
		request.AttributesToHighlight = highlight
	}

	if len(filterable) > 0 {
		task, err := index.UpdateFilterableAttributes(&filterable)
		if err != nil {
			return nil, err
		}

		waitTask, err := c.client.WaitForTask(task.TaskUID, meilisearch.WaitParams{
			Context: ctx,
		})
		if err != nil {
			return nil, err
		}
		if waitTask.Status == "failed" {
			return nil, errors.New(waitTask.Error.Message)
		}
	}
	if len(filterQuery) > 0 {
		request.Filter = filterQuery
	}

	if skip >= 0 {
		request.Offset = int64(skip)
	}
	if take >= 0 {
		request.Limit = int64(take)
	}

	responded := make(chan struct{})
	responseChan := make(chan *meilisearch.SearchResponse)
	go func() {
		defer close(responded)
		defer close(responseChan)

		response, err := index.Search(query, &request)
		if err != nil {
			return
		}

		responseChan <- response
		responded <- struct{}{}
	}()

	out := make(chan T)

	go func() {
		defer close(out)
		select {
		case <-ctx.Done():
			return
		case response := <-responseChan:
			for _, hit := range response.Hits {
				var data T

				hitAsJson, err := json.Marshal(hit)
				if err != nil {
					continue
				}
				if err := json.Unmarshal(hitAsJson, &data); err != nil {
					continue
				}

				out <- data
			}
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-responded:
		return out, nil
	}
}
