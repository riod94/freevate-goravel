package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"

	"goravel/app/http/controllers"
)

func Api() {
	// GROUP API Prefix
	facades.Route().Prefix("/api").Group(func(router route.Router) {
		// Root api
		router.Get("/", func(ctx http.Context) http.Response {
			return ctx.Response().Json(http.StatusOK, http.Json{
				"Hello":   "Goravel",
				"version": support.Version,
			})
		})

		router.Prefix("/users").Group(func(router route.Router) {
			userController := controllers.NewUserController()
			router.Get("", userController.Index)
			router.Post("", userController.Store)
			router.Get("/{id}", userController.Show)
			router.Put("/{id}", userController.Update)
			router.Delete("/{id}", userController.Destroy)
		})

		router.Prefix("/banks").Group(func(router route.Router) {
			bankController := controllers.NewBankController()
			router.Get("", bankController.Index)
			router.Post("", bankController.Store)
			router.Get("/{id}", bankController.Show)
			router.Put("/{id}", bankController.Update)
			router.Delete("/{id}", bankController.Destroy)
		})
	})
}
