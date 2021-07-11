package clerk

import (
	uuid "github.com/satori/go.uuid"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	type Message struct {
		Id string `bson:"_id"`
		Subject string
		Body string
	}

	connection, err := NewMongoConnection("mongodb://root:changeme@localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close(func(err error) {
		log.Fatal(err)
	})

	collection := NewDatabase("clerk").GetCollection("test-collection")

	messageId := uuid.NewV4().String()

	createCommand := NewMongoCreateCommand(collection, &Message{
		Id: messageId,
		Subject: "Hello!",
		Body: "Hello World",
	})
	err = connection.SendCommand(createCommand)
	if err != nil {
		log.Fatal(err)
	}

	singleQuery := NewMongoSingleQuery(collection).Where("_id", messageId)
	iterator, err := connection.SendQuery(singleQuery)
	if err != nil {
		log.Fatal(err)
	}

	retrievedMessage := Message{}
	err = iterator.Single(&retrievedMessage)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(retrievedMessage)

	updateCommand := NewMongoUpdateCommand(collection, &Message{
		Id: messageId,
		Subject: "Hello!",
		Body: "Hello Updated",
	}).Where("_id", messageId)
	err = connection.SendCommand(updateCommand)
	if err != nil {
		log.Fatal(err)
	}

	listQuery := NewMongoListQuery(collection)
	iterator, err = connection.SendQuery(listQuery)
	if err != nil {
		log.Fatal(err)
	}

	retrievedMessages := []Message{}
	for iterator.Next() {
		retrievedMessage := Message{}
		err := iterator.Decode(&retrievedMessage)
		if err != nil {
			log.Fatal(err)
		}
		retrievedMessages = append(retrievedMessages, retrievedMessage)

		deleteCommand := NewMongoDeleteCommand(collection).Where("_id", retrievedMessage.Id)
		err = connection.SendCommand(deleteCommand)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println(retrievedMessages)
}
