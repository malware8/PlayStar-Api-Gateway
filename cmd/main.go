package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/malware8/PlayStar-Api-Gateway/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfigs(); err != nil {
		logrus.Fatalf("Configuration initialization failed. Error: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Environment variables loading failed. Error: %s", err.Error())
	}

	repository := repository.Repository{}

	repository.NewConnection(
		viper.GetString("postgres.host"),
		viper.GetString("postgres.port"),
		viper.GetString("postgres.username"),
		os.Getenv("POSTGRES_PASSWORD"),
		viper.GetString("postgres.dbname"),
	)

}

func initConfigs() error {
	viper.AddConfigPath("./pkg/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
