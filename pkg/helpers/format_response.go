package helpers

import "github.com/gofiber/fiber/v2"

type (
	Meta struct {
		Status  string `json:"status"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	BaseResponse struct {
		Meta       Meta        `json:"meta"`
		Data       any         `json:"data"`
		Errors     any         `json:"errors,omitempty"`
		Pagination *Pagination `json:"pagination,omitempty"`
	}
)

// 200 - OK
func SuccessOK(message string, data any, pagination ...Pagination) BaseResponse {
	var p *Pagination

	if len(pagination) > 0 {
		p = &pagination[0]
	}

	return BaseResponse{
		Meta: Meta{
			Status:  "ok",
			Code:    fiber.StatusOK,
			Message: message,
		},
		Data:       data,
		Pagination: p,
	}
}

// 201 - Created
func SuccessCreated(message string, data any) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Status:  "created",
			Code:    fiber.StatusCreated,
			Message: message,
		},
		Data: data,
	}
}

// 400 - Bad Request
func BadRequest(message string, errors any) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Status:  "bad request",
			Code:    fiber.StatusBadRequest,
			Message: message,
		},
		Errors: errors,
	}
}

// 401 - Unauthorized
func Unauthorized(message string) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Status:  "unauthorized",
			Code:    fiber.StatusUnauthorized,
			Message: message,
		},
	}
}

// 403 - Forbidden
func Forbidden(message string) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Status:  "forbidden",
			Code:    fiber.StatusForbidden,
			Message: message,
		},
	}
}

// 404 - Not Found
func NotFound(message string) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Status:  "not found",
			Code:    fiber.StatusNotFound,
			Message: message,
		},
	}
}

// 409 - Conflict
func Conflict(message string) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Status:  "conflict",
			Code:    fiber.StatusConflict,
			Message: message,
		},
	}
}

// 422 - Unprocessable Entity
func UnprocessableEntity(message string) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Status:  "unprocessable entity",
			Code:    fiber.StatusUnprocessableEntity,
			Message: message,
		},
	}
}

// 500 - Internal Server Error
func InternalServerError(message string) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Status:  "internal server error",
			Code:    fiber.StatusInternalServerError,
			Message: message,
		},
	}
}
