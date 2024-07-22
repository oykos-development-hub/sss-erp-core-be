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
			rt.Put("/users/{id}", handlers.UserHandler.UpdateUser)
			rt.Delete("/users/{id}", handlers.UserHandler.DeleteUser)
		})

		rt.Group(func(rt chi.Router) {
			rt.Post("/roles", handlers.RoleHandler.CreateRole)
			rt.Get("/roles/{id}", handlers.RoleHandler.GetRoleById)
			rt.Get("/roles", handlers.RoleHandler.GetRoleList)
			rt.Put("/roles/{id}", handlers.RoleHandler.UpdateRole)
			rt.Delete("/roles/{id}", handlers.RoleHandler.DeleteRole)
			rt.Get("/roles/{id}/permissions", handlers.PermissionHandler.GetPermissionListForRole)
			rt.Post("/roles/{id}/permissions/sync", handlers.RolesPermissionHandler.SyncPermissions)
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

		rt.Post("/accounts", handlers.AccountHandler.CreateAccount)
		rt.Get("/accounts/{id}", handlers.AccountHandler.GetAccountById)
		rt.Get("/accounts", handlers.AccountHandler.GetAccountList)
		rt.Delete("/accounts/{id}", handlers.AccountHandler.DeleteAccount)

		rt.Post("/notifications", handlers.NotificationHandler.CreateNotification)
		rt.Get("/notifications/{id}", handlers.NotificationHandler.GetNotificationById)
		rt.Get("/notifications", handlers.NotificationHandler.GetNotificationList)
		rt.Put("/notifications/{id}", handlers.NotificationHandler.UpdateNotification)
		rt.Delete("/notifications/{id}", handlers.NotificationHandler.DeleteNotification)

		rt.Post("/permissions", handlers.PermissionHandler.CreatePermission)
		rt.Get("/permissions/{id}", handlers.PermissionHandler.GetPermissionById)
		rt.Get("/permissions", handlers.PermissionHandler.GetPermissionList)
		rt.Put("/permissions/{id}", handlers.PermissionHandler.UpdatePermission)
		rt.Delete("/permissions/{id}", handlers.PermissionHandler.DeletePermission)

		rt.Post("/bank-accounts", handlers.BankAccountHandler.CreateBankAccount)
		rt.Get("/bank-accounts/{id}", handlers.BankAccountHandler.GetBankAccountById)
		rt.Get("/bank-accounts", handlers.BankAccountHandler.GetBankAccountList)
		rt.Put("/bank-accounts/{id}", handlers.BankAccountHandler.UpdateBankAccount)
		rt.Delete("/bank-accounts/{title}", handlers.BankAccountHandler.DeleteBankAccount)

		rt.Post("/logs", handlers.LogHandler.CreateLog)
		rt.Get("/logs/{id}", handlers.LogHandler.GetLogById)
		rt.Get("/logs", handlers.LogHandler.GetLogList)
		rt.Put("/logs/{id}", handlers.LogHandler.UpdateLog)
		rt.Delete("/logs/{id}", handlers.LogHandler.DeleteLog)

		rt.Post("/templates", handlers.TemplateHandler.CreateTemplate)
		rt.Get("/templates/{id}", handlers.TemplateHandler.GetTemplateById)
		rt.Get("/templates", handlers.TemplateHandler.GetTemplateList)
		rt.Put("/templates/{id}", handlers.TemplateHandler.UpdateTemplate)
		rt.Delete("/templates/{id}", handlers.TemplateHandler.DeleteTemplate)

		rt.Post("/template-items", handlers.TemplateItemHandler.CreateTemplateItem)
		rt.Get("/template-items/{id}", handlers.TemplateItemHandler.GetTemplateItemById)
		rt.Get("/template-items", handlers.TemplateItemHandler.GetTemplateItemList)
		rt.Put("/template-items/{id}", handlers.TemplateItemHandler.UpdateTemplateItem)
		rt.Delete("/template-items/{id}", handlers.TemplateItemHandler.DeleteTemplateItem)

		rt.Get("/error-logs/{id}", handlers.ErrorLogHandler.GetErrorLogById)
		rt.Get("/error-logs", handlers.ErrorLogHandler.GetErrorLogList)
		rt.Put("/error-logs/{id}", handlers.ErrorLogHandler.UpdateErrorLog)
		rt.Delete("/error-logs/{id}", handlers.ErrorLogHandler.DeleteErrorLog)

		rt.Post("/bff-error-logs", handlers.BffErrorLogHandler.CreateBffErrorLog)
		rt.Get("/bff-error-logs/{id}", handlers.BffErrorLogHandler.GetBffErrorLogById)
		rt.Get("/bff-error-logs", handlers.BffErrorLogHandler.GetBffErrorLogList)
		rt.Put("/bff-error-logs/{id}", handlers.BffErrorLogHandler.UpdateBffErrorLog)
		rt.Delete("/bff-error-logs/{id}", handlers.BffErrorLogHandler.DeleteBffErrorLog)
	
		rt.Post("/customer-supports", handlers.CustomerSupportHandler.CreateCustomerSupport)
rt.Get("/customer-supports/{id}", handlers.CustomerSupportHandler.GetCustomerSupportById)
rt.Get("/customer-supports", handlers.CustomerSupportHandler.GetCustomerSupportList)
rt.Put("/customer-supports/{id}", handlers.CustomerSupportHandler.UpdateCustomerSupport)
rt.Delete("/customer-supports/{id}", handlers.CustomerSupportHandler.DeleteCustomerSupport)
	})

	return app.Routes
}
