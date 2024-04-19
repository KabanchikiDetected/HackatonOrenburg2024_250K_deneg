package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

func convertStringToObjectID(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}
