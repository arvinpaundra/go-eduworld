package utils

import "strings"

type SQLCondition struct {
	Column   string
	Operator string
	Value    any
}

// GetSelectedFields get desired columns.
// c must be at least represent one column.
func GetSelectedFields(c string) []string {
	if c == "" {
		return []string{"*"}
	}

	if strings.Contains(c, ", ") {
		return strings.Split(c, ", ")
	}

	return strings.Split(c, ",")
}
