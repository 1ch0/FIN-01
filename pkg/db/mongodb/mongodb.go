package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Requires the MongoDB Go Driver
	// https://go.mongodb.org/mongo-driver
	ctx := context.TODO()

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// ping 成功才代表连接成功
	err = client.Ping(context.Background(), nil)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Open an aggregation cursor
	coll := client.Database("mock").Collection("orders")
	_, err = coll.Aggregate(ctx, bson.A{
		bson.D{
			{"$match",
				bson.D{
					{"status", "completed"},
					{"orderDate",
						bson.D{
							{"$gte", time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)},
							{"$lt", time.Date(2019, 4, 1, 0, 0, 0, 0, time.UTC)},
						},
					},
				},
			},
		},
		bson.D{
			{"$group",
				bson.D{
					{"_id", primitive.Null{}},
					{"total", bson.D{{"$sum", "$total"}}},
					{"shippingFee", bson.D{{"$sum", "$shippingFee"}}},
					{"count", bson.D{{"$sum", 1}}},
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{
					{"gradTotal",
						bson.D{
							{"$add",
								bson.A{
									"$total",
									"$shippingFee",
								},
							},
						},
					},
					{"_id", 0},
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
