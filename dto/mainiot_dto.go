package dto

import (
	"time"
)

type MainIotGetAllRequest struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}
type MainIotUpdateRequest struct {
	DisplayName  string `json:"display_name"`
	Lat          string `json:"lat"`
	Lon          string `json:"lon"`
	Address      string `json:"address"`
	SerialNumber string `json:"serial_number"`
	MacAddress   string `json:"mac_address"`
	IpRemote     string `json:"ip_remote"`
}
type MainIotInsertRequest struct {
	DisplayName  string `json:"display_name"`
	Lat          string `json:"lat"`
	Lon          string `json:"lon"`
	Address      string `json:"address"`
	SerialNumber string `json:"serial_number"`
	MacAddress   string `json:"mac_address"`
	IpRemote     string `json:"ip_remote"`
}
type MainIotInfo struct {
	Id           int64     `json:"id"`
	DisplayName  string    `json:"display_name"`
	Lat          string    `json:"lat"`
	Lon          string    `json:"lon"`
	Address      string    `json:"address"`
	SerialNumber string    `json:"serial_number"`
	MacAddress   string    `json:"mac_address"`
	IpRemote     string    `json:"ip_remote"`
	Status       string    `json:"status"`
	Active       bool      `json:"active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type MainIotRegisterResponse struct {
	MainIot MainIotInfo `json:"mainiot"`
}
