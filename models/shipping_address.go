package models

type ShippingAddress struct {
	UintBase
	UserID   string `json:"user_id"`
	User     `json:"user"`
	District string `json:"district"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

func (s ShippingAddress) TableName() string {
	return "shipping_addresses"
}
