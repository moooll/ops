package dto

import (
	"opsd/models"

	"github.com/gofrs/uuid"
)

type OrderCreate struct {
	UserID   uuid.UUID `json:"userID"`
	Products []struct {
		ID       uuid.UUID `json:"id"`
		Quantity int       `json:"quantity"`
		} `json:"products"`
	// TODO: add tokens for integration with user service
}

type OrderCreateResponse struct {
	Order  *models.Order `json:"order"`
	Status string        `json:"status"`
}
