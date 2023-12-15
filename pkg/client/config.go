package client

import "time"

type UsersService struct {
	Name    string        `env:"USERS_SERVICE_NAME" envDefault:"users"`
	BaseUrl string        `env:"USERS_SERVICE_BASE_URL,required"`
	Timeout time.Duration `env:"USERS_SERVICE_TIMEOUT" envDefault:"3s"`
}
