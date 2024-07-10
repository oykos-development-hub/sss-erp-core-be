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

	ErrorLogService := services.NewErrorLogServiceImpl(cel, models.ErrorLog)
	ErrorLogHandler := handlers.NewErrorLogHandler(cel, ErrorLogService)

	UserService := services.NewUserServiceImpl(cel, models.User)
	UserHandler := handlers.NewUserHandler(cel, UserService, ErrorLogService)

	AuthService := services.NewAuthServiceImpl(cel, models.User, models.Log)
	AuthHandler := handlers.NewAuthHandler(cel, AuthService, ErrorLogService)

	RoleService := services.NewRoleServiceImpl(cel, models.Role)
	RoleHandler := handlers.NewRoleHandler(cel, RoleService, ErrorLogService)

	SettingService := services.NewSettingServiceImpl(cel, models.Setting)
	SettingHandler := handlers.NewSettingHandler(cel, SettingService, ErrorLogService)

	BankAccountService := services.NewBankAccountServiceImpl(cel, models.BankAccount)
	BankAccountHandler := handlers.NewBankAccountHandler(cel, BankAccountService, ErrorLogService)

	SupplierService := services.NewSupplierServiceImpl(cel, models.Supplier, models.BankAccount)
	SupplierHandler := handlers.NewSupplierHandler(cel, SupplierService, ErrorLogService)

	AccountService := services.NewAccountServiceImpl(cel, models.Account)
	AccountHandler := handlers.NewAccountHandler(cel, AccountService, ErrorLogService)

	NotificationService := services.NewNotificationServiceImpl(cel, models.Notification)
	NotificationHandler := handlers.NewNotificationHandler(cel, NotificationService, ErrorLogService)

	RolesPermissionService := services.NewRolesPermissionServiceImpl(cel, models.RolesPermission)
	RolesPermissionHandler := handlers.NewRolesPermissionHandler(cel, RolesPermissionService, ErrorLogService)

	PermissionService := services.NewPermissionServiceImpl(cel, models.Permission)
	PermissionHandler := handlers.NewPermissionHandler(cel, PermissionService, ErrorLogService)

	LogService := services.NewLogServiceImpl(cel, models.Log)
	LogHandler := handlers.NewLogHandler(cel, LogService, ErrorLogService)

	TemplateService := services.NewTemplateServiceImpl(cel, models.Template)
	TemplateHandler := handlers.NewTemplateHandler(cel, TemplateService, ErrorLogService)

	TemplateItemService := services.NewTemplateItemServiceImpl(cel, models.TemplateItem)
	TemplateItemHandler := handlers.NewTemplateItemHandler(cel, TemplateItemService, ErrorLogService)

	BffErrorLogService := services.NewBffErrorLogServiceImpl(cel, models.BffErrorLog)
	BffErrorLogHandler := handlers.NewBffErrorLogHandler(cel, BffErrorLogService, ErrorLogService)

	myHandlers := &handlers.Handlers{
		UserHandler:            UserHandler,
		AuthHandler:            AuthHandler,
		RoleHandler:            RoleHandler,
		SettingHandler:         SettingHandler,
		SupplierHandler:        SupplierHandler,
		AccountHandler:         AccountHandler,
		NotificationHandler:    NotificationHandler,
		RolesPermissionHandler: RolesPermissionHandler,
		PermissionHandler:      PermissionHandler,
		BankAccountHandler:     BankAccountHandler,
		LogHandler:             LogHandler,
		TemplateHandler:        TemplateHandler,
		TemplateItemHandler:    TemplateItemHandler,
		ErrorLogHandler:        ErrorLogHandler,
		BffErrorLogHandler:     BffErrorLogHandler,
	}

	myMiddleware := &middleware.Middleware{
		App: cel,
	}

	cel.Routes = routes(cel, myMiddleware, myHandlers)

	return cel
}
