package services

import (
	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/core-ms-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
)

type RolesPermissionServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.RolesPermission
}

func NewRolesPermissionServiceImpl(app *celeritas.Celeritas, repo data.RolesPermission) RolesPermissionService {
	return &RolesPermissionServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *RolesPermissionServiceImpl) SyncPermissions(roleID int, input []dto.RolesPermissionDTO) ([]dto.RolesPermissionResponseDTO, error) {
	err := h.repo.DeleteAllPermissionsByRole(roleID)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo role permission sync permissions")
	}

	for _, rolePermission := range input {
		data := rolePermission.ToRolesPermission()

		_, err := h.repo.Insert(*data)
		if err != nil {
			return nil, newErrors.Wrap(err, "repo role permission insert")
		}
	}

	data, err := h.repo.GetAll(nil)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo role permission get all")
	}

	response := dto.ToRolesPermissionListResponseDTO(data)

	return response, nil
}
