package dto

import (
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
)

type CreateRoleDTO struct {
	Title        string `json:"title" validate:"required"`
	Abbreviation string `json:"abbreviation" validate:"required"`
}

type UpdateRoleDTO struct {
	Title        *string `json:"title" validate:"omitempty"`
	Abbreviation *string `json:"abbreviation" validate:"omitempty"`
}

type RoleResponseDTO struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Abbreviation string    `json:"abbreviation"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (dto CreateRoleDTO) ToRole() *data.Role {
	return &data.Role{
		Title:        dto.Title,
		Abbreviation: dto.Abbreviation,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (dto UpdateRoleDTO) ToRole(data *data.Role) {
	if dto.Title != nil {
		data.Title = *dto.Title
	}
	if dto.Abbreviation != nil {
		data.Abbreviation = *dto.Abbreviation
	}

	data.UpdatedAt = time.Now()
}

func ToRoleResponseDTO(data data.Role) RoleResponseDTO {
	return RoleResponseDTO{
		ID:           data.ID,
		Title:        data.Title,
		Abbreviation: data.Abbreviation,
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
