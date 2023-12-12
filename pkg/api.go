package pkg

import (
	"fmt"
	"github.com/sigchat/sc-http/pkg/transport/controllerwrapper"
	"github.com/sigchat/sc-http/pkg/transport/server"
	usersV1 "github.com/sigchat/sc-users/pkg/app/users"
	"github.com/sigchat/sc-users/pkg/parameters"

	usersRepository "github.com/sigchat/sc-users/pkg/repository"
	usersInteractor "github.com/sigchat/sc-users/pkg/usecase/users"
)

func InitAPI() {
	serv := server.NewServer()

	usersRepository := usersRepository.NewRAMDBRepository()
	usersInteractor := usersInteractor.NewInteractor(usersRepository)
	usersV1 := usersV1.NewController(usersInteractor)

	wrapper := controllerwrapper.NewWrapper()

	serv.GetDefaultRouter().GET("/api/v1/users", wrapper.Data(usersV1.GetUsers))
	serv.GetDefaultRouter().GET(fmt.Sprintf("/api/v1/users/{%s}", parameters.UserIDParam), wrapper.Data(usersV1.GetUserByID))
	serv.GetDefaultRouter().POST("/api/v1/register", wrapper.Data(usersV1.RegisterUser))

	serv.Run()
}
