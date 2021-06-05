package models

//BrandSearchParams -> search Category Params
type BrandSearchParams struct {
	Keyword string `json:"keyword"`
}

//Category -> category model
type Brand struct {
	UintBase
	Name        string `json:"name"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}

// TableName  -> returns table name of model
func (b Brand) TableName() string {
	return "brands"
}
