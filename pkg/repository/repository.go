package repository

import (
	"context"
	"github.com/sigchat/sc-users/pkg/domain/dto"
	"github.com/sigchat/sc-users/pkg/domain/model"
)

type Repository interface {
	CreateUser(ctx context.Context, request *dto.CreateUserDTO) (int, error)
	GetUsers(ctx context.Context) ([]model.User, error)
	UpdateUser(ctx context.Context, id int, data *dto.UpdateUserDTO) (modified *model.User, err error)
	DeleteUser(ctx context.Context, id int) error
}
