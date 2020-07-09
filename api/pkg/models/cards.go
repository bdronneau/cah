package models

// Cards of app
type Cards struct {
	ID          uint64 `json:"id" form:"id" query:"id"`
	Description string `json:"description" form:"description" query:"description"`
}
