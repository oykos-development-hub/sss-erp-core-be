package main

import (
	"log"
	"os"

	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/handlers"
	"gitlab.sudovi.me/erp/core-ms-api/middleware"

	"github.com/oykos-development-hub/celeritas"
	"gitlab.sudovi.me/erp/core-ms-api/services"
)

func initApplication() *celeritas.Celeritas {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "gitlab.sudovi.me/erp/core-ms-api"

	models := data.New(cel.DB.Pool)

	UserService := services.NewUserServiceImpl(cel, models.User)
	UserHandler := handlers.NewUserHandler(cel, UserService)

	AuthService := services.NewAuthServiceImpl(cel, models.User)
	LogService := services.NewUserAccountLogServiceImpl(cel, models.UserAccountLog)
	AuthHandler := handlers.NewAuthHandler(cel, AuthService, LogService)

	RoleService := services.NewRoleServiceImpl(cel, models.Role)
	RoleHandler := handlers.NewRoleHandler(cel, RoleService)

	SettingService := services.NewSettingServiceImpl(cel, models.Setting)
	SettingHandler := handlers.NewSettingHandler(cel, SettingService)

	SupplierService := services.NewSupplierServiceImpl(cel, models.Supplier)
	SupplierHandler := handlers.NewSupplierHandler(cel, SupplierService)

	AccountService := services.NewAccountServiceImpl(cel, models.Account)
	AccountHandler := handlers.NewAccountHandler(cel, AccountService)

	NotificationService := services.NewNotificationServiceImpl(cel, models.Notification)
	NotificationHandler := handlers.NewNotificationHandler(cel, NotificationService)

	myHandlers := &handlers.Handlers{
		UserHandler:         UserHandler,
		AuthHandler:         AuthHandler,
		RoleHandler:         RoleHandler,
		SettingHandler:      SettingHandler,
		SupplierHandler:     SupplierHandler,
		AccountHandler:      AccountHandler,
		NotificationHandler: NotificationHandler,
	}

	myMiddleware := &middleware.Middleware{
		App: cel,
	}

	cel.Routes = routes(cel, myMiddleware, myHandlers)

	return cel
}
