package dto

import (
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
)

type UserRegistrationDTO struct {
	FirstName      string  `json:"first_name" validate:"required"`
	LastName       string  `json:"last_name" validate:"required"`
	Email          string  `json:"email" validate:"required"`
	Password       string  `json:"password" validate:"required"`
	SecondaryEmail *string `json:"secondary_email" validate:"omitempty"`
	Phone          string  `json:"phone" validate:"required"`
	Pin            string  `json:"pin" validate:"len=4"`
	Active         *bool   `json:"active" validate:"omitempty,boolean"`
	VerifiedEmail  *bool   `json:"verified_email" validate:"omitempty,boolean"`
	VerifiedPhone  *bool   `json:"verified_phone" validate:"omitempty,boolean"`
	FolderId       *int    `json:"folder_id" validate:"omitempty"`
	RoleId         int     `json:"role_id" validate:"omitempty"`
}

func (dto *UserRegistrationDTO) ToUser() *data.User {
	return &data.User{
		FirstName:      dto.FirstName,
		LastName:       dto.LastName,
		Email:          dto.Email,
		SecondaryEmail: dto.SecondaryEmail,
		Password:       dto.Password,
		Phone:          dto.Phone,
		Pin:            dto.Pin,
		Active:         true,
		VerifiedEmail:  false,
		VerifiedPhone:  false,
		FolderId:       dto.FolderId,
		RoleId:         dto.RoleId,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

type UserUpdateDTO struct {
	FirstName      *string `json:"first_name" validate:"omitempty"`
	LastName       *string `json:"last_name" validate:"omitempty"`
	Email          *string `json:"email" validate:"omitempty,email"`
	Active         *bool   `json:"active" validate:"omitempty,boolean"`
	SecondaryEmail *string `json:"secondary_email" validate:"omitempty,email"`
	Phone          *string `json:"phone" validate:"omitempty"`
	FolderId       *int    `json:"folder_id" validate:"omitempty"`
	RoleId         *int    `json:"role_id" validate:"omitempty"`
}

type ValidatePinInput struct {
	Pin string `json:"pin" validate:"required,len=4"`
}

func (dto *UserUpdateDTO) ToUser(u *data.User) {
	if dto.FirstName != nil {
		u.FirstName = *dto.FirstName
	}
	if dto.LastName != nil {
		u.LastName = *dto.LastName
	}
	if dto.Phone != nil {
		u.Phone = *dto.Phone
	}
	if dto.Email != nil {
		u.Email = *dto.Email
	}
	if dto.Active != nil {
		u.Active = *dto.Active
	}
	if dto.RoleId != nil {
		u.RoleId = *dto.RoleId
	}

	u.SecondaryEmail = dto.SecondaryEmail
	u.FolderId = dto.FolderId

	u.UpdatedAt = time.Now()
}

type UserResponseDTO struct {
	ID             int       `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Active         bool      `json:"active"`
	SecondaryEmail *string   `json:"secondary_email"`
	Phone          string    `json:"phone"`
	VerifiedEmail  bool      `json:"verified_email"`
	VerifiedPhone  bool      `json:"verified_phone"`
	FolderId       *int      `json:"folder_id"`
	RoleId         int       `json:"role_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type GetUserListDTO struct {
	Page     *int    `json:"page" validate:"omitempty"`
	Size     *int    `json:"size" validate:"omitempty"`
	IsActive *bool   `json:"is_active" validate:"omitempty"`
	Email    *string `json:"email" validate:"omitempty"`
}

func ToUserResponseDTO(user data.User) *UserResponseDTO {
	return &UserResponseDTO{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		Active:         user.Active,
		SecondaryEmail: user.SecondaryEmail,
		Phone:          user.Phone,
		VerifiedEmail:  user.VerifiedEmail,
		VerifiedPhone:  user.VerifiedPhone,
		FolderId:       user.FolderId,
		RoleId:         user.RoleId,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
}

func ToUsersResponseDTO(users []*data.User) []UserResponseDTO {
	dtoUsers := make([]UserResponseDTO, len(users))
	for i, user := range users {
		dtoUsers[i] = *ToUserResponseDTO(*user)
	}
	return dtoUsers
}
