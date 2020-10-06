package svc

import (
	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
	"opsd/dto"
	"opsd/models"
)

func (i *Instance) orders(ctx *fasthttp.RequestCtx) {
	response := &dto.OrderList{}

	orders := models.Orders{}
	err := i.db.All(&orders)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}

	for _, order := range orders {
		orderProductInfos := models.OrderProductInfos{}
		err := i.db.Where("order_id = ?", order.ID).All(&orderProductInfos)
		if err != nil {
			handleError(ctx, 500, err)
			return
		}
		op := []dto.OrderProduct{}
		for _, product := range orderProductInfos {
			op = append(op, dto.OrderProduct{
				ID:          product.ProductID,
				Name:        product.ProductName,
				Description: product.ProductDescription,
				Quantity:    product.Quantity,
			})
		}

		response.List = append(response.List, dto.OrderListItem{
			ID:       order.ID,
			OwnerID:  order.OwnerID,
			Products: op,
		})
	}

	respBytes, err := ffjson.Marshal(response)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetBody(respBytes)
}

