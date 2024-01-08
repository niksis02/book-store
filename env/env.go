package env

import (
	"fmt"

	"github.com/spf13/viper"
)

type envVars struct {
	PORT        int
	DB_PORT     int
	DB_NAME     string
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	JWT_SECRET  string
}

// This config is global and is accessible across all the app
var Env *envVars = &envVars{}

// Loads Environment variables from .env file
func LoadEnvVars() error {
	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	db_host := viper.GetString("DB_HOST")
	if db_host == "" {
		return fmt.Errorf("missing required env variable db_host")
	}
	Env.DB_HOST = db_host

	db_name := viper.GetString("DB_NAME")
	if db_name == "" {
		return fmt.Errorf("missing required env variable db_name")
	}
	Env.DB_NAME = db_name

	db_user := viper.GetString("DB_USER")
	if db_user == "" {
		return fmt.Errorf("missing required env variable db_user")
	}
	Env.DB_USER = db_user

	db_password := viper.GetString("DB_PASSWORD")
	if db_password == "" {
		return fmt.Errorf("missing required env variable db_password")
	}
	Env.DB_PASSWORD = db_password

	db_port := viper.GetInt("DB_PORT")
	if db_port == 0 {
		return fmt.Errorf("missing required env variable db_port")
	}
	Env.DB_PORT = db_port

	port := viper.GetInt("PORT")
	if port == 0 {
		Env.PORT = 8080
	}
	Env.PORT = port

	jwt_secret := viper.GetString("JWT_SECRET")
	if jwt_secret == "" {
		return fmt.Errorf("missing required env variable jwt_secret")
	}
	Env.JWT_SECRET = jwt_secret

	return nil
}
