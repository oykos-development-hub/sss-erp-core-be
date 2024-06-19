package services

import (
	"context"

	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/core-ms-api/pkg/errors"

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

func (h *RoleServiceImpl) CreateRole(ctx context.Context, input dto.CreateRoleDTO) (*dto.RoleResponseDTO, error) {
	data := input.ToRole()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo role insert")

	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo role get")
	}

	res := dto.ToRoleResponseDTO(*data)

	return &res, nil
}

func (h *RoleServiceImpl) UpdateRole(ctx context.Context, id int, input dto.CreateRoleDTO) (*dto.RoleResponseDTO, error) {
	data := input.ToRole()
	data.ID = id

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo role update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo role get")
	}

	response := dto.ToRoleResponseDTO(*data)

	return &response, nil
}

func (h *RoleServiceImpl) DeleteRole(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {

		return newErrors.Wrap(err, "repo role delete")
	}

	return nil
}

func (h *RoleServiceImpl) GetRole(id int) (*dto.RoleResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo role get")
	}
	response := dto.ToRoleResponseDTO(*data)

	return &response, nil
}

func (h *RoleServiceImpl) GetRoleList() ([]dto.RoleResponseDTO, error) {
	data, err := h.repo.GetAll(nil)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo role get all")
	}
	response := dto.ToRoleListResponseDTO(data)

	return response, nil
}
