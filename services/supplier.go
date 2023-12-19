package services

import (
	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type SupplierServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Supplier
}

func NewSupplierServiceImpl(app *celeritas.Celeritas, repo data.Supplier) SupplierService {
	return &SupplierServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *SupplierServiceImpl) CreateSupplier(input dto.SupplierDTO) (*dto.SupplierResponseDTO, error) {
	data := input.ToSupplier()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToSupplierResponseDTO(*data)

	return &res, nil
}

func (h *SupplierServiceImpl) UpdateSupplier(id int, input dto.SupplierDTO) (*dto.SupplierResponseDTO, error) {
	data := input.ToSupplier()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToSupplierResponseDTO(*data)

	return &response, nil
}

func (h *SupplierServiceImpl) DeleteSupplier(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *SupplierServiceImpl) GetSupplier(id int) (*dto.SupplierResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToSupplierResponseDTO(*data)

	return &response, nil
}

func (h *SupplierServiceImpl) GetSupplierList(input dto.GetSupplierListInput) ([]dto.SupplierResponseDTO, *uint64, error) {
	var cond []up.LogicalExpr
	var combinedCond *up.AndExpr
	if input.Search != nil {
		search := "%" + *input.Search + "%"
		h.App.InfoLog.Println(search)
		searchCond := up.Or(
			up.Cond{"title ILIKE": search},
			up.Cond{"abbreviation ILIKE": search},
			up.Cond{"address ILIKE": search},
			up.Cond{"description ILIKE": search},
			up.Cond{"official_id ILIKE": search},
		)
		cond = append(cond, searchCond)
	}

	if input.Entity != nil {
		searchCond := up.Cond{"entity": input.Entity}
		cond = append(cond, searchCond)
	} else {
		searchCond := up.Cond{"entity": "supplier"}
		cond = append(cond, searchCond)
	}

	if len(cond) > 0 {
		combinedCond = up.And(cond...)
	}

	data, total, err := h.repo.GetAll(input.Page, input.Size, combinedCond)

	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToSupplierListResponseDTO(data)

	return response, total, nil
}
