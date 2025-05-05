package ports

type AuthToken struct {
	AcessToken string          `json:"access_token"`
	ExpiresIn  int             `json:"expires_in"`
	Claims     AuthTokenClaims `json:"claims"`
}

type AuthTokenClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type AuthTokenProvider interface {
	GenerateToken(userID, username string) (AuthToken, error)
	ValidateToken(token string) (AuthTokenClaims, error)
}