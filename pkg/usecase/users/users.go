package users

import (
	"context"
	"fmt"
	"github.com/sigchat/sc-http/pkg/transport/errors"
	"github.com/sigchat/sc-users/pkg/domain/dto"
	"github.com/sigchat/sc-users/pkg/domain/model"
	"github.com/sigchat/sc-users/pkg/items"
	"github.com/sigchat/sc-users/pkg/repository"
	"net/http"
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

func (in *Interactor) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	u, err := in.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := items.List[model.User](u).Filter(func(item model.User, index int) bool {
		return item.ID == id
	})
	if len(users) == 0 {
		return nil, errors.NewHttpError().
			WithCode(http.StatusNotFound).
			WithMessage(fmt.Sprintf("user with id %d not found", id))

	}

	return &users[0], nil
}

func (in *Interactor) UpdateUserByID(ctx context.Context, id int, data *dto.UpdateUserDTO) (modified *model.User, err error) {
	return in.repository.UpdateUserByID(ctx, id, data)
}

func (in *Interactor) DeleteUser(ctx context.Context, id int) error {
	return in.repository.DeleteUser(ctx, id)
}
