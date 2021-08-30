package database

import (
	"blog_rest_api_gin/src/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateTodo(todo models.Todo) (*mongo.InsertOneResult, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(todoCollection)
	result, err := collection.InsertOne(ctx, todo)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindAllTodo() ([]bson.M, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(todoCollection)

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

func FindUserTodo(userId string) ([]bson.M, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(todoCollection)

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

func UpdateTodo(todo models.Todo) (*mongo.UpdateResult, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(todoCollection)

	filter := bson.M{"todoid": todo.TodoId}

	update := bson.M{
		"$set": todo,
	}
	updateResult, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func DeleteTodo(todo models.Todo) (*mongo.DeleteResult, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(todoCollection)

	filter := bson.M{"userid": todo.UserId, "todoid": todo.TodoId}

	deleteResult, err := collection.DeleteMany(ctx, filter)

	if err != nil {
		return nil, err
	}
	return deleteResult, nil
}
