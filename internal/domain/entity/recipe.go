package entity

type Recipe struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ImageUrl    string   `json:"imageurl"`
	Ingredients []string `json:"ingredients"`
}
