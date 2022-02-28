# clerk
📒 A minimalistic library for abstracting database operations

## Installation

Adding *clerk* to your Go module is as easy as calling this command in your project

```shell
go get github.com/cozy-hosting/clerk
```

## Usage

Being a minimalistic library, *clerk* only provides the basics. The rest is up to your specific need.

### Creating a connection

```go
connection, err := mongodb.NewMongoConnection("mongodb://root:change-me@host.docker.internal:27017")
if err != nil {
	panic(err)
}

defer connection.Close(func(err error) {
	if err != nil {
		panic(err)
	}
})
```

### Defining a database & collection

```go
collection := clerk.NewDatabase("foo").Collection("bar")
```

### Persisting a data in a collection

```go
type Message struct {
    Id   string `bson:"_id"`
    Body string
}

create := clerk.NewCreate(collection, Message{
    Id:   "0",
    Body: "Hello World",
})
if err := create.Execute(connection.Context()); err != nil {
    panic(err)
}
```

### Querying the collection

```go
query := clerk.NewQuery(collection).Where("_id", "0")
results, err := query.Execute(connection.Context())
if err != nil {
    panic(err)
}

fmt.Println(results...)
```

---

Copyright © 2021 - The cozy team **& contributors**