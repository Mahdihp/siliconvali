package dto

import "time"

type UppInfo struct {
	Id                       int       `json:"id"`
	Amount                   int64     `json:"amount"`
	ReferenceNumber          string    `json:"reference_number"`
	TransactionNumber        string    `json:"transaction_number"`
	SourceAccountNumber      string    `json:"source_account_number"`
	DestinationAccountNumber bool      `json:"destination_account_number"`
	Deleted                  string    `json:"deleted"`
	CreatedAt                time.Time `json:"created_at"`
}

type GetAllUppRequest struct {
	PageIndex       int    `json:"page_index"`
	PageSize        int    `json:"page_size"`
	UserId          int64  `json:"user_id"`
	ReferenceNumber string `json:"reference_number"`
}

type UppUpdateRequest struct {
}
type UppInsertRequest struct {
}
