package configs

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	PORT                int
	POSTGRES_PORT       int
	POSTGRES_USER       string
	POSTGRES_HOST       string
	POSTGRES_PASSWORD   string
	POSTGRES_DB         string
	RPC_CLIENT_ENDPOINT string
	APP_NAME            string
}

var EnvConfig = Config{}

func LoadConfig() {
	EnvConfig.RPC_CLIENT_ENDPOINT = os.Getenv("RPC_CLIENT_ENDPOINT")
	EnvConfig.POSTGRES_DB = os.Getenv("POSTGRES_DB")
	EnvConfig.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	EnvConfig.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	EnvConfig.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	EnvConfig.APP_NAME = os.Getenv("APP_NAME")
	EnvConfig.PORT = 5566
	EnvConfig.POSTGRES_PORT = 5432
	PORT, error := strconv.Atoi(os.Getenv("PORT"))
	if error == nil {
		EnvConfig.PORT = PORT
	}
	POSTGRES_PORT, error := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if error == nil {
		EnvConfig.POSTGRES_PORT = POSTGRES_PORT
	}
}
func init() {
	LoadConfig()
}
func GetEnvConfig() Config {
	if EnvConfig.RPC_CLIENT_ENDPOINT == "" {
		LoadConfig()
	}
	return EnvConfig
}
