package domain

// declare all book in system
type Book struct {
	Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Category_id UUID   `json:"category_id"`
}
