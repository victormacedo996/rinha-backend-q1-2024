package config

import (
	"os"
	"strconv"
	"sync"
)

type WebServer struct {
	SERVER_PORT int
	TIMEOUT     int
}

type Database struct {
	USER     string
	PWD      string
	HOST     string
	PORT     string
	DATABASE string
	MAX_CONS int
	MIN_CONS int
}

type Redis struct {
	HOST string
	DB   int
	PWD  string
	USER string
}

type Config struct {
	WebServer WebServer
	Database  Database
	Redis     Redis
}

var once sync.Once

var config *Config

func GetInstance() *Config {

	if config == nil {
		once.Do(
			func() {
				webserver_config := WebServer{
					SERVER_PORT: parseEnvToInt("SERVER_PORT", "5000"),
					TIMEOUT:     parseEnvToInt("TIMEOUT", "10"),
				}

				database_config := Database{
					USER:     getEnv("POSTGRES_USERNAME", "postgres"),
					PWD:      getEnv("POSTGRES_PASSWORD", "postgres"),
					HOST:     getEnv("POSTGRES_HOST", "localhost"),
					PORT:     getEnv("POSTGRES_PORT", "5432"),
					DATABASE: getEnv("POSTGRES_DATABASE", "postgres"),
					MAX_CONS: parseEnvToInt("POSTGRES_MIN_CONNS", "2"),
					MIN_CONS: parseEnvToInt("POSTGRES_MAX_CONNS", "5"),
				}

				redis_config := Redis{}

				config = &Config{
					WebServer: webserver_config,
					Database:  database_config,
					Redis:     redis_config,
				}
			},
		)
	}

	return config

}

func parseEnvToInt(envName, defaultValue string) int {
	num, err := strconv.Atoi(getEnv(envName, defaultValue))
	if err != nil {
		return 0
	}
	return num
}

func getEnv(env, defaultValue string) string {
	enviroment := os.Getenv(env)
	if enviroment == "" {
		return defaultValue
	}

	return enviroment
}
