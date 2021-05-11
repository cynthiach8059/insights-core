package jwt

type TokenPayload struct {
	UserID   string `json:"user_id"`
	UserType string `json:"user_type"`
	UUID     string `json:"-"`
}
