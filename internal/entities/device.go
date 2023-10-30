package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type Device struct {
	bun.BaseModel `bun:"table:devices"`

	ID        string `bun:",pk"`
	UserId    string `bun:",notnull"`
	Name      string `bun:",notnull"`
	IPAddress string `bun:"ip_address"`
	Platform  *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
