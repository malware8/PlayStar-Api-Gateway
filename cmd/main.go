package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfigs(); err != nil {
		logrus.Fatal("Configuration initialization failed")
	}

	engine := gin.Default()

	authSvc := *auth.

}

func initConfigs() error {
	viper.AddConfigPath("./pkg/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
