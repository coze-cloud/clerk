package main

import (
	"context"
	"fmt"

	clerk "github.com/coze-cloud/clerk/src"
	"github.com/coze-cloud/clerk/src/meilisearch"
)

func main() {
	conn, err := meilisearch.NewMeillisearchConnection(
		context.Background(),
		"http://localhost:7700",
		"KvcfjKliQSSJC5X5OWEL",
	)
	if err != nil {
		panic(err)
	}

	type Message struct {
		Id   string `json:"id"`
		Text string `json:"text"`
	}

	operator := meilisearch.NewMeillisearchOperator[Message](conn)
	messages := clerk.NewCollection("messages")

	create := clerk.NewCreate(messages, Message{
		Id:   "1",
		Text: "Hello world!",
	})
	if err := create.Execute(context.Background(), operator); err != nil {
		panic(err)
	}

	query := clerk.NewQuery[Message](messages).
		Where("text", "Hello, world!").
		Where("id", "1655882087")
	queryStream, err := query.Execute(context.Background(), operator)
	if err != nil {
		panic(err)
	}
	for message := range queryStream {
		fmt.Println(message)
	}

	// TODO: Query - retrieves data from the index
	// TODO: Update / use as Index update ...?
}
