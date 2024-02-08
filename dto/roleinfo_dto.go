package dto

type RoleInfo struct {
	RoleID      int16  `json:"role_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type GetAllRoleRequest struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}
