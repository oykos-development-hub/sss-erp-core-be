package services

import (
	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type CustomerSupportServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.CustomerSupport
}

func NewCustomerSupportServiceImpl(app *celeritas.Celeritas, repo data.CustomerSupport) CustomerSupportService {
	return &CustomerSupportServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *CustomerSupportServiceImpl) CreateCustomerSupport(input dto.CustomerSupportDTO) (*dto.CustomerSupportResponseDTO, error) {
	data := input.ToCustomerSupport()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToCustomerSupportResponseDTO(*data)

	return &res, nil
}

func (h *CustomerSupportServiceImpl) UpdateCustomerSupport(id int, input dto.CustomerSupportDTO) (*dto.CustomerSupportResponseDTO, error) {
	data := input.ToCustomerSupport()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToCustomerSupportResponseDTO(*data)

	return &response, nil
}

func (h *CustomerSupportServiceImpl) DeleteCustomerSupport(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *CustomerSupportServiceImpl) GetCustomerSupport(id int) (*dto.CustomerSupportResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToCustomerSupportResponseDTO(*data)

	return &response, nil
}

func (h *CustomerSupportServiceImpl) GetCustomerSupportList(filter dto.CustomerSupportFilterDTO) ([]dto.CustomerSupportResponseDTO, *uint64, error) {
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
	response := dto.ToCustomerSupportListResponseDTO(data)

	return response, total, nil
}
