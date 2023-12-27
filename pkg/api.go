package pkg

import (
	"context"
	"fmt"
	"github.com/sigchat/sc-http/pkg/transport/controllerwrapper"
	"github.com/sigchat/sc-http/pkg/transport/server"
	usersV1 "github.com/sigchat/sc-users/pkg/app/users"
	"github.com/sigchat/sc-users/pkg/config"
	"github.com/sigchat/sc-users/pkg/domain/dto"
	"github.com/sigchat/sc-users/pkg/parameters"
	sessionsRepository "github.com/sigchat/sc-users/pkg/repository/sessions"
	usersRepository "github.com/sigchat/sc-users/pkg/repository/users"
	sessionsInteractor "github.com/sigchat/sc-users/pkg/usecase/sessions"

	usersInteractor "github.com/sigchat/sc-users/pkg/usecase/users"
)

func InitAPI() {
	var cfg config.JWTAuth
	serv := server.NewServer()
	serv.ParseConfig(&cfg)

	noAuthGroup := serv.GetDefaultRouter()
	authGroup := serv.GetJWTRouterGroup("")

	sessionsRepository := sessionsRepository.NewRAMDBRepository([]byte(cfg.AuthPrivateKey))
	usersRepository := usersRepository.NewRAMDBRepository()

	sessionsInteractor := sessionsInteractor.NewInteractor(sessionsRepository)
	usersInteractor := usersInteractor.NewInteractor(
		usersRepository,
		sessionsInteractor,
	)

	{
		// Тестовые пользователи
		usersInteractor.RegisterUser(context.Background(), &dto.RegisterUserRequestDTO{
			Username: "user1",
			Password: "user1",
		})
		usersInteractor.RegisterUser(context.Background(), &dto.RegisterUserRequestDTO{
			Username: "user2",
			Password: "user2",
		})
		usersInteractor.RegisterUser(context.Background(), &dto.RegisterUserRequestDTO{
			Username: "user3",
			Password: "user3",
		})

	}

	usersV1 := usersV1.NewController(usersInteractor)

	wrapper := controllerwrapper.NewWrapper()

	authGroup.GET("/api/v1/users", wrapper.Data(usersV1.GetUsers))
	authGroup.GET(fmt.Sprintf("/api/v1/users/{%s}", parameters.UserIDParam), wrapper.Data(usersV1.GetUserByID))
	authGroup.GET("/api/v1/users/search", wrapper.Data(usersV1.SearchUsersByUsername))
	noAuthGroup.POST("/api/v1/register", wrapper.Data(usersV1.RegisterUser))
	noAuthGroup.POST("/api/v1/login", wrapper.Data(usersV1.LoginUser))

	serv.Run()
}
