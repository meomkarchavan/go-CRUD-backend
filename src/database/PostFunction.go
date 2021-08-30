package database

import (
	"blog_rest_api_gin/src/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePost(post models.Post) (*mongo.InsertOneResult, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}
	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(postCollection)
	result, err := collection.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindAllPost() ([]bson.M, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(postCollection)

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

func FindUserPost(userId string) ([]bson.M, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(postCollection)

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

func UpdatePost(post models.Post) (*mongo.UpdateResult, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(postCollection)

	filter := bson.M{"userid": post.UserId, "postid": post.PostId}

	update := bson.M{
		"$set": post,
	}
	updateResult, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func DeletePost(post models.Post) (*mongo.DeleteResult, error) {
	client, ctx, cancel, err := createConnection(url)
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	collection := client.Database(db).Collection(postCollection)

	filter := bson.M{"userid": post.UserId, "postid": post.PostId}

	deleteResult, err := collection.DeleteMany(ctx, filter)

	if err != nil {
		return nil, err
	}
	return deleteResult, nil
}
