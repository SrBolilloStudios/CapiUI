package models

type Product struct {
	ID          int      `json:"id"`
	SKU         string   `json:"sku"`
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Description string   `json:"description"`
	Img         string   `json:"img"`
	Category    Category `json:"category"`
	Stock       int      `json:"stock"`
}
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
