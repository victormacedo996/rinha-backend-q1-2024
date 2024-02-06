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

type Config struct {
	WebServer WebServer
	Database  Database
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

				database_config := Database{}

				config = &Config{
					WebServer: webserver_config,
					Database:  database_config,
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
