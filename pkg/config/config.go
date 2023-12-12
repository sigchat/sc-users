package config

type Config struct {
	Name string `env:"USERS_SERVICE_NAME" envDefault:"users"`
}
