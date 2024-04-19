package models

const (
	BaseUser = "user"
	Deputy   = "deputy"
	Admin    = "admin"
)

// User model to store in storage
type User struct {
	ID       string
	Email    string `bson:"email"`
	Role     string `bson:"role"`
	Password string `bson:"password"`
}
