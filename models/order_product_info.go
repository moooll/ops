package models

import (
	"encoding/json"

	"github.com/gofrs/uuid"
)

// OrderProductInfo is used by pop to map your order_products database table to your go code.
type OrderProductInfo struct {
	OrderID            uuid.UUID `json:"order_id" db:"order_id"`
	ProductID          uuid.UUID `json:"product_id" db:"product_id"`
	ProductName        string    `json:"product_name" db:"product_name"`
	ProductDescription string    `json:"product_description" db:"product_description"`
	Quantity           int       `json:"quantity" db:"quantity"`
}

func (opi OrderProductInfo) TableName() string {
	return "order_products_info"
}

// String is not required by pop and may be deleted
func (o OrderProductInfo) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// OrderProducts is not required by pop and may be deleted
type OrderProductInfos []OrderProductInfo

// String is not required by pop and may be deleted
func (o OrderProductInfos) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}
