package handlers

import (
	"net/http"
)

type Handlers struct {
	UserHandler     UserHandler
	AuthHandler     AuthHandler
	RoleHandler     RoleHandler
	SettingHandler  SettingHandler
	SupplierHandler SupplierHandler
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
