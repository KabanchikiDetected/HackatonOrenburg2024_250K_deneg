package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

func ConvertStringToObjectID(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}
