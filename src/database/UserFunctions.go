package database

import (
	"go_visitors_maintain_backend/src/models"

	"go.mongodb.org/mongo-driver/bson"
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

	filter := bson.M{"username": username}

	var result models.User

	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return models.User{}, err
	}
	return result, nil
}
func FindUserFromID(userId string) (models.User, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(userCollection)

	filter := bson.M{"userid": userId}

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

	filter := bson.M{"userid": user.UserId}
	update := bson.M{
		"$set": user,
	}
	updateResult, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	return updateResult, nil
}
func FindAllUsers() ([]bson.M, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(userCollection)

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
func DeleteUser(userId string) (*mongo.DeleteResult, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(userCollection)
	filter := bson.M{"userid": userId}

	deleteResult, err := collection.DeleteMany(ctx, filter)

	if err != nil {
		return nil, err
	}
	return deleteResult, nil
}
