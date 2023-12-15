package users

import (
	"context"
	"fmt"
	"github.com/sigchat/sc-http/pkg/domain/tokens"
	"github.com/sigchat/sc-http/pkg/transport/errors"
	"github.com/sigchat/sc-users/pkg/domain/dto"
	"github.com/sigchat/sc-users/pkg/domain/model"
	"github.com/sigchat/sc-users/pkg/domain/presenter"
	"github.com/sigchat/sc-users/pkg/items"
	"github.com/sigchat/sc-users/pkg/repository/users"
	"github.com/sigchat/sc-users/pkg/usecase/sessions"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

type Interactor struct {
	repository         users.Repository
	userPresenter      *presenter.UserPresenter
	sessionsInteractor *sessions.Interactor
}

func NewInteractor(
	repository users.Repository,
	sessionsInteractor *sessions.Interactor,
) *Interactor {
	return &Interactor{
		userPresenter:      presenter.NewUserPresenter(),
		repository:         repository,
		sessionsInteractor: sessionsInteractor,
	}
}

func (in *Interactor) RegisterUser(
	ctx context.Context,
	request *dto.RegisterUserRequestDTO,
) (responseDTO *dto.RegisterUserResponseDTO, err error) {
	createdUserID, err := in.repository.CreateUser(ctx, request)
	if err != nil {
		return nil, err
	}

	session, err := in.sessionsInteractor.GetOrCreateSession(ctx, createdUserID)
	if err != nil {
		return nil, err
	}

	response := &dto.RegisterUserResponseDTO{
		UserID:      createdUserID,
		AccessToken: session.AccessToken,
	}

	return response, nil
}

func (in *Interactor) LoginUser(
	ctx context.Context,
	request *dto.LoginUserRequest,
) (responseDTO *dto.LoginUserResponse, err error) {
	usersSlice, _ := in.repository.GetUsers(ctx)
	usersList := items.List[model.User](usersSlice)
	usersList.Filter(func(item model.User, index int) bool {
		return item.Username == request.Username
	})

	if len(usersList) == 0 {
		return nil, errors.NewHttpError().
			WithCode(http.StatusUnauthorized).
			WithMessage(fmt.Sprintf("user with username=%s not found", request.Username))
	}

	foundUser := usersList[0]
	if err := bcrypt.CompareHashAndPassword(foundUser.Password, []byte(request.Password)); err != nil {
		return nil, errors.NewHttpError().
			WithCode(http.StatusUnauthorized).
			WithMessage(fmt.Sprintf("invalid credentials"))
	}

	session, err := in.sessionsInteractor.GetOrCreateSession(ctx, foundUser.ID)
	if err != nil {
		return nil, err
	}

	response := &dto.LoginUserResponse{
		UserID:      foundUser.ID,
		AccessToken: session.AccessToken,
	}

	return response, nil
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

func (in *Interactor) GetUsersLikeUsername(ctx context.Context, likeUsername string, curUser *tokens.CurrentUser) ([]model.User, error) {
	usersSlice, _ := in.repository.GetUsers(ctx)
	usersList := items.List[model.User](usersSlice)
	usersList = usersList.Filter(func(item model.User, index int) bool {
		if item.ID == curUser.UserID {
			return false
		}
		return strings.Contains(item.Username, likeUsername)
	})

	return in.userPresenter.UserSearchResponse(usersList.Slice()), nil
}

func (in *Interactor) UpdateUserByID(ctx context.Context, id int, data *dto.UpdateUserDTO) (modified *model.User, err error) {
	return in.repository.UpdateUserByID(ctx, id, data)
}

func (in *Interactor) DeleteUser(ctx context.Context, id int) error {
	return in.repository.DeleteUser(ctx, id)
}
