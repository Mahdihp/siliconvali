package dto

type RoleInfo struct {
	Id          int16  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type GetAllRoleRequest struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}
