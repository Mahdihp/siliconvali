package dto

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User   UserInfo `json:"user"`
	Tokens Tokens   `json:"tokens"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type UserInfo struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Mobile       string `json:"mobile"`
	NationalCode string `json:"national_code"`
	Active       bool   `json:"active"`
	Address      string `json:"address"`
}

type RegisterRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Mobile       string `json:"mobile"`
	NationalCode string `json:"national_code"`
	Address      string `json:"address"`
}
type RegisterResponse struct {
	User UserInfo `json:"user"`
}
