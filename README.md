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
connection, err := NewMongoConnection("mongodb://root:root@localhost:27017")
if err != nil {
    log.Fatal(err)
}

defer connection.Close(func(err error) {
    log.Fatal(err)
})
```

### Defining a database & collection

```go
collection := NewDatabase("clerk").GetCollection("messages")
```

### Persisting a data structure in a collection

```go
type Message struct {
    Id string `bson:"_id"`
    Body string
}

createCommand := NewMongoCreateCommand(collection, &Message{
    Id: uuid.NewV4().String(),
    Body: "Hello World",
})

if err = connection.SendCommand(createCommand); err != nil {
    log.Fatal(err)
}
```

### Querying the collection for a data structure

```go
type Message struct {
    Id string `bson:"_id"`
    Body string
}

singleQuery := NewMongoSingleQuery(collection).Where("_id", uuid.NewV4().String())

iterator, err := connection.SendQuery(singleQuery)
if err != nil {
    log.Fatal(err)
}

message := Message{}
if err = iterator.Single(&message); err != nil {
    log.Fatal(err)
}
```


```go
listQuery := NewMongoListQuery(collection)

iterator, err = connection.SendQuery(listQuery)
if err != nil {
    log.Fatal(err)
}

messages := []Message{}
for iterator.Next() {
    message := Message{}
    if err := iterator.Decode(&message); err != nil {
        log.Fatal(err)
    }
    messages = append(messages, message)
}
```

## Future plans

* [x] Unit tests for the existing components
* [ ] Support for more database implementations

---

Copyright © 2021 - The cozy team **& contributors**