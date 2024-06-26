package users

import (
	"context"
	"sort"

	"github.com/KabanchikiDetected/hackaton/events/internal/domain"
	"github.com/KabanchikiDetected/hackaton/events/internal/errors"
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

func (s *Storage) DicrementRating(ctx context.Context, id string, rating int) error {
	objectID, err := mongoStorage.ConvertStringToObjectID(id)
	if err != nil {
		return err
	}
	var user domain.EventsToStudent
	if err = s.usersCollection.FindOne(ctx, bson.M{"user_id": objectID}).Decode(&user); err != nil {
		return err
	}
	if user.Rating < rating {
		return errors.BadRequest
	}

	_, err = s.usersCollection.UpdateOne(
		ctx,
		bson.M{"user_id": objectID},
		bson.D{
			{
				Key: "$inc", Value: bson.D{
					{Key: "rating", Value: -rating},
				},
			},
		},
	)
	return err
}

func (s *Storage) GetAllUserRatings(ctx context.Context) ([]domain.UserRating, error) {
	cursor, err := s.usersCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var eventsToStudents []domain.EventsToStudent
	if err = cursor.All(ctx, &eventsToStudents); err != nil {
		return nil, err
	}
	var userRatings []domain.UserRating
	for _, eventsToStudent := range eventsToStudents {
		userRatings = append(userRatings, domain.UserRating{
			UserID: eventsToStudent.UserID,
			Rating: eventsToStudent.Rating,
		})
	}
	// sort this slice by rating
	sort.Slice(userRatings, func(i, j int) bool {
		return userRatings[i].Rating > userRatings[j].Rating
	})
	return userRatings, nil
}
