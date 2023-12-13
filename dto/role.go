package dto

import (
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
)

type CreateRoleDTO struct {
	Title        string `json:"title" validate:"required"`
	Abbreviation string `json:"abbreviation" validate:"required"`
	Active       bool   `json:"active"`
}

type RoleResponseDTO struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Abbreviation string    `json:"abbreviation"`
	Active       bool      `json:"active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (dto CreateRoleDTO) ToRole() *data.Role {
	return &data.Role{
		Title:        dto.Title,
		Abbreviation: dto.Abbreviation,
		Active:       dto.Active,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func ToRoleResponseDTO(data data.Role) RoleResponseDTO {
	return RoleResponseDTO{
		ID:           data.ID,
		Title:        data.Title,
		Abbreviation: data.Abbreviation,
		Active:       data.Active,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func ToRoleListResponseDTO(Roles []*data.Role) []RoleResponseDTO {
	dtoList := make([]RoleResponseDTO, len(Roles))
	for i, x := range Roles {
		dtoList[i] = ToRoleResponseDTO(*x)
	}
	return dtoList
}
