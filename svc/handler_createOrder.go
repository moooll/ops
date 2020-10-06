package svc

import (
	"opsd/dto"
	"opsd/models"

	"github.com/gobuffalo/pop/v5"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
)

func (i *Instance) createOrder(ctx *fasthttp.RequestCtx) {
	createRequest := &dto.OrderCreate{}

	err := ffjson.Unmarshal(ctx.Request.Body(), createRequest)
	if err != nil {
		handleError(ctx, 400, err)
		return
	}

	order := &models.Order{
		OwnerID: createRequest.UserID,
		Status:  "queued",
	}
	err = i.db.Transaction(func(tx *pop.Connection) error {
		err := tx.Create(order)
		if err != nil {
			return err
		}
		for _, product := range createRequest.Products {
			err = tx.Create(&models.OrderProduct{
				OrderID:   order.ID,
				ProductID: product.ID,
				Quantity:  product.Quantity,
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		handleError(ctx, 500, err)
		return
	}

	response := &dto.OrderCreateResponse{
		Order:  order,
		Status: "success",
	}
	respBytes, err := ffjson.Marshal(response)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}
	ctx.SetStatusCode(201)
	ctx.SetContentType("application/json")
	ctx.SetBody(respBytes)
}
