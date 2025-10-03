package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	port int,
	connStr string,
	tokenTTL time.Duration,
) *App {
	ggrpcApp := grpcapp.New(log, port)

	return &App{
		GRPCServer: ggrpcApp,
	}
}
