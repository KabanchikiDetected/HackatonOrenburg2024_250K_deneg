package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"name"`
	Image string             `bson:"image"`
	Price int                `bson:"price"`
}
