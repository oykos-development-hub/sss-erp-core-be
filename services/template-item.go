package services

import (
	"context"

	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/core-ms-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type TemplateItemServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.TemplateItem
}

func NewTemplateItemServiceImpl(app *celeritas.Celeritas, repo data.TemplateItem) TemplateItemService {
	return &TemplateItemServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *TemplateItemServiceImpl) CreateTemplateItem(ctx context.Context, input dto.TemplateItemDTO) (*dto.TemplateItemResponseDTO, error) {
	data := input.ToTemplateItem()

	id, err := h.repo.Insert(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo template item create")
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo template item get")
	}

	res := dto.ToTemplateItemResponseDTO(*data)

	return &res, nil
}

func (h *TemplateItemServiceImpl) UpdateTemplateItem(ctx context.Context, id int, input dto.TemplateItemDTO) (*dto.TemplateItemResponseDTO, error) {
	data := input.ToTemplateItem()
	data.ID = id

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo template item update")
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo template item get")
	}

	response := dto.ToTemplateItemResponseDTO(*data)

	return &response, nil
}

func (h *TemplateItemServiceImpl) DeleteTemplateItem(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {
		return newErrors.Wrap(err, "repo template item delete")
	}

	return nil
}

func (h *TemplateItemServiceImpl) GetTemplateItem(id int) (*dto.TemplateItemResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo template item get")
	}
	response := dto.ToTemplateItemResponseDTO(*data)

	return &response, nil
}

func (h *TemplateItemServiceImpl) GetTemplateItemList(filter dto.TemplateItemFilterDTO) ([]dto.TemplateItemResponseDTO, *uint64, error) {
	conditionAndExp := &up.AndExpr{}
	var orders []interface{}

	if filter.TemplateID != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"template_id": *filter.TemplateID})
	}

	if filter.OrganizationUnitID != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"organization_unit_id": *filter.OrganizationUnitID})
	}

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
		return nil, nil, newErrors.Wrap(err, "repo template item get all")
	}
	response := dto.ToTemplateItemListResponseDTO(data)

	return response, total, nil
}
