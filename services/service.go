package services

import (
	"context"

	jwtdto "github.com/oykos-development-hub/celeritas/jwt/dto"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
)

type BaseService interface {
	RandomString(n int) string
	Encrypt(text string) (string, error)
	Decrypt(crypto string) (string, error)
}

type UserService interface {
	CreateUser(ctx context.Context, input dto.UserRegistrationDTO) (*dto.UserResponseDTO, error)
	UpdateUser(ctx context.Context, id int, input dto.UserUpdateDTO) (*dto.UserResponseDTO, error)
	GetUser(id int) (*dto.UserResponseDTO, error)
	GetUserList(data dto.GetUserListDTO) ([]dto.UserResponseDTO, *uint64, error)
	DeleteUser(ctx context.Context, id int) error
}

type AuthService interface {
	Login(loginInput dto.LoginInput) (*dto.LoginResponse, error)
	RefreshToken(userId int, refreshToken string, iat string) (*jwtdto.Token, error)
	Logout(userId int) error
	ForgotPassword(input dto.ForgotPassword) error
	ResetPasswordVerify(email, token string) (*dto.ResetPasswordVerifyResponse, error)
	ResetPassword(input dto.ResetPassword) error
	ValidatePin(id int, data dto.ValidatePinInput) error
}

type RoleService interface {
	CreateRole(ctx context.Context, input dto.CreateRoleDTO) (*dto.RoleResponseDTO, error)
	UpdateRole(ctx context.Context, id int, input dto.CreateRoleDTO) (*dto.RoleResponseDTO, error)
	DeleteRole(ctx context.Context, id int) error
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
	CreateAccountList(ctx context.Context, input []dto.AccountDTO) ([]dto.AccountResponseDTO, error)
	DeleteAccount(ctx context.Context, id int) error
	GetAccount(id int) (*dto.AccountResponseDTO, error)
	GetAccountList(input dto.GetAccountsFilter) ([]dto.AccountResponseDTO, int, error)
}

type NotificationService interface {
	CreateNotification(input dto.NotificationDTO) (*dto.NotificationResponseDTO, error)
	UpdateNotification(id int, input dto.NotificationDTO) (*dto.NotificationResponseDTO, error)
	DeleteNotification(id int) error
	GetNotification(id int) (*dto.NotificationResponseDTO, error)
	GetNotificationList(input dto.GetNotificationListInput) ([]dto.NotificationResponseDTO, *uint64, error)
}

type RolesPermissionService interface {
	SyncPermissions(roleID int, input []dto.RolesPermissionDTO) ([]dto.RolesPermissionResponseDTO, error)
}

type PermissionService interface {
	CreatePermission(input dto.PermissionDTO) (*dto.PermissionResponseDTO, error)
	UpdatePermission(id int, input dto.PermissionDTO) (*dto.PermissionResponseDTO, error)
	DeletePermission(id int) error
	GetPermission(id int) (*dto.PermissionResponseDTO, error)
	GetPermissionList() ([]dto.PermissionResponseDTO, error)
	GetPermissionListForRole(roleID int) ([]dto.PermissionWithRolesResponseDTO, error)
}

type BankAccountService interface {
	CreateBankAccount(input dto.BankAccountDTO) (*dto.BankAccountResponseDTO, error)
	UpdateBankAccount(id int, input dto.BankAccountDTO) (*dto.BankAccountResponseDTO, error)
	DeleteBankAccount(title string) error
	GetBankAccount(id int) (*dto.BankAccountResponseDTO, error)
	GetBankAccountList(filter dto.BankAccountFilterDTO) ([]dto.BankAccountResponseDTO, *uint64, error)
}

type LogService interface {
	CreateLog(input dto.LogDTO) (*dto.LogResponseDTO, error)
	UpdateLog(id int, input dto.LogDTO) (*dto.LogResponseDTO, error)
	DeleteLog(id int) error
	GetLog(id int) (*dto.LogResponseDTO, error)
	GetLogList(filter dto.LogFilterDTO) ([]dto.LogResponseDTO, *uint64, error)
}

type TemplateService interface {
	CreateTemplate(input dto.TemplateDTO) (*dto.TemplateResponseDTO, error)
	UpdateTemplate(id int, input dto.TemplateDTO) (*dto.TemplateResponseDTO, error)
	DeleteTemplate(id int) error
	GetTemplate(id int) (*dto.TemplateResponseDTO, error)
	GetTemplateList(filter dto.TemplateFilterDTO) ([]dto.TemplateResponseDTO, *uint64, error)
}

type TemplateItemService interface {
	CreateTemplateItem(ctx context.Context, input dto.TemplateItemDTO) (*dto.TemplateItemResponseDTO, error)
	UpdateTemplateItem(ctx context.Context, id int, input dto.TemplateItemDTO) (*dto.TemplateItemResponseDTO, error)
	DeleteTemplateItem(ctx context.Context, id int) error
	GetTemplateItem(id int) (*dto.TemplateItemResponseDTO, error)
	GetTemplateItemList(filter dto.TemplateItemFilterDTO) ([]dto.TemplateItemResponseDTO, *uint64, error)
}
