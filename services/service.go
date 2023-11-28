package services

import (
	jwtdto "github.com/oykos-development-hub/celeritas/jwt/dto"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
)

type BaseService interface {
	RandomString(n int) string
	Encrypt(text string) (string, error)
	Decrypt(crypto string) (string, error)
}

type UserService interface {
	CreateUser(input dto.UserRegistrationDTO) (*dto.UserResponseDTO, error)
	UpdateUser(id int, input dto.UserUpdateDTO) (*dto.UserResponseDTO, error)
	GetUser(id int) (*dto.UserResponseDTO, error)
	GetUserList(data dto.GetUserListDTO) ([]dto.UserResponseDTO, *uint64, error)
	DeleteUser(id int) error
}

type AuthService interface {
	Login(loginInput dto.LoginInput) (*dto.LoginResponse, error)
	RefreshToken(userId int, refreshToken string, iat string) (*jwtdto.Token, error)
	Logout(userId int) error
	ForgotPassword(input dto.ForgotPassword) error
	ForgotPasswordV2(input dto.ForgotPassword) error
	ResetPasswordVerify(email, token string) (bool, error)
	ResetPassword(input dto.ResetPassword) error
	ValidatePin(id int, data dto.ValidatePinInput) error
}

type RoleService interface {
	CreateRole(input dto.CreateRoleDTO) (*dto.RoleResponseDTO, error)
	UpdateRole(id int, input dto.UpdateRoleDTO) (*dto.RoleResponseDTO, error)
	DeleteRole(id int) error
	GetRole(id int) (*dto.RoleResponseDTO, error)
	GetRoleList() ([]dto.RoleResponseDTO, error)
}

type SettingService interface {
	CreateSetting(input dto.SettingDTO) (*dto.SettingResponseDTO, error)
	UpdateSetting(id int, input dto.SettingDTO) (*dto.SettingResponseDTO, error)
	DeleteSetting(id int) error
	GetSetting(id int) (*dto.SettingResponseDTO, error)
	GetSettingList(dto.GetSettingsDTO) ([]dto.SettingResponseDTO, *uint64, error)
}

type UserAccountLogService interface {
	CreateUserAccountLog(input dto.UserAccountLogDTO) (*dto.UserAccountLogResponseDTO, error)
	DeleteUserAccountLog(id int) error
}

type SupplierService interface {
	CreateSupplier(input dto.SupplierDTO) (*dto.SupplierResponseDTO, error)
	UpdateSupplier(id int, input dto.SupplierDTO) (*dto.SupplierResponseDTO, error)
	DeleteSupplier(id int) error
	GetSupplier(id int) (*dto.SupplierResponseDTO, error)
	GetSupplierList(input dto.GetSupplierListInput) ([]dto.SupplierResponseDTO, *uint64, error)
}

type AccountService interface {
	CreateAccount(input dto.AccountDTO) (*dto.AccountResponseDTO, error)
	UpdateAccount(id int, input dto.AccountDTO) (*dto.AccountResponseDTO, error)
	DeleteAccount(id int) error
	GetAccount(id int) (*dto.AccountResponseDTO, error)
	GetAccountList(input dto.GetAccountsFilter) ([]dto.AccountResponseDTO, int, error)
}

type NotificationService interface {
	CreateNotification(input dto.NotificationDTO) (*dto.NotificationResponseDTO, error)
	UpdateNotification(id int, input dto.NotificationDTO) (*dto.NotificationResponseDTO, error)
	DeleteNotification(id int) error
	GetNotification(id int) (*dto.NotificationResponseDTO, error)
	GetNotificationList() ([]dto.NotificationResponseDTO, error)
}
