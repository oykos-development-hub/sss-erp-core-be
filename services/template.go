package services

import (
	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type TemplateServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Template
}

func NewTemplateServiceImpl(app *celeritas.Celeritas, repo data.Template) TemplateService {
	return &TemplateServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *TemplateServiceImpl) CreateTemplate(input dto.TemplateDTO) (*dto.TemplateResponseDTO, error) {
	data := input.ToTemplate()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToTemplateResponseDTO(*data)

	return &res, nil
}

func (h *TemplateServiceImpl) UpdateTemplate(id int, input dto.TemplateDTO) (*dto.TemplateResponseDTO, error) {
	data := input.ToTemplate()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToTemplateResponseDTO(*data)

	return &response, nil
}

func (h *TemplateServiceImpl) DeleteTemplate(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *TemplateServiceImpl) GetTemplate(id int) (*dto.TemplateResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToTemplateResponseDTO(*data)

	return &response, nil
}

func (h *TemplateServiceImpl) GetTemplateList(filter dto.TemplateFilterDTO) ([]dto.TemplateResponseDTO, *uint64, error) {
	conditionAndExp := &up.AndExpr{}
	var orders []interface{}

	// example of making conditions
	// if filter.Year != nil {
	// 	conditionAndExp = up.And(conditionAndExp, &up.Cond{"year": *filter.Year})
	// }

	if filter.SortByTitle != nil {
		if *filter.SortByTitle == "asc" {
			orders = append(orders, "-title")
		} else {
			orders = append(orders, "title")
		}
	}

	orders = append(orders, "-created_at")
	

	data, total, err := h.repo.GetAll(filter.Page, filter.Size, conditionAndExp, orders)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToTemplateListResponseDTO(data)

	return response, total, nil
}
