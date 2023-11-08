package utils

import "time"

// Timestamp format t into readable format
func Timestamp(t time.Time) string {
	return t.Format("2006-02-02 15:04:05")
}
