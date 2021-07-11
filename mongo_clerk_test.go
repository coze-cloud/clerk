package clerk

import (
	uuid "github.com/satori/go.uuid"
	"log"
	"testing"
)

type Message struct {
	Id string `bson:"_id"`
	Title string
	Body string
}

func TestMain(m *testing.M) {

	//ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	//defer cancel()
	//
	//clientOptions := options.Client().ApplyURI("mongodb://root:changeme@localhost:27017")
	//
	//client, err := mongo.Connect(ctx, clientOptions)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer func() {
	//	err := client.Disconnect(ctx)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	//
	//ctx, cancel = context.WithTimeout(context.Background(), 2 * time.Second)
	//defer cancel()
	//err = client.Ping(ctx, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//collection := client.Database("clerk").Collection("testing")
	//
	//messageId := uuid.NewV4().String()
	//message := Message{Id: messageId, Title: "Subject", Body: "Hello World"}
	//
	//ctx, cancel = context.WithTimeout(context.Background(), 5 * time.Second)
	//defer cancel()
	//_, err = collection.InsertOne(ctx, message)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//filter := bson.D{{"_id", messageId}}
	//ctx, cancel = context.WithTimeout(context.Background(), 5 * time.Second)
	//defer cancel()
	//
	//retrievedMessage := new(Message)
	//
	//err = collection.FindOne(ctx, filter).Decode(retrievedMessage)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println(retrievedMessage)

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
		Subject: "Hello",
		Body: "Hello World",
	})
	err = connection.SendCommand(createCommand)
	if err != nil {
		log.Fatal(err)
	}

	updateCommand := NewMongoUpdateCommand(collection, &Message{
		Id: messageId,
		Subject: "Hello",
		Body: "Hello Updated",
	}).Where("_id", messageId)
	err = connection.SendCommand(updateCommand)
	if err != nil {
		log.Fatal(err)
	}
}
