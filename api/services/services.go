package services

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewGmailService),
	fx.Provide(NewFirebaseService),
	fx.Provide(NewUserService),
)
