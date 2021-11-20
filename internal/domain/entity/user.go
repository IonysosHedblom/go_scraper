package entity

type User struct {
	Id          string `json:"user_id"`
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	InventoryId string `json:"inventory_id"`
}
