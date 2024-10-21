package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	// MONGO_URL               = "mongodb://localhost:27017"
	// MONGO_DB                = "card_game_db_1"
	MONGO_URL               = "mongodb+srv://dev:e7xiosdK0aCY7bYx@dev-trnbo.mongodb.net/"
	MONGO_DB                = "card_game_db_3"
	MONGO_COLLECTION_PLAYER = "player"
	MONGO_COLLECTION_CLAN   = "clan"
)

type mongoDBOption struct {
	dialTimeout time.Duration
	url         string
	readPref    *readpref.ReadPref
	bsonOptions *options.BSONOptions
}

var (
	opt    *mongoDBOption
	Client *mongo.Client
)

func init() {
	opt := &mongoDBOption{
		dialTimeout: 3 * time.Second,
		url:         MONGO_URL,
		readPref:    readpref.Primary(),
		bsonOptions: &options.BSONOptions{
			UseJSONStructTags: true,
			NilMapAsEmpty:     true,
			NilSliceAsEmpty:   true,
		},
	}

	ctx, canceler := context.WithTimeout(context.Background(), time.Second*3)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(opt.url).SetReadPreference(opt.readPref).SetBSONOptions(opt.bsonOptions))
	canceler()

	if err != nil {
		panic(err)
	}

	Client = client
}

func main() {
	filter := bson.D{}
	results := make([]map[string]any, 0, 64)
	cursor, err := Client.Database(MONGO_DB).Collection(MONGO_COLLECTION_CLAN).Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		result := make(map[string]any)
		err = cursor.Decode(result)
		if err != nil {
			panic(err)
		}
		results = append(results, result)
	}
	fmt.Printf("results = %v\n", len(results))
	for _, result := range results {
		fmt.Printf("result = %v\n", result)
	}
}
