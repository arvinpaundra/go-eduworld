package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID             string  `bun:",pk"`
	InterestId     string  `bun:",notnull"`
	Email          *string `bun:",unique"`
	Username       string  `bun:",unique,notnull"`
	Password       string
	Fullname       string `bun:",notnull"`
	Status         string `bun:",notnull"`
	Role           string `bun:",notnull"`
	Bio            *string
	Phone          *string
	BirthDate      *time.Time
	ProfilePicture *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Interest       *Interest `bun:"rel:belongs-to,join:interest_id=id"`
	Device         *Device   `bun:"rel:has-one,join:id=user_id"`
	Session        *Session  `bun:"rel:has-one,join:id=user_id"`
}
