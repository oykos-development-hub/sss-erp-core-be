package dto

import (
	"gitlab.sudovi.me/erp/core-ms-api/data"
)

type GetPermissionListInput struct {
	RoleID int `json:"role_id"`
}

type PermissionDTO struct {
	Title    string `json:"title"`
	ParentID *int   `json:"parent_id"`
}

type PermissionResponseDTO struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Path     string `json:"route"`
	ParentID *int   `json:"parent_id"`
}

type PermissionWithRolesResponseDTO struct {
	PermissionResponseDTO
	CanCreate bool `json:"create"`
	CanRead   bool `json:"read"`
	CanUpdate bool `json:"update"`
	CanDelete bool `json:"delete"`
}

func (dto PermissionDTO) ToPermission() *data.Permission {
	return &data.Permission{
		Title:    dto.Title,
		ParentID: dto.ParentID,
	}
}

func ToPermissionResponseDTO(data data.Permission) PermissionResponseDTO {
	return PermissionResponseDTO{
		ID:       data.ID,
		Title:    data.Title,
		ParentID: data.ParentID,
		Path:     data.Path,
	}
}

func ToPermissionListResponseDTO(permissions []*data.Permission) []PermissionResponseDTO {
	dtoList := make([]PermissionResponseDTO, len(permissions))
	for i, x := range permissions {
		dtoList[i] = ToPermissionResponseDTO(*x)
	}
	return dtoList
}

func ToPermissionWithRolesResponseDTO(data data.PermissionWithRoles) PermissionWithRolesResponseDTO {
	return PermissionWithRolesResponseDTO{
		PermissionResponseDTO: ToPermissionResponseDTO(data.Permission),
		CanCreate:             data.CanCreate,
		CanUpdate:             data.CanUpdate,
		CanDelete:             data.CanDelete,
		CanRead:               data.CanRead,
	}
}

func ToPermissionListWithRoleResponseDTO(permissions []*data.PermissionWithRoles) []PermissionWithRolesResponseDTO {
	dtoList := make([]PermissionWithRolesResponseDTO, len(permissions))
	for i, x := range permissions {
		dtoList[i] = ToPermissionWithRolesResponseDTO(*x)
	}
	return dtoList
}
