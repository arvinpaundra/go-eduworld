package utils

import "strings"

// GetSelectedFields get desired columns.
// c must be at least represent one column.
func GetSelectedFields(c string) []string {
	if strings.Contains(c, ", ") {
		return strings.Split(c, ", ")
	}

	return strings.Split(c, ",")
}
