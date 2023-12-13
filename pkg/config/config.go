package config

type UserService struct {
	Name    string `env:"USERS_SERVICE_NAME" envDefault:"users"`
	BaseURL string `env:"USERS_SERVICE_BASE_URL,required"`
}

type JWTAuth struct {
	AuthPrivateKey string `env:"AUTH_PRIVATE_KEY,required"`
}
