package services

import (
	"strings"

	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
)

type RoleServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Role
}

func NewRoleServiceImpl(app *celeritas.Celeritas, repo data.Role) RoleService {
	return &RoleServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *RoleServiceImpl) CreateRole(input dto.CreateRoleDTO) (*dto.RoleResponseDTO, error) {
	data := input.ToRole()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToRoleResponseDTO(*data)

	return &res, nil
}

func (h *RoleServiceImpl) UpdateRole(id int, input dto.UpdateRoleDTO) (*dto.RoleResponseDTO, error) {
	data, _ := h.repo.Get(id)
	input.ToRole(data)

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToRoleResponseDTO(*data)

	return &response, nil
}

func (h *RoleServiceImpl) DeleteRole(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)

		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return errors.ErrRoleStillAssigned
		}

		return errors.ErrInternalServer
	}

	return nil
}

func (h *RoleServiceImpl) GetRole(id int) (*dto.RoleResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToRoleResponseDTO(*data)

	return &response, nil
}

func (h *RoleServiceImpl) GetRoleList() ([]dto.RoleResponseDTO, error) {
	data, err := h.repo.GetAll(nil)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}
	response := dto.ToRoleListResponseDTO(data)

	return response, nil
}
