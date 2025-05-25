package adapters

import (
	"context"
	"time"

	"go-hex/core"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoOrderRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewMongoOrderRepository(db *mongo.Database) core.OrderRepository {
	return &MongoOrderRepository{
		db:         db,
		collection: db.Collection("orders"),
	}
}

func (r *MongoOrderRepository) Save(order core.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// สร้าง document ที่แมพจาก core.Order
	doc := bson.M{
		"order_name": order.Order_name, // แปลง uint → string
		"total":      order.Total,
	}

	_, err := r.collection.InsertOne(ctx, doc)
	return err
}
