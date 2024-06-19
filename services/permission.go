package services

import (
	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/core-ms-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
)

type PermissionServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Permission
}

func NewPermissionServiceImpl(app *celeritas.Celeritas, repo data.Permission) PermissionService {
	return &PermissionServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *PermissionServiceImpl) CreatePermission(input dto.PermissionDTO) (*dto.PermissionResponseDTO, error) {
	data := input.ToPermission()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo permission insert")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo permission get")
	}

	res := dto.ToPermissionResponseDTO(*data)

	return &res, nil
}

func (h *PermissionServiceImpl) UpdatePermission(id int, input dto.PermissionDTO) (*dto.PermissionResponseDTO, error) {
	data := input.ToPermission()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo permission update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo permission get")
	}

	response := dto.ToPermissionResponseDTO(*data)

	return &response, nil
}

func (h *PermissionServiceImpl) DeletePermission(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		return newErrors.Wrap(err, "repo permission delete")
	}

	return nil
}

func (h *PermissionServiceImpl) GetPermission(id int) (*dto.PermissionResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo permission get")
	}
	response := dto.ToPermissionResponseDTO(*data)

	return &response, nil
}

func (h *PermissionServiceImpl) GetPermissionList() ([]dto.PermissionResponseDTO, error) {
	data, err := h.repo.GetAll(nil)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo permission get all")
	}
	response := dto.ToPermissionListResponseDTO(data)

	return response, nil
}

func (h *PermissionServiceImpl) GetPermissionListForRole(roleID int) ([]dto.PermissionWithRolesResponseDTO, error) {
	data, err := h.repo.GetAllPermissionOfRole(roleID)
	data[0].CanRead = true
	if err != nil {
		return nil, newErrors.Wrap(err, "repo permission get permission list for role")
	}
	response := dto.ToPermissionListWithRoleResponseDTO(data)

	return response, nil
}
