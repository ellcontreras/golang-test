package main

import (
	"github.com/labstack/echo/v4"

	handler "stratplus/cmd/users"
	"stratplus/gateway"
	"stratplus/types"
)

func main() {
	db, err := gateway.ConnectDB()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&types.User{})

	// Handlers
	userHandler := handler.NewUserHandler(db)

	e := echo.New()

	e.POST("/api/v1/users", userHandler.CreateUserHandler)
	e.POST("/api/v1/users/auth", userHandler.AuthenticateUserHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
