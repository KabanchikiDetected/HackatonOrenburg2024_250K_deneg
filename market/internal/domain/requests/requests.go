package requests

type CreateProduct struct {
	Name  string `json:"name"`
	Image string `json:"image"` // base64 encoded image
	Price int    `json:"price"`
}
