package main

import (
	"gitlab.sudovi.me/erp/core-ms-api/handlers"
	"gitlab.sudovi.me/erp/core-ms-api/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/oykos-development-hub/celeritas"
)

func routes(app *celeritas.Celeritas, middleware *middleware.Middleware, handlers *handlers.Handlers) *chi.Mux {
	// middleware must come before any routes

	//api
	app.Routes.Route("/api", func(rt chi.Router) {

		rt.With(middleware.JwtVerifyRefreshToken).Get("/refresh", handlers.AuthHandler.RefreshToken)
		rt.Post("/users/login", handlers.AuthHandler.Login)
		rt.Post("/users/password/forgot", handlers.AuthHandler.ForgotPassword)
		rt.Get("/users/password/validate-email", handlers.AuthHandler.ResetPasswordVerify)
		rt.Post("/users/password/reset", handlers.AuthHandler.ResetPassword)
		rt.Post("/users", handlers.UserHandler.CreateUser)

		rt.Group(func(rt chi.Router) {
			rt.With(middleware.JwtVerifyToken).Post("/users/validate-pin", handlers.AuthHandler.ValidatePin)
			rt.With(middleware.JwtVerifyToken).Get("/logged-in-user", handlers.UserHandler.GetLoggedInUser)
			rt.With(middleware.JwtVerifyToken).Post("/users/logout", handlers.AuthHandler.Logout)
			rt.Get("/users/{id}", handlers.UserHandler.GetUserById)
			rt.Get("/users", handlers.UserHandler.GetUserList)
			rt.Patch("/users/{id}", handlers.UserHandler.UpdateUser)
			rt.Delete("/users/{id}", handlers.UserHandler.DeleteUser)
		})

		rt.Group(func(rt chi.Router) {
			rt.Post("/roles", handlers.RoleHandler.CreateRole)
			rt.Get("/roles/{id}", handlers.RoleHandler.GetRoleById)
			rt.Get("/roles", handlers.RoleHandler.GetRoleList)
			rt.Patch("/roles/{id}", handlers.RoleHandler.UpdateRole)
			rt.Delete("/roles/{id}", handlers.RoleHandler.DeleteRole)
		})

		rt.Post("/settings", handlers.SettingHandler.CreateSetting)
		rt.Get("/settings/{id}", handlers.SettingHandler.GetSettingById)
		rt.Get("/settings", handlers.SettingHandler.GetSettingList)
		rt.Put("/settings/{id}", handlers.SettingHandler.UpdateSetting)
		rt.Delete("/settings/{id}", handlers.SettingHandler.DeleteSetting)

		rt.Post("/suppliers", handlers.SupplierHandler.CreateSupplier)
		rt.Get("/suppliers/{id}", handlers.SupplierHandler.GetSupplierById)
		rt.Get("/suppliers", handlers.SupplierHandler.GetSupplierList)
		rt.Put("/suppliers/{id}", handlers.SupplierHandler.UpdateSupplier)
		rt.Delete("/suppliers/{id}", handlers.SupplierHandler.DeleteSupplier)
	})

	return app.Routes
}
