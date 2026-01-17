package main

import (
	"advancedbank/internal/handlers"
	"advancedbank/internal/service"
	"advancedbank/internal/storage"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	store := storage.New()
	svc := service.New(store)
	h := handlers.New(svc)

	handlers.Init(r, h)

	go func() {
		err := r.Run(":8080")
		if err != nil {
			log.Fatal(err)
		}
	}()

	Shutdown()
}

func Shutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shut down successfully")
}
