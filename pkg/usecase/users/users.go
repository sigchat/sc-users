package users

import (
	"context"
	"github.com/sigchat/sc-users/pkg/domain/dto"
	"github.com/sigchat/sc-users/pkg/repository"
)

type Interactor struct {
	repository repository.Repository
}

func NewInteractor(repository repository.Repository) *Interactor {
	return &Interactor{repository: repository}
}

func (in *Interactor) CreateUser(ctx context.Context, request *dto.CreateUserDTO) (id int, err error) {
	in.repository.CreateUser(ctx)
}
