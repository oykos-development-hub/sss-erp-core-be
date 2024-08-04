package services

import (
	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type ListOfParameterServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.ListOfParameter
}

func NewListOfParameterServiceImpl(app *celeritas.Celeritas, repo data.ListOfParameter) ListOfParameterService {
	return &ListOfParameterServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *ListOfParameterServiceImpl) CreateListOfParameter(input dto.ListOfParameterDTO) (*dto.ListOfParameterResponseDTO, error) {
	data := input.ToListOfParameter()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToListOfParameterResponseDTO(*data)

	return &res, nil
}

func (h *ListOfParameterServiceImpl) UpdateListOfParameter(id int, input dto.ListOfParameterDTO) (*dto.ListOfParameterResponseDTO, error) {
	data := input.ToListOfParameter()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToListOfParameterResponseDTO(*data)

	return &response, nil
}

func (h *ListOfParameterServiceImpl) DeleteListOfParameter(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *ListOfParameterServiceImpl) GetListOfParameter(id int) (*dto.ListOfParameterResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToListOfParameterResponseDTO(*data)

	return &response, nil
}

func (h *ListOfParameterServiceImpl) GetListOfParameterList(filter dto.ListOfParameterFilterDTO) ([]dto.ListOfParameterResponseDTO, *uint64, error) {
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
	response := dto.ToListOfParameterListResponseDTO(data)

	return response, total, nil
}
