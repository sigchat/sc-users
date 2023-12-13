package users

import (
	"encoding/json"
	"github.com/sigchat/sc-http/pkg/transport/errors"
	"github.com/sigchat/sc-users/pkg/domain/dto"
	"github.com/sigchat/sc-users/pkg/parameters"
	"github.com/sigchat/sc-users/pkg/usecase/users"
	"github.com/valyala/fasthttp"
	"net/http"
)

type Controller struct {
	interactor *users.Interactor
}

func NewController(usersInteractor *users.Interactor) *Controller {
	return &Controller{interactor: usersInteractor}
}

func (ctrl *Controller) RegisterUser(ctx *fasthttp.RequestCtx) (interface{}, error) {
	var userDTO dto.RegisterUserRequestDTO
	if err := json.Unmarshal(ctx.Request.Body(), &userDTO); err != nil {
		return nil, errors.NewHttpError().WithCode(http.StatusBadRequest)
	}
	return ctrl.interactor.RegisterUser(ctx, &userDTO)
}

func (ctrl *Controller) LoginUser(ctx *fasthttp.RequestCtx) (any, error) {
	var userDTO dto.LoginUserRequest
	if err := json.Unmarshal(ctx.Request.Body(), &userDTO); err != nil {
		return nil, errors.NewHttpError().WithCode(http.StatusBadRequest)
	}
	return ctrl.interactor.LoginUser(ctx, &userDTO)
}

func (ctrl *Controller) GetUsers(ctx *fasthttp.RequestCtx) (interface{}, error) {
	return ctrl.interactor.GetUsers(ctx)
}

func (ctrl *Controller) GetUserByID(ctx *fasthttp.RequestCtx) (interface{}, error) {
	var params parameters.UserIDParams
	if err := params.Get(ctx); err != nil {
		return nil, err
	}
	return ctrl.interactor.GetUserByID(ctx, params.UserID)
}

func (ctrl *Controller) UpdateUserByID(ctx *fasthttp.RequestCtx) (interface{}, error) {
	var params parameters.UserIDParams
	if err := params.Get(ctx); err != nil {
		return nil, err
	}

	var userDTO dto.UpdateUserDTO
	if err := json.Unmarshal(ctx.Request.Body(), &userDTO); err != nil {
		return nil, errors.NewHttpError().WithCode(http.StatusBadRequest)
	}
	return ctrl.interactor.UpdateUserByID(ctx, params.UserID, &userDTO)
}

func (ctrl *Controller) DeleteUserByID(ctx *fasthttp.RequestCtx) error {
	var params parameters.UserIDParams
	if err := params.Get(ctx); err != nil {
		return errors.NewHttpError().WithCode(http.StatusBadRequest)
	}

	return ctrl.interactor.DeleteUser(ctx, params.UserID)
}
