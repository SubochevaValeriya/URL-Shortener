package main

import (
	urls "URLShortener"
	"URLShortener/internal/handler"
	"URLShortener/internal/repository"
	"URLShortener/internal/service"
	"context"
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
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing congigs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	//mongoDbConnection := "mongodb://localhost:27017" // заменить

	mongoConfig := repository.MongoConfig{
		Host:            viper.GetString("db.host"),
		Port:            viper.GetString("db.port"),
		DefaultDatabase: viper.GetString("db.default_database"),
	}

	//	mongoDbConnection := "mongodb://" + viper.GetString("db.host") + ":" + viper.GetString("db.port")

	dbTables := repository.DbTables{CollectionName: mongoConfig.DefaultDatabase}
	mh, err := repository.NewHandler(mongoConfig) //Create an instance of MongoHander with the connection string provided

	//dbTables := repository.DbTables{CollectionName: viper.GetString("db.default_database")}
	//mh, err := repository.NewHandler(mongoDbConnection, dbTables.CollectionName) //Create an instance of MongoHander with the connection string provided
	if err != nil {
		logrus.Fatalf("failed to inititalize db: %s", err.Error())
	}

	// dependency injection
	repos := repository.NewRepository(mh, dbTables)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(urls.Server)
	r := handler.RegisterRoutes(handlers)

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

	//log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), r))

	//walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
	//	fmt.Printf("%s %s\n", method, route)
	//	return nil
	//}
	//
	//if err := chi.Walk(r, walkFunc); err != nil {
	//	fmt.Printf("Logging err: %s\n", err.Error())
	//}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
