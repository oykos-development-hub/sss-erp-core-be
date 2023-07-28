package services

import (
	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
)

type UserAccountLogServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.UserAccountLog
}

func NewUserAccountLogServiceImpl(app *celeritas.Celeritas, repo data.UserAccountLog) UserAccountLogService {
	return &UserAccountLogServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *UserAccountLogServiceImpl) CreateUserAccountLog(input dto.UserAccountLogDTO) (*dto.UserAccountLogResponseDTO, error) {
	data := input.ToUserAccountLog()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToUserAccountLogResponseDTO(*data)

	return &res, nil
}

func (h *UserAccountLogServiceImpl) DeleteUserAccountLog(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}
