package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags|log.Lshortfile)

	srv := server.NewServer(logger)

	logger.Println("Сервер запускается на порту :8080...")

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
