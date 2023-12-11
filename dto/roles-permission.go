package dto

import (
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
)

type RolesPermissionDTO struct {
	PermissionID int  `json:"permission_id"`
	RoleID       int  `json:"role_id"`
	CanCreate    bool `json:"create"`
	CanRead      bool `json:"read"`
	CanUpdate    bool `json:"update"`
	CanDelete    bool `json:"delete"`
}

type RolesPermissionResponseDTO struct {
	ID           int       `json:"id"`
	PermissionID int       `json:"permission_id"`
	RoleID       int       `json:"role_id"`
	CanCreate    bool      `json:"create"`
	CanRead      bool      `json:"read"`
	CanUpdate    bool      `json:"update"`
	CanDelete    bool      `json:"delete"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (dto RolesPermissionDTO) ToRolesPermission() *data.RolesPermission {
	return &data.RolesPermission{
		RoleID:       dto.RoleID,
		PermissionID: dto.PermissionID,
		CanCreate:    dto.CanCreate,
		CanUpdate:    dto.CanUpdate,
		CanDelete:    dto.CanDelete,
		CanRead:      dto.CanRead,
	}
}

func ToRolesPermissionResponseDTO(data data.RolesPermission) RolesPermissionResponseDTO {
	return RolesPermissionResponseDTO{
		ID:           data.ID,
		RoleID:       data.RoleID,
		PermissionID: data.PermissionID,
		CanCreate:    data.CanCreate,
		CanUpdate:    data.CanUpdate,
		CanDelete:    data.CanDelete,
		CanRead:      data.CanRead,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func ToRolesPermissionListResponseDTO(rolespermissions []*data.RolesPermission) []RolesPermissionResponseDTO {
	dtoList := make([]RolesPermissionResponseDTO, len(rolespermissions))
	for i, x := range rolespermissions {
		dtoList[i] = ToRolesPermissionResponseDTO(*x)
	}
	return dtoList
}
