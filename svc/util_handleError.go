package svc

import (
	"fmt"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"opsd/dto"
)

func handleError(ctx *fasthttp.RequestCtx, code int, err error) {
	ctx.SetContentType("application/json")
	b, err := ffjson.Marshal(&dto.Error{
		Status:  "failure",
		Message: fmt.Sprintf("could not create order: %s", err),
	})
	zap.L().With(zap.Error(err)).Error("could not marshal response")
	if err != nil {
		ctx.SetStatusCode(500)
		// if error marshaling is broken, use hardcoded marshaled error message
		ctx.SetBody([]byte(`{"status": "failure", "message": "Internal Server Error"`))
		zap.L().With(zap.Error(err)).Error("could not marshal response")
		return
	}
	ctx.SetStatusCode(400)
	ctx.SetBody(b)
}
