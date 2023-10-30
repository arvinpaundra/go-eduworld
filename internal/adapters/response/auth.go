package response

import "time"

type (
	Login struct {
		Token     string `json:"token"`
		IssuedAt  string `json:"issued_at"`
		ExpiredAt string `json:"expired_at"`
	}

	CheckSession struct {
		Device     string    `json:"device"`
		IPAdress   string    `json:"ip_address"`
		HasSession bool      `json:"has_session"`
		Platform   *string   `json:"platform"`
		Time       time.Time `json:"time"`
	}
)
