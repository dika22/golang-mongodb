package repository

import (
	"context"
	"dataon-test/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepositoryImpl interface {
	InsertOrder(ctx context.Context, payload model.Orders) error
	FindByID(ctx context.Context, id string) (*model.Orders, error)
	List(ctx context.Context) ([]*model.Orders, error)
	Update(ctx context.Context, id string, payload model.OrderUpdate) error
}

type orderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(mongodb *mongo.Database) *orderRepository {
	return &orderRepository{
		collection: mongodb.Collection("orders"),
	}
} 


func (r orderRepository) InsertOrder(ctx context.Context, payload model.Orders) error  {

	res, err := r.collection.InsertOne(ctx, payload)
	if err != nil {
		return err
	}
	res.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (r orderRepository) FindByID(ctx context.Context, id string) (*model.Orders, error) {
	order := &model.Orders{}
	Id, _ := primitive.ObjectIDFromHex(id)
	err := r.collection.FindOne(ctx, bson.M{"_id": Id}).Decode(order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("order not found")
			return nil, err
		}
		log.Println("error find by id", err)
		return nil, err
	}

	return order, nil
}

func (r orderRepository) List(ctx context.Context) ([]*model.Orders, error) {
	var results []*model.Orders

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r orderRepository) Update(ctx context.Context, id string, payload model.OrderUpdate) error {
	Id, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{"$set": payload}

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": Id}, update)
	if err != nil {
		log.Println("error update office", err)
		return err
	}
	return nil


}