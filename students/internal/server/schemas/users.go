package schemas

import (
	"time"

	"github.com/KabanchikiDetected/hackaton/students/internal/domain"
)

const layout = "2006.01.02"

type UserSchema struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Description string `json:"description"`
	Birthday    string `json:"birthday"`
	FacultyID   string `json:"faculty_id"`
}

func (s *UserSchema) ToDomain() (domain.User, error) {
	birthday, err := time.Parse(layout, s.Birthday)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		FirstName:   s.FirstName,
		LastName:    s.LastName,
		Description: s.Description,
		Birthday:    birthday,
		FacultyID:   s.FacultyID,
	}, nil
}
