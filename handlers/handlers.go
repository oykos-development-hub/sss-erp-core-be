package handlers

import (
	"net/http"
)

type Handlers struct {
	UserHandler            UserHandler
	AuthHandler            AuthHandler
	RoleHandler            RoleHandler
	SettingHandler         SettingHandler
	SupplierHandler        SupplierHandler
	AccountHandler         AccountHandler
	NotificationHandler    NotificationHandler
	RolesPermissionHandler RolesPermissionHandler
	PermissionHandler      PermissionHandler
	BankAccountHandler     BankAccountHandler
	LogHandler             LogHandler
	TemplateHandler        TemplateHandler
	TemplateItemHandler    TemplateItemHandler
	ErrorLogHandler        ErrorLogHandler
	BffErrorLogHandler     BffErrorLogHandler
	CustomerSupportHandler CustomerSupportHandler
	ListOfParameterHandler ListOfParameterHandler
	}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	GetUserList(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	GetLoggedInUser(w http.ResponseWriter, r *http.Request)
}

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	ValidatePin(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	ForgotPassword(w http.ResponseWriter, r *http.Request)
	ResetPasswordVerify(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
}

type RoleHandler interface {
	CreateRole(w http.ResponseWriter, r *http.Request)
	UpdateRole(w http.ResponseWriter, r *http.Request)
	DeleteRole(w http.ResponseWriter, r *http.Request)
	GetRoleById(w http.ResponseWriter, r *http.Request)
	GetRoleList(w http.ResponseWriter, r *http.Request)
}

type SettingHandler interface {
	CreateSetting(w http.ResponseWriter, r *http.Request)
	UpdateSetting(w http.ResponseWriter, r *http.Request)
	DeleteSetting(w http.ResponseWriter, r *http.Request)
	GetSettingById(w http.ResponseWriter, r *http.Request)
	GetSettingList(w http.ResponseWriter, r *http.Request)
}

type SupplierHandler interface {
	CreateSupplier(w http.ResponseWriter, r *http.Request)
	UpdateSupplier(w http.ResponseWriter, r *http.Request)
	DeleteSupplier(w http.ResponseWriter, r *http.Request)
	GetSupplierById(w http.ResponseWriter, r *http.Request)
	GetSupplierList(w http.ResponseWriter, r *http.Request)
}

type AccountHandler interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	DeleteAccount(w http.ResponseWriter, r *http.Request)
	GetAccountById(w http.ResponseWriter, r *http.Request)
	GetAccountList(w http.ResponseWriter, r *http.Request)
}

type NotificationHandler interface {
	CreateNotification(w http.ResponseWriter, r *http.Request)
	UpdateNotification(w http.ResponseWriter, r *http.Request)
	DeleteNotification(w http.ResponseWriter, r *http.Request)
	GetNotificationById(w http.ResponseWriter, r *http.Request)
	GetNotificationList(w http.ResponseWriter, r *http.Request)
}

type RolesPermissionHandler interface {
	SyncPermissions(w http.ResponseWriter, r *http.Request)
}

type PermissionHandler interface {
	CreatePermission(w http.ResponseWriter, r *http.Request)
	UpdatePermission(w http.ResponseWriter, r *http.Request)
	DeletePermission(w http.ResponseWriter, r *http.Request)
	GetPermissionById(w http.ResponseWriter, r *http.Request)
	GetPermissionList(w http.ResponseWriter, r *http.Request)
	GetPermissionListForRole(w http.ResponseWriter, r *http.Request)
	GetUsersByPermission(w http.ResponseWriter, r *http.Request)
}

type BankAccountHandler interface {
	CreateBankAccount(w http.ResponseWriter, r *http.Request)
	UpdateBankAccount(w http.ResponseWriter, r *http.Request)
	DeleteBankAccount(w http.ResponseWriter, r *http.Request)
	GetBankAccountById(w http.ResponseWriter, r *http.Request)
	GetBankAccountList(w http.ResponseWriter, r *http.Request)
}

type LogHandler interface {
	CreateLog(w http.ResponseWriter, r *http.Request)
	UpdateLog(w http.ResponseWriter, r *http.Request)
	DeleteLog(w http.ResponseWriter, r *http.Request)
	GetLogById(w http.ResponseWriter, r *http.Request)
	GetLogList(w http.ResponseWriter, r *http.Request)
}

type TemplateHandler interface {
	CreateTemplate(w http.ResponseWriter, r *http.Request)
	UpdateTemplate(w http.ResponseWriter, r *http.Request)
	DeleteTemplate(w http.ResponseWriter, r *http.Request)
	GetTemplateById(w http.ResponseWriter, r *http.Request)
	GetTemplateList(w http.ResponseWriter, r *http.Request)
}

type TemplateItemHandler interface {
	CreateTemplateItem(w http.ResponseWriter, r *http.Request)
	UpdateTemplateItem(w http.ResponseWriter, r *http.Request)
	DeleteTemplateItem(w http.ResponseWriter, r *http.Request)
	GetTemplateItemById(w http.ResponseWriter, r *http.Request)
	GetTemplateItemList(w http.ResponseWriter, r *http.Request)
}

type ErrorLogHandler interface {
	UpdateErrorLog(w http.ResponseWriter, r *http.Request)
	DeleteErrorLog(w http.ResponseWriter, r *http.Request)
	GetErrorLogById(w http.ResponseWriter, r *http.Request)
	GetErrorLogList(w http.ResponseWriter, r *http.Request)
}

type BffErrorLogHandler interface {
	CreateBffErrorLog(w http.ResponseWriter, r *http.Request)
	UpdateBffErrorLog(w http.ResponseWriter, r *http.Request)
	DeleteBffErrorLog(w http.ResponseWriter, r *http.Request)
	GetBffErrorLogById(w http.ResponseWriter, r *http.Request)
	GetBffErrorLogList(w http.ResponseWriter, r *http.Request)
}

type CustomerSupportHandler interface {
	CreateCustomerSupport(w http.ResponseWriter, r *http.Request)
	UpdateCustomerSupport(w http.ResponseWriter, r *http.Request)
	DeleteCustomerSupport(w http.ResponseWriter, r *http.Request)
	GetCustomerSupportById(w http.ResponseWriter, r *http.Request)
	GetCustomerSupportList(w http.ResponseWriter, r *http.Request)
}

type ListOfParameterHandler interface {
	CreateListOfParameter(w http.ResponseWriter, r *http.Request)
	UpdateListOfParameter(w http.ResponseWriter, r *http.Request)
	DeleteListOfParameter(w http.ResponseWriter, r *http.Request)
	GetListOfParameterById(w http.ResponseWriter, r *http.Request)
	GetListOfParameterList(w http.ResponseWriter, r *http.Request)
}
