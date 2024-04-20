package responses

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"` // url
	Price int    `json:"price"`
}
