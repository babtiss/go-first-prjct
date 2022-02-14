package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	todo "go-application"
	"go-application/pkg/handler"
	repository "go-application/pkg/repository"
	"go-application/pkg/service"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if errConfig := initConfig(); errConfig != nil {
		logrus.Fatalf("Error init configs: %s", errConfig.Error())
	}
	if errEnv := godotenv.Load(); errEnv != nil {
		logrus.Fatalf("Error loading env variables: %s", errEnv.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if errSrv := srv.Run(viper.GetString("port"), handlers.InitRoutes()); errSrv != nil {
		logrus.Fatalf("Error occered while running http sever: %s", errSrv.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
