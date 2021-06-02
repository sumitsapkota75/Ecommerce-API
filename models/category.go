package models

//CategorySearchParams -> search Category Params
type CategorySearchParams struct {
	Keyword string `json:"keyword"`
}

//Category -> category model
type Category struct {
	UintBase
	Name        string `json:"name"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
}

// TableName  -> returns table name of model
func (c Category) TableName() string {
	return "categories"
}
