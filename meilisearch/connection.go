package meilisearch

import (
	"context"

	"github.com/meilisearch/meilisearch-go"
)

type MeilisearchConnection struct {
	ctx    context.Context
	client *meilisearch.Client
}

func NewMeillisearchConnection(ctx context.Context, url string, apiKey string) (*MeilisearchConnection, error) {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   url,
		APIKey: apiKey,
	})

	return &MeilisearchConnection{
		ctx:    ctx,
		client: client,
	}, nil
}

func (c *MeilisearchConnection) Close(handler func(err error)) {
	handler(nil)
}
