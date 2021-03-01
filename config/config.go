package config

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseHost   string
	DatabasePort   string
	DatabaseUser   string
	DatabasePass   string
	DatabaseName   string
	DatabaseDriver string
}

func GetEnvFile(key, defVal string) string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return defVal
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion %v", key)
	}
	return value
}

func NewConfig() *Config {
	conf := Config{}
	conf.DatabaseHost = os.Getenv("DB_HOST")
	if os.Getenv("DB_HOST") == "" {
		conf.DatabaseHost = GetEnvFile("DB_HOST", "localhost")
	}
	conf.DatabasePort = os.Getenv("DB_PORT")
	if os.Getenv("DB_PORT") == "" {
		conf.DatabasePort = GetEnvFile("DB_PORT", "3306")
	}
	conf.DatabaseUser = os.Getenv("DB_USER")
	if os.Getenv("DB_USER") == "" {
		conf.DatabaseUser = GetEnvFile("DB_USER", "root")
	}
	conf.DatabasePass = os.Getenv("DB_PASS")
	if os.Getenv("DB_PASS") == "" {
		conf.DatabasePass = GetEnvFile("DB_PASS", "")
	}
	conf.DatabaseName = os.Getenv("DB_SCHEMA")
	if os.Getenv("DB_SCHEMA") == "" {
		conf.DatabaseName = GetEnvFile("DB_SCHEMA", "db_user")
	}
	conf.DatabaseDriver = os.Getenv("DB_DRIVER")
	if os.Getenv("DB_DRIVER") == "" {
		conf.DatabaseDriver = GetEnvFile("DB_DRIVER", "mysql")
	}
	return &conf
}
func StartRouter(router *mux.Router) {
	serverHost := os.Getenv("SERVER_HOST")
	if os.Getenv("SERVER_HOST") == "" {
		serverHost = GetEnvFile("SERVER_HOST", "localhost")
	}

	serverPort := os.Getenv("SERVER_PORT")
	if os.Getenv("SERVER_PORT") == "" {
		serverPort = GetEnvFile("SERVER_PORT", "9999")
	}
	address := fmt.Sprintf("%s:%s", serverHost, serverPort)
	log.Printf("Server is now listening at %v.....\n", serverPort)
	err := http.ListenAndServe(address, router)
	if err != nil {
		panic(err)
	}

}
