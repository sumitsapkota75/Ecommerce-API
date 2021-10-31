package services

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewFirebaseService),
	fx.Provide(NewUserService),
	fx.Provide(NewCategoryService),
	fx.Provide(NewBrandService),
	fx.Provide(NewProductService),
	fx.Provide(NewOrderService),
	fx.Provide(NewVendorService),
	fx.Provide(NewShippingAddressService),
)
