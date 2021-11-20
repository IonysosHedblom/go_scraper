package entity

type Inventory struct {
	Id          int64           `json:"inventory_id"`
	Ingredients []InventoryItem `json:"ingredients"`
}

type InventoryItem struct {
	Id         int64  `json:"inventory_item_id"`
	Name       string `json:"name"`
	ImageUrl   string `json:"image_url"`
	ExpireDate string `json:"expire_date"`
	Volume     Volume `json:"volume"`
	Category   string `json:"category"`
	Quantity   int    `json:"quantity"`
}

type Volume struct {
	Amount int    `json:"amount"`
	Unit   string `json:"unit"`
}
