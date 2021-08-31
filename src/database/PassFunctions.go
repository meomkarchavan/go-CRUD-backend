package database

import (
	"go_visitors_maintain_backend/src/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePass(pass models.Pass) (*mongo.InsertOneResult, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(passCollection)
	result, err := collection.InsertOne(ctx, pass)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindAllPass() ([]bson.M, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(passCollection)

	var result []bson.M
	var cursor *mongo.Cursor
	cursor, err = collection.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func FindUserPass(userId string) ([]bson.M, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(passCollection)

	filter := bson.M{"userid": userId}

	var result []bson.M
	var cursor *mongo.Cursor
	cursor, err = collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &result); err != nil {
		log.Fatal(err)
	}
	return result, nil
}

func UpdatePass(pass models.Pass) (*mongo.UpdateResult, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(passCollection)

	filter := bson.M{"passid": pass.PassId}

	update := bson.M{
		"$set": pass,
	}
	updateResult, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func DeletePass(pass models.Pass) (*mongo.DeleteResult, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(passCollection)

	filter := bson.M{"userid": pass.UserId, "passid": pass.PassId}

	deleteResult, err := collection.DeleteMany(ctx, filter)

	if err != nil {
		return nil, err
	}
	return deleteResult, nil
}
