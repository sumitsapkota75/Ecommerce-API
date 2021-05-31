package middlewares

import (
	"go.uber.org/fx"
)

// Module Middleware exported
var Module = fx.Options(
	fx.Provide(NewAuthMiddleware),
	fx.Provide(NewDBTransactionMiddleware),
)
