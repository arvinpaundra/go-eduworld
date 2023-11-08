package helpers

import (
	"encoding/base64"
	"math"
)

type (
	Pagination struct {
		*OffsetPagination `json:",omitempty"`
		*CursorPagination `json:",omitempty"`
	}

	OffsetPagination struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		TotalData  int `json:"total_data"`
		TotalPages int `json:"total_pages"`
	}

	CursorPagination struct {
		Previous *string `json:"previous"`
		Next     *string `json:"next"`
	}
)

func getLimit(limit int) int {
	if limit < 1 {
		return 1
	}

	return limit
}

func getPage(page int) int {
	page += 1

	if page < 1 {
		return 1
	}

	return page
}

func getTotalPages(totalData, limit int) int {
	total := math.Ceil(float64(totalData) / float64(limit))

	if total < 1 {
		return 1
	}

	return int(total)
}

// NewOffsetPagination common pagination technique using limit and offset
func NewOffsetPagination(page, perPage, totalData int) *OffsetPagination {
	return &OffsetPagination{
		Page:       getPage(page),
		PerPage:    getLimit(perPage),
		TotalPages: getTotalPages(totalData, perPage),
		TotalData:  totalData,
	}
}

// NewCursorPagination pagination suite with ordered data and cannot be skipped with limit only one
func NewCursorPagination(previous, next string) *CursorPagination {
	// encode into base64 url safe
	previousBase64 := base64.RawURLEncoding.EncodeToString([]byte(previous))
	nextBase64 := base64.RawStdEncoding.EncodeToString([]byte(next))
	return &CursorPagination{
		Previous: &previousBase64,
		Next:     &nextBase64,
	}
}
