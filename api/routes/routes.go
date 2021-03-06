package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserRoutes),
	fx.Provide(NewCategoryRoutes),
	fx.Provide(NewBrandRoutes),
	fx.Provide(NewProductRoute),
	fx.Provide(NewOrderRoutes),
	fx.Provide(NewVendorRoutes),
	fx.Provide(NewShippingAddressRoute),
	fx.Provide(NewRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	userRoutes UserRoutes,
	categoryRoutes CategoryRoutes,
	brandRoutes BrandRoutes,
	productRoutes ProductRoutes,
	orderRoutes OrderRoutes,
	vendorRoutes VendorRoutes,
	shippingRoutes ShippingAddressRoute,
) Route {
	return Routes{
		userRoutes,
		categoryRoutes,
		brandRoutes,
		productRoutes,
		orderRoutes,
		vendorRoutes,
		shippingRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
