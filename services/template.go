package services

import (
	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/core-ms-api/pkg/errors"

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
		return nil, newErrors.Wrap(err, "repo template create")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo template get")
	}

	res := dto.ToTemplateResponseDTO(*data)

	return &res, nil
}

func (h *TemplateServiceImpl) UpdateTemplate(id int, input dto.TemplateDTO) (*dto.TemplateResponseDTO, error) {
	data := input.ToTemplate()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo template update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo template get")
	}

	response := dto.ToTemplateResponseDTO(*data)

	return &response, nil
}

func (h *TemplateServiceImpl) DeleteTemplate(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		return newErrors.Wrap(err, "repo template delete")
	}

	return nil
}

func (h *TemplateServiceImpl) GetTemplate(id int) (*dto.TemplateResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo template get")
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
		return nil, nil, newErrors.Wrap(err, "repo template get all")
	}
	response := dto.ToTemplateListResponseDTO(data)

	return response, total, nil
}
