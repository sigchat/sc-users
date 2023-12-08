package repository

import (
	"context"
	"github.com/sigchat/sc-users/pkg/domain/dto"
)

type Repository interface {
	CreateUser(ctx context.Context, request *dto.CreateUserDTO) error
}
