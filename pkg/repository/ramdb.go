package repository

import (
	"context"
	"fmt"
	"github.com/sigchat/sc-users/pkg/domain/dto"
	"github.com/sigchat/sc-users/pkg/domain/model"
	"github.com/sigchat/sc-users/pkg/items"
	"golang.org/x/crypto/bcrypt"
	"sync"
	"time"
)

var ()

var _userIDCNT = 1

type RAMDBRepository struct {
	items items.List[*model.User]
	m     sync.RWMutex
}

func NewRAMDBRepository() *RAMDBRepository {
	return &RAMDBRepository{
		items: items.List[*model.User]{},
		m:     sync.RWMutex{},
	}
}

func (r *RAMDBRepository) CreateUser(ctx context.Context, request *dto.RegisterUserDTO) (id int, err error) {
	r.m.Lock()
	defer r.m.Unlock()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	newID := _userIDCNT
	newUser := &model.User{
		ID:            newID,
		Username:      request.Username,
		Password:      hashedPassword,
		CreatedAt:     time.Now(),
		LastUpdatedAt: time.Now(),
		Active:        true,
	}
	_userIDCNT++
	r.items = append(r.items, newUser)

	return newID, nil
}

func (r *RAMDBRepository) GetUsers(ctx context.Context) ([]model.User, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	u := make([]model.User, 0, len(r.items))

	r.items.Each(func(item *model.User, index int) bool {
		u = append(u, *item)
		return false
	})
	return u, nil
}

func (r *RAMDBRepository) UpdateUserByID(ctx context.Context, id int, data *dto.UpdateUserDTO) (modified *model.User, err error) {
	r.m.Lock()
	defer r.m.Unlock()

	r.items.Each(func(item *model.User, index int) bool {
		if item.ID == id {
			modified = &*item
			return true
		}
		return false
	})

	if modified == nil {
		return modified, fmt.Errorf(`user with id=%d not found`, id)
	}

	modified.Username = data.Username
	modified.Password = data.HashedPassword
	modified.LastUpdatedAt = time.Now()
	modified.LastOnline = &data.LastOnline
	modified.Active = data.Active
	return
}

func (r *RAMDBRepository) DeleteUser(ctx context.Context, id int) error {
	r.m.Lock()
	defer r.m.Unlock()

	delIndex := -1
	r.items.Each(func(item *model.User, index int) bool {
		if item.ID == id {
			delIndex = index
			return true
		}
		return false
	})

	if delIndex == -1 {
		return fmt.Errorf(`user with id=%d not found`, id)
	}

	r.items = append(r.items[:delIndex], r.items[delIndex+1:]...)
	return nil
}
