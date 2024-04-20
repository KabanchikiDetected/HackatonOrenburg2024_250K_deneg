package users

import (
	"context"

	"github.com/KabanchikiDetected/hackaton/events/internal/domain"
	mongoStorage "github.com/KabanchikiDetected/hackaton/events/internal/storage/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	usersCollection  *mongo.Collection
	eventsCollection *mongo.Collection
}

func New(
	usersColection *mongo.Collection,
	eventsCollection *mongo.Collection,
) *Storage {
	return &Storage{
		usersCollection:  usersColection,
		eventsCollection: eventsCollection,
	}
}

func (s *Storage) UserEvents(ctx context.Context, id string) (domain.EventsToStudent, error) {
	const op = "mongo.UserEvents"

	objectID, err := mongoStorage.ConvertStringToObjectID(id)
	if err != nil {
		return domain.EventsToStudent{}, err
	}
	var eventsToStudent domain.EventsToStudent
	err = s.usersCollection.FindOne(ctx, bson.M{"user_id": objectID}).Decode(&eventsToStudent)
	if err != nil {
		return domain.EventsToStudent{}, err
	}
	return eventsToStudent, nil
}

func (s *Storage) AddEventToUser(ctx context.Context, studentID string, eventID string) error {
	objectID, err := mongoStorage.ConvertStringToObjectID(studentID)
	if err != nil {
		return err
	}
	eventObjectID, err := mongoStorage.ConvertStringToObjectID(eventID)
	if err != nil {
		return err
	}

	var event domain.Event
	if err = s.eventsCollection.FindOne(ctx, bson.M{"_id": eventObjectID}).Decode(&event); err != nil {
		return err
	}

	_, err = s.usersCollection.UpdateOne(
		ctx,
		bson.M{"user_id": objectID},
		bson.D{
			{
				Key: "$push", Value: bson.D{
					{Key: "events", Value: event},
				},
			},
			{
				Key: "$inc", Value: bson.D{
					{Key: "rating", Value: event.Rating},
				},
			},
		},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return err
	}
	return nil
}
