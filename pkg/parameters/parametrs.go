package parameters

import (
	"context"
	"fmt"
	"github.com/sigchat/sc-http/pkg/transport/errors"
	"github.com/valyala/fasthttp"
	"net/http"
	"strconv"
)

const (
	UserIDParam      = "userID"
	SearchQueryParam = "q"
)

type UserIDParams struct {
	UserID int
}

func (q *UserIDParams) Get(ctx context.Context) error {
	//if !ctx.QueryArgs().Has(UserIDParam) {
	//	return fmt.Errorf("query param %s not found", UserIDParam)
	//}
	//param := ctx.QueryArgs().GetUintOrZero(UserIDParam)
	//if param == 0 {
	//	return fmt.Errorf("invalid query param %s", UserIDParam)
	//}
	//
	//q.UserIDParam = param
	userID, err := strconv.Atoi(ctx.Value(UserIDParam).(string))
	if err != nil {

		return errors.NewHttpError().
			WithCode(http.StatusBadRequest).
			WithMessage(fmt.Sprintf("cannot validate %s: %v", UserIDParam, err))
	}
	q.UserID = userID

	return nil
}

type SearchQueryParams struct {
	Query string
}

func (p *SearchQueryParams) Get(ctx *fasthttp.RequestCtx) error {
	if !ctx.QueryArgs().Has(SearchQueryParam) {
		return fmt.Errorf("query param %s not found", SearchQueryParam)
	}
	p.Query = string(ctx.QueryArgs().Peek(SearchQueryParam))
	return nil
}
