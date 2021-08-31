package database

import (
	"go_visitors_maintain_backend/src/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateVisit(visit models.Visit) (*mongo.InsertOneResult, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(visitCollection)
	result, err := collection.InsertOne(ctx, visit)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindAllVisit() ([]bson.M, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(visitCollection)

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

func FindUserVisit(userId string) ([]bson.M, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(visitCollection)

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

func UpdateVisit(visit models.Visit) (*mongo.UpdateResult, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(visitCollection)

	filter := bson.M{"visitid": visit.VisitId}

	update := bson.M{
		"$set": visit,
	}
	updateResult, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func DeleteVisit(visit models.Visit) (*mongo.DeleteResult, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(visitCollection)

	filter := bson.M{"userid": visit.UserId, "visitid": visit.VisitId}

	deleteResult, err := collection.DeleteMany(ctx, filter)

	if err != nil {
		return nil, err
	}
	return deleteResult, nil
}
