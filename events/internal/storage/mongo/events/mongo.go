package mongo

import (
	"context"

	"github.com/KabanchikiDetected/hackaton/events/internal/domain"
	"github.com/KabanchikiDetected/hackaton/events/internal/errors"
	mongoStorage "github.com/KabanchikiDetected/hackaton/events/internal/storage/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	collection *mongo.Collection
}

func New(collection *mongo.Collection) *Storage {
	return &Storage{
		collection: collection,
	}
}

func (s *Storage) Event(ctx context.Context, id string) (domain.Event, error) {
	const op = "mongo.Event"

	objectID, err := mongoStorage.ConvertStringToObjectID(id)
	if err != nil {
		return domain.Event{}, err
	}

	var event domain.Event

	if err := s.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&event); err != nil {
		return domain.Event{}, err
	}
	return event, nil
}

func (s *Storage) Events(ctx context.Context, isFinished bool) ([]domain.Event, error) {
	const op = "mongo.Events"
	filter := bson.M{}
	if isFinished {
		filter["is_finished"] = isFinished
	}

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var events []domain.Event
	if err = cursor.All(ctx, &events); err != nil {
		return nil, err
	}
	return events, nil
}

func (s *Storage) AddEvent(ctx context.Context, event domain.Event) (string, error) {
	const op = "mongo.AddEvent"
	result, err := s.collection.InsertOne(ctx, event)
	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

func (s *Storage) UpdateEvent(ctx context.Context, id string, event domain.Event) error {
	const op = "mongo.UpdateEvent"

	objectID, err := mongoStorage.ConvertStringToObjectID(id)

	if err != nil {
		return err
	}

	if _, err := s.Event(ctx, id); err != nil {
		return errors.NotFound
	}

	_, err = s.collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.D{
			{
				Key: "$set", Value: bson.D{
					{Key: "title", Value: event.Title},
					{Key: "description", Value: event.Description},
					{Key: "start_date", Value: event.StartDate},
					{Key: "end_date", Value: event.EndDate},
					{Key: "is_finished", Value: event.IsFinished},
					{Key: "rating", Value: event.Rating},
					{Key: "faculty_id", Value: event.FacultyID},
				},
			},
		},
	)
	return err
}

func (s *Storage) DeleteEvent(ctx context.Context, id string) error {
	const op = "mongo.DeleteEvent"

	objectID, err := mongoStorage.ConvertStringToObjectID(id)
	if err != nil {
		return err
	}
	_, err = s.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

func (s *Storage) InsertImage(ctx context.Context, id string, image string) error {
	const op = "mongo.InsertImage"

	objectID, err := mongoStorage.ConvertStringToObjectID(id)
	if err != nil {
		return err
	}

	if _, err := s.Event(ctx, id); err != nil {
		return errors.NotFound
	}

	_, err = s.collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.D{
			{
				Key: "$set", Value: bson.D{
					{Key: "image", Value: image},
				},
			},
		},
	)
	return err
}
