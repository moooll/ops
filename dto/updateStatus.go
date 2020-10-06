package dto

import "github.com/gofrs/uuid"

type UpdateStatus struct {
	OrderID uuid.UUID`json:"order_id"`
	Status string `json:"status"`
}

type UpdateStatusResponse struct {
	Status string `json:"status"`
}