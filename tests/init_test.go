package tests

import (
	urls "github.com/SubochevaValeriya/URL-Shortener"
	"github.com/SubochevaValeriya/URL-Shortener/internal/handler"
	"github.com/SubochevaValeriya/URL-Shortener/internal/repository"
	"github.com/SubochevaValeriya/URL-Shortener/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func initForTest() (*handler.Handler, *repository.Repository) {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing congigs: %s", err.Error())
	}

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

	return handlers, repos
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func addURLForTest(repos *repository.Repository) {
	urlInfo := urls.UrlInfo{
		OriginalURL: "https://github.com/",
		ShortURL:    " ",
		CreatedAt:   time.Now(),
		Visits:      0,
	}

	repos.AddURL(&urlInfo)
}

func deleteURLForTest(repos *repository.Repository) {
	shortURL := "dhfh3"

	repos.DeleteURL(shortURL)
}
