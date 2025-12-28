package main

import (
	"advancedbank/internal/handlers"
	"advancedbank/internal/storage"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	store := storage.New()

	r := gin.Default()
	h := handlers.New(store)
	handlers.Init(r, h)

	go r.Run(":8080")

	Shutdown()
}

func Shutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shut down successfully")
}
