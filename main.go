package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MYK12397/go-messenger/internal/adapters/handler"
	"github.com/MYK12397/go-messenger/internal/adapters/repository"
	"github.com/MYK12397/go-messenger/internal/core/services"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var (
	repo      = flag.String("db", "postgres", "Database for storing messages")
	redisHost = "localhost:6379"
	srv       *services.MessengerService
)

func main() {
	err := godotenv.Load()
	if err != nil {

		log.Fatal("Error loading .env file")
	}

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

	e := InitRoutes()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	go func() {

		if err := e.Start(":" + port); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v ", err)
		}
		log.Println("Serving new connections stopped.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	ctxShutdown, releaseShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer releaseShutdown()

	if err := e.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("shutdown error: %v", err)
	}
	log.Println("graceful shutdown completed.")

}

func InitRoutes() *echo.Echo {
	e := echo.New()
	handler := handler.NewHTTPHandler(*srv)

	e.POST("/messages", handler.SaveMessage)
	e.GET("/messages", handler.ReadMessages)
	e.GET("/messages/:id", handler.ReadMessage)
	e.DELETE("/messages/:id", handler.DeleteMessage)

	return e
}
