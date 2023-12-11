package parameters

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

const (
	userID = "userID"
)

type UserIDParams struct {
	UserID int `json:"userID"`
}

func (q *UserIDParams) Get(ctx *fasthttp.RequestCtx) error {
	if !ctx.QueryArgs().Has(userID) {
		return fmt.Errorf("query param %s not found", userID)
	}
	param := ctx.QueryArgs().GetUintOrZero(userID)
	if param == 0 {
		return fmt.Errorf("invalid query param %s", userID)
	}

	q.UserID = param
	return nil
}
