package request

type (
	Module struct {
		Title       string  `json:"title" form:"title" validate:"required,min=3"`
		Description *string `json:"description" form:"description"`
	}
)
