package dto

import "time"

type DeviceIotInfo struct {
	DisplayName  string    `json:"display_name"`
	SerialNumber string    `json:"serial_number"`
	TypeDevice   int       `json:"type_device"`
	Status       string    `json:"status"`
	Active       bool      `json:"active"`
	Lat          float64   `json:"lat"`
	Lon          float64   `json:"lon"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type DeviceIotGetAllRequest struct {
	PageIndex int   `json:"page_index"`
	PageSize  int   `json:"page_size"`
	MainIotId int64 `json:"mainiot_id"`
}
type DeviceIotUpdateRequest struct {
}
type DeviceIotInsertRequest struct {
}

type DeviceIotRegisterResponse struct {
	DeviceIot DeviceIotInfo `json:"deviceiot"`
}
