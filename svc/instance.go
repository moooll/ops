package svc

import (
	"github.com/fasthttp/router"
	"github.com/gobuffalo/pop/v5"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
	"opsd/dto"
	"opsd/models"
)

type Instance struct {
	db *pop.Connection
}

func New(db *pop.Connection) *Instance {
	return &Instance{
		db,
	}
}

func (i *Instance) RouterExtender(r *router.Group) {
	r.GET("/orders", i.orders)
	r.POST("/orders/create", i.createOrder)
	r.POST("/orders/update", i.updateOrder)
}

func (i *Instance) updateOrder(ctx *fasthttp.RequestCtx) {
	request := &dto.UpdateStatus{}
	err := ffjson.Unmarshal(ctx.Request.Body(), request)

	if err != nil {
		handleError(ctx, 500, err)
		return
	}
	order := &models.Order{}
	err = i.db.Find(order, request.OrderID)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}

	order.Status = request.Status
	err = i.db.Update(order)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}
	response := dto.UpdateStatusResponse{}
	respBytes, err := ffjson.Marshal(response)
	if err != nil {
		handleError(ctx, 500, err)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetBody(respBytes)
}
