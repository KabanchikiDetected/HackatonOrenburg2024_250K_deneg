package responses

type Product struct {
	Name  string `json:"name"`
	Image string `json:"image"` // url
	Price int    `json:"price"`
}
