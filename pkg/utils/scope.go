package utils

import "gorm.io/gorm"

type gormScope func(*gorm.DB) *gorm.DB

// MenteeFullname is a scope contain condition to prevent full-text search column `mentees`.`fullname`
// if not necessary needed
func MenteeFullname(fullname string) gormScope {
	return func(db *gorm.DB) *gorm.DB {
		if fullname != "" {
			return db.Where("mentees.fullname LIKE ?", "%"+fullname+"%")
		}

		return db
	}
}

func Conditions(conditions string, args ...interface{}) gormScope {
	return func(db *gorm.DB) *gorm.DB {
		if conditions != "" {
			return db.Where(conditions, args...)
		}

		return db
	}
}
