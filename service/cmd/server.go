package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/gox7/authorizathion/internal/repository"
	"github.com/gox7/authorizathion/internal/services"
	"github.com/gox7/authorizathion/internal/transport"
	"github.com/gox7/authorizathion/models"
)

func main() {
	config := new(models.LocalConfig)
	services.NewConfig(config)

	serverLog := new(slog.Logger)
	databaseLog := new(slog.Logger)
	services.NewLoggger("server", &serverLog)
	services.NewLoggger("database", &databaseLog)

	postgres := new(repository.Postgres)
	repository.NewPostgres(config, databaseLog, postgres)
	postgres.Migration()

	sessions := new(services.SessionsService)
	authorizathion := new(services.AuthorizathionService)
	services.NewSessions(postgres, sessions)
	services.NewAuthorizathion(postgres, authorizathion)

	transport.SetResource(
		authorizathion,
		sessions,
		config,
	)

	engine := gin.Default()
	transport.Register(engine, serverLog)
	transport.Listen(engine, config)
}
