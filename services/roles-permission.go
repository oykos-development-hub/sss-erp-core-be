package services

import (
	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"

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

func (h *RolesPermissionServiceImpl) CreateRolesPermission(input dto.RolesPermissionDTO) (*dto.RolesPermissionResponseDTO, error) {
	data := input.ToRolesPermission()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToRolesPermissionResponseDTO(*data)

	return &res, nil
}

func (h *RolesPermissionServiceImpl) UpdateRolesPermission(id int, input dto.RolesPermissionDTO) (*dto.RolesPermissionResponseDTO, error) {
	data := input.ToRolesPermission()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToRolesPermissionResponseDTO(*data)

	return &response, nil
}

func (h *RolesPermissionServiceImpl) DeleteRolesPermission(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *RolesPermissionServiceImpl) GetRolesPermission(id int) (*dto.RolesPermissionResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToRolesPermissionResponseDTO(*data)

	return &response, nil
}

func (h *RolesPermissionServiceImpl) GetRolesPermissionList() ([]dto.RolesPermissionResponseDTO, error) {
	data, err := h.repo.GetAll(nil)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}
	response := dto.ToRolesPermissionListResponseDTO(data)

	return response, nil
}
