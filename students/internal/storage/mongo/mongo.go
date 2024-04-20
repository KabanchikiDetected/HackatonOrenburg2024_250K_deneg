package mongo

import (
	"context"

	"github.com/KabanchikiDetected/hackaton/students/internal/domain"
	"github.com/KabanchikiDetected/hackaton/students/internal/errors"
	"go.mongodb.org/mongo-driver/bson"
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

func (s *Storage) User(ctx context.Context, id string) (domain.User, error) {
	var user domain.User
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (s *Storage) Users(ctx context.Context) ([]domain.User, error) {
	var students []domain.User
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &students); err != nil {
		return nil, err
	}
	return students, nil
}

func (s *Storage) CreateUser(ctx context.Context, user domain.User) error {
	_, err := s.collection.InsertOne(ctx, user)
	return err
}

func (s *Storage) UpdateUser(ctx context.Context, id string, user domain.User) error {
	if _, err := s.User(ctx, id); err != nil {
		return errors.NotFound
	}
	_, err := s.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{
				Key: "$set", Value: bson.D{
					{Key: "first_name", Value: user.FirstName},
					{Key: "last_name", Value: user.LastName},
					{Key: "birthday", Value: user.Birthday},
					{Key: "description", Value: user.Description},
					{Key: "faculty_id", Value: user.FacultyID},
				},
			},
		},
	)
	return err
}

func (s *Storage) DeleteUser(ctx context.Context, id string) error {
	_, err := s.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (s *Storage) InsertImage(ctx context.Context, id string, image string) error {
	if _, err := s.User(ctx, id); err != nil {
		return errors.NotFound
	}
	_, err := s.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{
				Key: "$set", Value: bson.D{
					{Key: "photo", Value: image},
				},
			},
		},
	)
	return err
}
