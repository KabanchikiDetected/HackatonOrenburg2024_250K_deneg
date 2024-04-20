package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	BaseUser = "user"
	Deputy   = "deputy"
	Admin    = "admin"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Role     string             `bson:"role"`
	Password string             `bson:"password"`
}
