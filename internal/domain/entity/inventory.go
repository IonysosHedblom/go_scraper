package entity

type Inventory struct {
	Id     int64  `json:"inventory_id"`
	UserId string `json:"user_id"`
}

type Item struct {
	Id           int64  `json:"item_id"`
	Name         string `json:"name"`
	ImageUrl     string `json:"image_url"`
	ExpireDate   string `json:"expire_date"`
	VolumeAmount string `json:"volume_amount"`
	VolumeUnit   string `json:"volume_unit"`
	Category     string `json:"category"`
	Quantity     int    `json:"quantity"`
}

type InventoryItem struct {
	Id          int64 `json:"inventory_item_id"`
	ItemId      int64 `json:"item_id"`
	InventoryId int64 `json:"inventory_id"`
}
