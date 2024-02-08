package dto

import "time"

type DeviceDetailsInfo struct {
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
type GetAllDeviceDetailsRequest struct {
	PageIndex       int   `json:"page_index"`
	PageSize        int   `json:"page_size"`
	DeviceDetailsId int64 `json:"devicedetails_id"`
}
type DeviceDetailsUpdateRequest struct {
}
type DeviceDetailsInsertRequest struct {
}

type DeviceDetailsRegisterResponse struct {
	DeviceDetails DeviceDetailsInfo `json:"devicedetails"`
}
