package services

import (
	"fmt"

	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type SettingServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Setting
}

func NewSettingServiceImpl(app *celeritas.Celeritas, repo data.Setting) SettingService {
	return &SettingServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *SettingServiceImpl) CreateSetting(input dto.SettingDTO) (*dto.SettingResponseDTO, error) {
	data := input.ToSetting()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToSettingResponseDTO(*data)

	return &res, nil
}

func (h *SettingServiceImpl) UpdateSetting(id int, input dto.SettingDTO) (*dto.SettingResponseDTO, error) {
	data := input.ToSetting()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}

	response := dto.ToSettingResponseDTO(*data)

	return &response, nil
}

func (h *SettingServiceImpl) DeleteSetting(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *SettingServiceImpl) GetSetting(id int) (*dto.SettingResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(id, err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToSettingResponseDTO(*data)

	return &response, nil
}

func (h *SettingServiceImpl) GetSettingList(data dto.GetSettingsDTO) ([]dto.SettingResponseDTO, *uint64, error) {
	conditionAndExp := &up.AndExpr{}
	conditionAndExp = up.And(conditionAndExp, &up.Cond{"entity": data.Entity})
	if data.Search != nil && *data.Search != "" {
		likeCondition := fmt.Sprintf("%%%s%%", *data.Search)
		cond := up.And(
			up.Or(
				up.Cond{"title ILIKE": likeCondition},
				up.Cond{"description ILIKE": likeCondition},
				up.Cond{"abbreviation ILIKE": likeCondition},
			),
		)
		conditionAndExp = up.And(conditionAndExp, cond)
	}

	if data.Value != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"value": *data.Value})
	}

	if data.ParentID != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"parent_id": *data.ParentID})
	} else {
		cond := up.And(
			up.Or(
				up.Cond{"parent_id is": nil},
				up.Cond{"parent_id": 0},
			),
		)
		conditionAndExp = up.And(conditionAndExp, cond)
	}

	res, total, err := h.repo.GetAll(data.Page, data.Size, conditionAndExp)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToSettingListResponseDTO(res)

	return response, total, nil
}
