package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type Session struct {
	bun.BaseModel `bun:"table:sessions"`

	ID               string  `bun:",pk"`
	UserId           string  `bun:",notnull"`
	Token            string  `bun:",notnull"`
	GoogleOAuthToken *string `bun:"google_oauth_token"`
	FCMToken         *string `bun:"fcm_token"`
	RefreshToken     *string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
