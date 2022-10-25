package main

import (
	"context"
	urls "github.com/SubochevaValeriya/URL-Shortener"
	"github.com/SubochevaValeriya/URL-Shortener/internal/handler"
	"github.com/SubochevaValeriya/URL-Shortener/internal/repository"
	"github.com/SubochevaValeriya/URL-Shortener/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// logs
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// configs
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing congigs: %s", err.Error())
	}

	// environmental variables
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	// DB Config
	mongoConfig := repository.MongoConfig{
		Host:            viper.GetString("db.host"),
		Port:            viper.GetString("db.port"),
		DefaultDatabase: viper.GetString("db.default_database"),
	}

	dbTables := repository.DbTables{CollectionName: mongoConfig.DefaultDatabase}
	mh, err := repository.NewHandler(mongoConfig)

	if err != nil {
		logrus.Fatalf("failed to inititalize db: %s", err.Error())
	}

	// dependency injection
	repos := repository.NewRepository(mh, dbTables)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(urls.Server)
	r := handler.RegisterRoutes(handlers)

	// create index on Short URL (because we will search on this parameter)
	err = repos.CreateIndex("short_URL")
	if err != nil {
		logrus.Errorf("failed to create index")
	}

	go func() {
		if err := srv.Run(viper.GetString("port"), r); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("URL-Shortener Started")

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("URL-Shortener Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
