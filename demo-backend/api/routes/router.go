package routes

import (
	"demo-backend/api/handlers"
	"demo-backend/api/middlewares"
	"demo-backend/infrastructure/datastores"
	"demo-backend/infrastructure/identity"
	"demo-backend/use_cases/productsuc"
	"demo-backend/use_cases/usermgmtuc"

	"github.com/gofiber/fiber/v2"
)

func InitPublicRoutes(app *fiber.App) {

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to My Demo Rest API"))
	})

	grp := app.Group("/api/v1")

	identityManager := identity.NewIdentityManager()
	registerUseCase := usermgmtuc.NewRegisterUseCase(identityManager)

	grp.Post("/user", handlers.RegisterHandler(registerUseCase))
}

func InitProtectedRoutes(app *fiber.App) {

	grp := app.Group("/api/v1")

	productsDataStore := datastores.NewProductsDataStore()

	createProductUseCase := productsuc.NewCreateProductUseCase(productsDataStore)
	grp.Post("/products", middlewares.NewRequiresRealmRole("admin"),
		handlers.CreateProductHandler(createProductUseCase))

	getProductsUseCase := productsuc.NewGetProductsUseCase(productsDataStore)
	grp.Get("/products", middlewares.NewRequiresRealmRole("viewer"),
		handlers.GetProductsHandler(getProductsUseCase))
}
