package services

import (
	"context"

	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"

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
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToTemplateItemResponseDTO(*data)

	return &res, nil
}

func (h *TemplateItemServiceImpl) UpdateTemplateItem(ctx context.Context, id int, input dto.TemplateItemDTO) (*dto.TemplateItemResponseDTO, error) {
	data := input.ToTemplateItem()
	data.ID = id

	err := h.repo.Update(ctx, *data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToTemplateItemResponseDTO(*data)

	return &response, nil
}

func (h *TemplateItemServiceImpl) DeleteTemplateItem(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *TemplateItemServiceImpl) GetTemplateItem(id int) (*dto.TemplateItemResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
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
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToTemplateItemListResponseDTO(data)

	return response, total, nil
}
