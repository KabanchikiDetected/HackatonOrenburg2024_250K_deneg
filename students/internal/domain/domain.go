package domain

import "time"

type Role string

var (
	Student Role = "student"
	Deputy  Role = "deputy"
)

type User struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	FirstName   string    `json:"name" bson:"name"`
	LastName    string    `json:"last_name" bson:"last_name"`
	Description string    `json:"description" bson:"description"`
	Birthday    time.Time `json:"birthday" bson:"birthday"`
	FacultyID   string    `json:"faculty_id" bson:"faculty_id"`
	Photo       string    `json:"photo" bson:"photo"`
	Role        Role      `json:"role" bson:"role"`
}
