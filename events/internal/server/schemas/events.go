package schemas

import (
	"time"

	"github.com/KabanchikiDetected/hackaton/events/internal/domain"
)

const layout = "2006-01-02 15:04"

type EventSchema struct {
	FacultyID   string `json:"faculty_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	IsFinished  bool   `json:"is_finished"`
	Rating      int    `json:"rating"`
}

func (s *EventSchema) ToDomain() (domain.Event, error) {
	startDate, err := time.Parse(layout, s.StartDate)
	if err != nil {
		return domain.Event{}, err
	}

	endDate, err := time.Parse(layout, s.EndDate)
	if err != nil {
		return domain.Event{}, err
	}
	return domain.Event{
		FacultyID:   s.FacultyID,
		Title:       s.Title,
		Description: s.Description,
		StartDate:   startDate,
		EndDate:     endDate,
		IsFinished:  s.IsFinished,
		Rating:      float64(s.Rating),
	}, nil
}
