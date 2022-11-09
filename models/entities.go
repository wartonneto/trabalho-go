package models

type Livro struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
}
