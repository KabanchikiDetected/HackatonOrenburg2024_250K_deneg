package requests

type CreateProduct struct {
	Name  string
	Image string // base64 encoded image
	Price int
}
