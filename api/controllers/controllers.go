package controllers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserController),
	fx.Provide(NewCategoryController),
	fx.Provide(NewBrandController),
	fx.Provide(NewProductController),
	fx.Provide(NewOrderController),
	fx.Provide(NewVendorController),
	fx.Provide(NewShippingAddressController),
)
