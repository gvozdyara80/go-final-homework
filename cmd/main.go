package main

import (
	"log"

	"github.com/go-final-homework/cmd/server"
	configs "github.com/go-final-homework/config"
	"github.com/go-final-homework/repository"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := configs.InitConfig(); err != nil {
		log.Fatalf("error occured while initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("error occured while initializing db: %s", err.Error())
	}

	log.Println("successfully connected to db")

	server := server.NewServer(viper.GetString("port"), db)
	if err := server.Run(); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}

	log.Println("Server is running on port", viper.GetString("port"))
}
