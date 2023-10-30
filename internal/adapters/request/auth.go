package request

type (
	Login struct {
		Username         string  `json:"username" form:"username" validate:"required"`
		Password         string  `json:"password" form:"password" validate:"required"`
		Device           string  `json:"device" form:"device" validate:"required"`
		Platform         *string `json:"platform" form:"platform"`
		FCMToken         *string `json:"fcm_token" form:"fcm_token"`
		GoogleOAuthToken *string `json:"google_oauth_token" form:"google_oauth_token"`
		IPAddress        string
	}

	Register struct {
		InterestId string `json:"interest_id" form:"interest_id" validate:"required"`
		Username   string `json:"username" form:"username" validate:"required,min=3,max=50"`
		Password   string `json:"password" form:"password" validate:"required,min=8,max=50"`
		Fullname   string `json:"fullname" form:"fullname" validate:"required,min=3,max=255"`
		Role       string `json:"role" form:"role" validate:"required"`
	}
)
