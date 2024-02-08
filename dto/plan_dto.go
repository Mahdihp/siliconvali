package dto

type PlanInfo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Period      int    `json:"period"`
	Active      bool   `json:"active"`
	Description string `json:"description"`
}
type GetAllPlanRequest struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}
