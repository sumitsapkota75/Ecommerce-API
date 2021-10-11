package models

// Vendor -> Model

type Vendor struct {
	UserBase
	Name         string `json:"name"`
	StoreName    string `json:"store_name"`
	DocumentType string `json:"document_type"`
	DocumentID   string `json:"document_id"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	Thumbnail    string `json:"thumbnail"`
}

// TableName  -> returns table name of model
func (u Vendor) TableName() string {
	return "vendors"
}
