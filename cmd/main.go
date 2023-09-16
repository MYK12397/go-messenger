package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/MYK12397/go-messenger/internal/adapters/handler"
	"github.com/MYK12397/go-messenger/internal/adapters/repository"
	"github.com/MYK12397/go-messenger/internal/core/services"
	"github.com/labstack/echo/v4"
)

var (
	repo        = flag.String("db", "postgres", "Database for storing messages")
	redisHost   = "localhost:6379"
	httpHandler *handler.HTTPHandler
	srv         *services.MessengerService
)

func main() {
	flag.Parse()
	fmt.Printf("Application running using %s\n", *repo)

	switch *repo {
	case "redis":
		store := repository.NewMessengerRedisRepository(redisHost)

		srv = services.NewMessengerService(store)
	default:
		store := repository.NewMessengerPostgresRepository()
		srv = services.NewMessengerService(store)
	}

	InitRoutes()
}

func InitRoutes() {
	e := echo.New()
	handler := handler.NewHTTPHandler(*srv)

	e.POST("/messages", handler.SaveMessage)
	e.GET("/messages/:id", handler.ReadMessage)
	e.GET("/messages", handler.ReadMessages)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
