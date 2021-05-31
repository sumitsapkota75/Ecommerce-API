package fxmodule

import (
	"context"
	"travel/api/routes"
	"travel/infrastructure"

	"go.uber.org/fx"
)

var Module = fx.Options(
	infrastructure.Module,
	routes.Module,
	fx.Invoke(fxmodule),
)

func fxmodule(
	lifecycle fx.Lifecycle,
	handler infrastructure.RequestHandler,
	logger infrastructure.Logger,
	env infrastructure.Env,
	migrations infrastructure.Migrations,
	database infrastructure.Database,
	routes routes.Routes,
) {
	conn, _ := database.DB.DB()
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("------- Starting Application --")
			logger.Zap.Info("-------------------------------")
			logger.Zap.Info("------- Travel Agency API 📒 -------")
			logger.Zap.Info("-------------------------------")
			logger.Zap.Info(" 🚌 Migrating DB Schema .......")
			migrations.Migrate()

			conn.SetMaxOpenConns(10)
			go func() {
				logger.Zap.Info("🖇️  Seting up route ....")
				routes.Setup()

				logger.Zap.Info(" 🌱 Seeding data ......")
				// seeds.Run()
				if env.ServerPort == "" {
					handler.Gin.Run()
				} else {
					handler.Gin.Run(env.ServerPort)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Zap.Info(" 🛑 Stopping Application .....")
			conn.Close()
			return nil
		},
	})
}
