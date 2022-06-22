# clerk

📒 A minimalistic library for abstracting database operations

## Installation

Adding _clerk_ to your Go module is as easy as calling this command in your project

```shell
go get github.com/coze-cloud/clerk
```

## Supported databases

_clerk_ has builtin support for the following database/search engines:

- [MongoDB](https://www.mongodb.com/) - MongoDB is a document-oriented database
- [Meilisearch](https://www.meilisearch.com/) - Meilisearch is a powerful and fast search engine

## Usage

Being a minimalistic library, _clerk_ only provides the basics. The rest is up to your specific need.

### Creating a connection

```go
connection, err := mongodb.NewMongoConnection("mongodb://root:root@localhost:27017")
if err != nil {
	panic(err)
}

defer connection.Close(func(err error) {
	if err != nil {
		panic(err)
	}
})
```

### Defining a database operator instance

```go
operator := mongodb.NewMongodbOperator[T](connection)
```

The generic parameter T defines the data type which the operator can interact with.
An operator has to be defined for each data type in use with _clerk_.

### Defining a database & collection

```go
collection := clerk.NewDatabase("foo").Collection("bar")
```

Certain operators only work with collections and don't need a database:

```go
collection := clerk.NewCollection("foo")
```

### Persisting a data in a collection

```go
type Message struct {
    Id   string `bson:"_id"`
    Body string
}

createCtx, createCancel := context.WithTimeout(
    context.Background(),
    time.Second * 5,
)
defer createCancel()

create := clerk.NewCreate(collection, Message{
    Id:   "0",
    Body: "Hello World",
})
if err := create.Execute(createCtx, operator); err != nil {
    panic(err)
}
```

### Querying the collection

```go
type Message struct {
    Id   string `bson:"_id"`
    Body string
}

results := []Message{}

queryCtx, queryCancel := context.WithTimeout(
    context.Background(),
    time.Second * 5,
)
defer queryCancel()

query := clerk.NewQuery[Message](collection).Where("_id", "0")
queryChan, err := query.Execute(queryCtx, operator)
if err != nil {
    panic(err)
}

for result := range queryChan {
    results := append(results, result)
}

fmt.Println(results...)
```

---

Copyright © 2022 - The cozy team **& contributors**
