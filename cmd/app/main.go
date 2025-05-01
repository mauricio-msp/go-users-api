package main

import (
	"log/slog"

	"github.com/mauricio-msp/go-users-api/cmd/internal/domain"
	"github.com/mauricio-msp/go-users-api/cmd/internal/repository"
	"github.com/mauricio-msp/go-users-api/cmd/internal/service"
	"github.com/mauricio-msp/go-users-api/cmd/internal/web/server"
)

func main() {
	dbInMemory := make(map[string]*domain.User)

	userRepository := repository.NewUserRepository(dbInMemory)
	userService := service.NewUserService(userRepository)

	srv := server.NewServer(userService)
	srv.Routes()

	if err := srv.Start(); err != nil {
		slog.Error("starting server: ", "error", err)
	}
}
