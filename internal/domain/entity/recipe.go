package entity

type Recipe struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ImageUrl    string   `json:"image_url"`
	Ingredients []string `json:"ingredients"`
}
