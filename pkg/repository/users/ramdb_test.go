package users

import (
	"context"
	"github.com/sigchat/sc-users/pkg/domain/dto"
	"github.com/sigchat/sc-users/pkg/domain/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	repo := NewRAMDBRepository()

	createUserDTO := &dto.RegisterUserRequestDTO{
		Username: "testuser",
		Password: "hashedpassword",
	}

	id, err := repo.CreateUser(context.Background(), createUserDTO)

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
	assert.Len(t, repo.items, 1)

	user := repo.items[0]
	assert.Equal(t, id, user.ID)
	assert.Equal(t, createUserDTO.Username, user.Username)
	assert.Equal(t, createUserDTO.Password, user.Password)
	assert.False(t, user.Active)
}

func TestGetUsers(t *testing.T) {
	repo := NewRAMDBRepository()

	users, err := repo.GetUsers(context.Background())

	assert.NoError(t, err)
	assert.Len(t, users, 0)

	// Add a user to the repository
	repo.items = append(repo.items, &model.User{ID: 1, Username: "testuser", Password: []byte("hashedpassword"), Active: false})

	users, err = repo.GetUsers(context.Background())

	assert.NoError(t, err)
	assert.Len(t, users, 1)

	user := users[0]
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "testuser", user.Username)
	assert.Equal(t, "hashedpassword", user.Password)
	assert.False(t, user.Active)
}

func TestUpdateUser(t *testing.T) {
	repo := NewRAMDBRepository()

	// Add a user to the repository
	repo.items = append(repo.items, &model.User{ID: 1, Username: "testuser", Password: []byte("hashedpassword"), Active: false})

	updateUserDTO := &dto.UpdateUserDTO{
		Username:       "updateduser",
		HashedPassword: []byte("updatedpassword"),
		LastOnline:     time.Now(),
		Active:         true,
	}

	modifiedUser, err := repo.UpdateUserByID(context.Background(), 1, updateUserDTO)

	assert.NoError(t, err)
	assert.NotNil(t, modifiedUser)
	assert.Equal(t, updateUserDTO.Username, modifiedUser.Username)
	assert.Equal(t, updateUserDTO.HashedPassword, modifiedUser.Password)
	assert.Equal(t, updateUserDTO.LastOnline, modifiedUser.LastOnline)
	assert.Equal(t, updateUserDTO.Active, modifiedUser.Active)
}

func TestUpdateUserNotFound(t *testing.T) {
	repo := NewRAMDBRepository()

	updateUserDTO := &dto.UpdateUserDTO{
		Username:       "updateduser",
		HashedPassword: []byte("updatedpassword"),
		LastOnline:     time.Now(),
		Active:         true,
	}

	modifiedUser, err := repo.UpdateUserByID(context.Background(), 1, updateUserDTO)

	assert.Error(t, err)
	assert.Nil(t, modifiedUser)
	assert.EqualError(t, err, "user with id=1 not found")
}

func TestDeleteUser(t *testing.T) {
	repo := NewRAMDBRepository()

	// Add a user to the repository
	repo.items = append(repo.items, &model.User{ID: 1, Username: "testuser", Password: []byte("hashedpassword"), Active: false})

	err := repo.DeleteUser(context.Background(), 1)

	assert.NoError(t, err)
	assert.Len(t, repo.items, 0)
}

func TestDeleteUserNotFound(t *testing.T) {
	repo := NewRAMDBRepository()

	err := repo.DeleteUser(context.Background(), 1)

	assert.Error(t, err)
	assert.EqualError(t, err, "user with id=1 not found")
}
