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

	err := store.LoadAccounts()
	if err != nil {
		slog.Error("failed to load accounts from file", "error", err)
		return
	}

	r := gin.Default()
	h := handlers.New(store)
	handlers.Init(r, h)

	go r.Run(":8080")

	Shutdown(store)
}

func Shutdown(store *storage.Storage) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down...")
	store.SaveAccounts()
	slog.Info("shut down successfully")
}
