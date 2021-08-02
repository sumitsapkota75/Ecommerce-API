package fxmodule

import (
	"context"
	"travel/api/controllers"
	"travel/api/middlewares"
	"travel/api/repository"
	"travel/api/routes"
	"travel/api/services"
	"travel/infrastructure"
	"travel/utils"

	"go.uber.org/fx"
)

var Module = fx.Options(
	infrastructure.Module,
	controllers.Module,
	services.Module,
	repository.Module,
	middlewares.Module,
	routes.Module,
	utils.Module,
	fx.Invoke(fxmodule),
)

func fxmodule(
	lifecycle fx.Lifecycle,
	handler infrastructure.RequestHandler,
	routes routes.Route,
	logger infrastructure.Logger,
	env infrastructure.Env,
	migrations infrastructure.Migrations,
	database infrastructure.Database,
	twilio utils.Twilio,

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
