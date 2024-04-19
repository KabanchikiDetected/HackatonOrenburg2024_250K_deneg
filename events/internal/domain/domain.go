package domain

import "time"

type Event struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	FacultyID   string    `json:"faculty_id" bson:"faculty_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	StartDate   time.Time `json:"start_date" bson:"start_date"`
	EndDate     time.Time `json:"end_date" bson:"end_date"`
	IsFinished  bool      `json:"is_finished" bson:"is_finished"`
	Image       string    `json:"image" bson:"image"`
	Rating      float64   `json:"rating" bson:"rating"`
}

type EventsToStudent struct {
	Events    []Event `json:"events" bson:"events"`
	StudentID string  `json:"student_id" bson:"student_id"`
}
