package users

import (
	"encoding/json"
	"fmt"
	"github.com/sigchat/sc-users/pkg/domain/dto"
	"github.com/sigchat/sc-users/pkg/parameters"
	"github.com/sigchat/sc-users/pkg/usecase/users"
	"github.com/valyala/fasthttp"
)

type Controller struct {
	interactor *users.Interactor
}

func NewController(interactor *users.Interactor) *Controller {
	return &Controller{interactor: interactor}
}

func (ctrl *Controller) RegisterUser(ctx *fasthttp.RequestCtx) (interface{}, error) {
	var userDTO dto.RegisterUserDTO
	if err := json.Unmarshal(ctx.Request.Body(), &userDTO); err != nil {
		return nil, fmt.Errorf("bad request")
	}
	return ctrl.interactor.CreateUser(ctx, &userDTO)
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
		return nil, fmt.Errorf("bad request: %v", err)
	}
	return ctrl.interactor.UpdateUserByID(ctx, params.UserID, &userDTO)
}

func (ctrl *Controller) DeleteUserByID(ctx *fasthttp.RequestCtx) error {
	var params parameters.UserIDParams
	if err := params.Get(ctx); err != nil {
		return fmt.Errorf("bad request: %v", err)
	}

	return ctrl.interactor.DeleteUser(ctx, params.UserID)
}
