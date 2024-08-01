package main

import (
	"preference-game/internal/api"
	"preference-game/internal/bootstrap"
	"preference-game/internal/repository"
	"preference-game/internal/service"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Options(
			service.NewModule(),
			repository.NewModule(),
			api.NewModule(),
		),
		fx.Provide(
			bootstrap.NewConfig,
			bootstrap.NewPostgresClient,
		),
	).
		Run()
}
