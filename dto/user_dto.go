package dto

type UserInsertRequest struct {
	Mobile       string `json:"mobile"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	NationalCode string `json:"national_code"`
	Address      string `json:"address"`
	RoleId       int16  `json:"role_id"`
}
type UserInsertResponse struct {
	User UserInfo `json:"user"`
}

type LoginRequest struct {
	Mobile   string `json:"mobile"`
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
	UserId       int64      `json:"user_id"`
	Mobile       string     `json:"mobile"`
	Password     string     `json:"-"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	NationalCode string     `json:"national_code"`
	Active       bool       `json:"active"`
	Deleted      bool       `json:"deleted"`
	Address      string     `json:"address"`
	Roles        []RoleInfo `json:"roles"`
}

type GetAllUserRequest struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}

type UserUpdateRequest struct {
	UserId       int64  `json:"user_id"`
	Mobile       string `json:"mobile"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	NationalCode string `json:"national_code"`
	Address      string `json:"address"`
	RoleId       int16  `json:"role_id"`
}
type UserUpdateResponse struct {
	User UserInfo `json:"user"`
}
