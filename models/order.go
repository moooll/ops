package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Order is used by pop to map your orders database table to your go code.
type Order struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Status    string    `json:"status" db:"status"`
	OwnerID   uuid.UUID `json:"owner_id" db:"owner_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (o Order) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Orders is not required by pop and may be deleted
type Orders []Order

// String is not required by pop and may be deleted
func (o Orders) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (o *Order) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: o.Status, Name: "Status"},
		&validators.UUIDIsPresent{Field: o.OwnerID, Name: "OwnerID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (o *Order) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return o.Validate(tx)
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (o *Order) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return o.Validate(tx)
}
