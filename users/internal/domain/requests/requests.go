package requests

type Register struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
