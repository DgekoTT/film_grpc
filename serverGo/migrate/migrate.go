package main

import (
	"film_grpc/serverGo/initializers"
	"film_grpc/serverGo/models"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(
		&models.Film{},
		&models.Genre{},
	)
	if err != nil {
		panic(err)
	}
}
