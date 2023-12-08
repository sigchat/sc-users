package users

import (
	"context"
	"sc-users/pkg/domain/dto"
	"sc-users/pkg/repository"
)

type Interactor struct {
	repository repository.Repository
}

func NewInteractor(repository repository.Repository) *Interactor {
	return &Interactor{repository: repository}
}

func (in *Interactor) CreateUser(ctx context.Context, request dto.CreateUserRequest) error {

}
