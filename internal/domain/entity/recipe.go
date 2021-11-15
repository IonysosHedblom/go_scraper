package entity

type Recipe struct {
	Id                 int64    `json:"recipe_id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	ImageUrl           string   `json:"imageurl"`
	Ingredients        []string `json:"ingredients"`
	QueryId            *int64   `json:"query_id"`
	IngredientSearchId *int64   `json:"ingredient_search_id"`
}
