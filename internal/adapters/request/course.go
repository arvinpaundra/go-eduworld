package request

import "mime/multipart"

type (
	QueryFindCourses struct {
		Keyword  string
		Category string
		Page     int
	}

	Course struct {
		CategoryId  string                `json:"category_id" form:"category_id" validate:"required"`
		UserId      string                `json:"user_id" form:"user_id" validate:"required"`
		InterestId  string                `json:"interest_id" form:"interest_id" validate:"required"`
		Title       string                `json:"title" form:"title" validate:"required,min=5"`
		IsPublished bool                  `json:"is_published" form:"is_published" validate:"required,boolean"`
		Level       string                `json:"level" form:"level" validate:"required"`
		Type        string                `json:"type" form:"type" validate:"required"`
		Price       *int                  `json:"price" form:"price" validate:"required_if=Type premium"`
		Description *string               `json:"description" form:"description"`
		Thumbnail   *multipart.FileHeader `json:"thumbnail" form:"thumbnail"`
	}
)
