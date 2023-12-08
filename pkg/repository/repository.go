package repository

import (
	"context"
	"sc-users/pkg/domain/dto"
)

type Repository interface {
	CreateUser(ctx context.Context, request dto.CreateUserRequest) error
}
