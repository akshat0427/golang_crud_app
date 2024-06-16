package app

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMongoClient() (*mongo.Client, context.Context, context.CancelFunc, error) {
	clientopts := options.Client().ApplyURI("mongodb://localhost:27017/")
	ctx, cancel := context.WithCancel(context.Background())
	client, err := mongo.Connect(ctx, clientopts)
	return client, ctx, cancel, err
}

type N struct {
	Name string
	Data string
}

func St() string {

	return "this works fine"

}

func Ins(n1 string, s string) {

	client, ctx, cancel, err := getMongoClient()

	if err != nil {
		log.Fatal(err)

	}

	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)

		}
	}()

	collection := client.Database("test").Collection("crud")

	inf := N{n1, s}

	collection.InsertOne(context.Background(), inf)

}

func Find(s2 string) string {
	client, ctx, cancel, err := getMongoClient()

	if err != nil {
		log.Fatal(err)

	}

	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)

		}
	}()

	collection2 := client.Database("test").Collection("crud")
	objectID, err := primitive.ObjectIDFromHex(s2)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": objectID}
	var res map[string]string
	collection2.FindOne(context.Background(), filter).Decode(&res)

	return res["name"]

}

func Update(f string, u string) {
	client, ctx, cancel, err := getMongoClient()

	if err != nil {
		log.Fatal(err)

	}

	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)

		}
	}()

	collection3 := client.Database("test").Collection("crud")
	objectID, err := primitive.ObjectIDFromHex(f)

	if err != nil {
		log.Fatal(err)
	}

	filter1 := bson.M{"_id": objectID}

	update := bson.M{"$set": bson.M{"name": u}}

	_, err = collection3.UpdateOne(context.Background(), filter1, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("updated")
	// fmt.Println(result.MatchedCount, result.ModifiedCount)

}

func DeleteCollection(s5 string) {
	client, ctx, cancel, err := getMongoClient()

	if err != nil {
		log.Fatal(err)

	}

	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)

		}
	}()

	err = client.Database("test").Collection(s5).Drop(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("database dropped")

}

func DropDocument(arg string) {

	client, ctx, cancel, err := getMongoClient()

	if err != nil {
		log.Fatal(err)

	}

	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)

		}
	}()

	collection2 := client.Database("test").Collection("crud")

	objectID, err := primitive.ObjectIDFromHex(arg)

	if err != nil {
		log.Fatal(err)
	}

	filter2 := bson.M{"_id": objectID}

	collection2.DeleteOne(context.Background(), filter2)

	print("dropped the document")

	// fmt.Println(res2["data"])

}

type Item struct {
	Name string
	Data string
}

func PrintCollection() ([]string, []string) {

	client, ctx, cancel, err := getMongoClient()

	if err != nil {
		log.Fatal(err)

	}

	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)

		}
	}()

	collection := client.Database("test").Collection("crud")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	name := []string{}
	id_ := []string{}

	for cursor.Next(context.Background()) {
		var result map[string]string
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}

		name = append(name, result["name"])
		id_ = append(id_, result["_id"])

	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return name, id_

}
