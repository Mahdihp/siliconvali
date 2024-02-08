package dto

type UserInsertRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Mobile       string `json:"mobile"`
	NationalCode string `json:"national_code"`
	Address      string `json:"address"`
}
type UserRegisterResponse struct {
	User UserInfo `json:"user"`
}

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
	UserId       int64  `json:"user_id"`
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Mobile       string `json:"mobile"`
	NationalCode string `json:"national_code"`
	Active       bool   `json:"active"`
	Deleted      bool   `json:"deleted"`
	Address      string `json:"address"`
}

type GetAllUserRequest struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}

type UserUpdateRequest struct {
	UserId       int64  `json:"user_id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Mobile       string `json:"mobile"`
	NationalCode string `json:"national_code"`
	Address      string `json:"address"`
}
type UserUpdateResponse struct {
	User UserInfo `json:"user"`
}
