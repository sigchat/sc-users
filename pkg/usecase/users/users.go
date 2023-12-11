package users

import (
	"context"
	"github.com/sigchat/sc-users/pkg/domain/dto"
	"github.com/sigchat/sc-users/pkg/domain/model"
	"github.com/sigchat/sc-users/pkg/repository"
)

type Interactor struct {
	repository repository.Repository
}

func NewInteractor(repository repository.Repository) *Interactor {
	return &Interactor{repository: repository}
}

func (in *Interactor) CreateUser(ctx context.Context, request *dto.RegisterUserDTO) (id int, err error) {
	return in.repository.CreateUser(ctx, request)
}

func (in *Interactor) GetUsers(ctx context.Context) ([]model.User, error) {
	return in.repository.GetUsers(ctx)
}

func (in *Interactor) UpdateUserByID(ctx context.Context, id int, data *dto.UpdateUserDTO) (modified *model.User, err error) {
	return in.repository.UpdateUserByID(ctx, id, data)
}

func (in *Interactor) DeleteUser(ctx context.Context, id int) error {
	return in.repository.DeleteUser(ctx, id)
}
