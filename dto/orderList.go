package dto

import "github.com/gofrs/uuid"

type OrderProduct struct {
	ID          uuid.UUID `json:"id"`
	Quantity    int       `json:"quantity"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type OrderListItem struct {
	ID       uuid.UUID      `json:"id"`
	OwnerID  uuid.UUID      `json:"ownerID"`
	Products []OrderProduct `json:"products"`
}

type OrderList struct {
	List   []OrderListItem `json:"list"`
	Status string          `json:"status"`
}
