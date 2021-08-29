package database

import (
	"blog_rest_api_gin/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user models.User) (*mongo.InsertOneResult, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(userCollection)
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func FindUser(username string) (models.User, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(userCollection)

	filter := bson.D{primitive.E{Key: "username", Value: username}}

	var result models.User

	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return models.User{}, err
	}
	return result, nil
}
func FindUserId(username string) (models.User, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(userCollection)

	filter := bson.M{"username": username}

	var result models.User

	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return models.User{}, err
	}
	return result, nil
}

func UpdateUser(user models.User) (*mongo.UpdateResult, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(userCollection)

	filter := bson.D{primitive.E{Key: "username", Value: user.Username}}
	update := bson.M{
		"$set": user,
	}
	updateResult, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func DeleteUser(username string) (*mongo.DeleteResult, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(userCollection)

	filter := bson.D{primitive.E{Key: "username", Value: username}}

	deleteResult, err := collection.DeleteMany(ctx, filter)

	if err != nil {
		return nil, err
	}
	return deleteResult, nil
}
