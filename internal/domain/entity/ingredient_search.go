package entity

type IngredientSearch struct {
	Id          int      `json:"id"`
	Ingredients []string `json:"query"`
}
