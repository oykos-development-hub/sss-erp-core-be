package dto

import jwtdto "github.com/oykos-development-hub/celeritas/jwt/dto"

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ForgotPassword struct {
	Email string `json:"email"`
}

type ResetPasswordVerify struct {
	Email string `json:"email" validate:"required"`
	Token string `json:"token" validate:"required"`
}

type ResetPassword struct {
	Email    string `json:"email" validate:"required"`
	Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	User  UserResponseDTO `json:"user"`
	Token jwtdto.Token    `json:"token"`
}
