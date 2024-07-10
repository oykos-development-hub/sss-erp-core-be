package services

import (
	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/core-ms-api/pkg/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type BffErrorLogServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.BffErrorLog
}

func NewBffErrorLogServiceImpl(app *celeritas.Celeritas, repo data.BffErrorLog) BffErrorLogService {
	return &BffErrorLogServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *BffErrorLogServiceImpl) CreateBffErrorLog(input dto.BffErrorLogDTO) (*dto.BffErrorLogResponseDTO, error) {
	dataToInsert := input.ToBffErrorLog()

	var id int

	err := data.Upper.Tx(func(tx up.Session) error {
		var err error
		_, err = h.repo.Insert(tx, *dataToInsert)
		if err != nil {
			return newErrors.Wrap(err, "repo insert")
		}

		return nil
	})

	if err != nil {
		return nil, newErrors.Wrap(err, "upper tx")
	}

	dataToInsert, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo log get")
	}

	res := dto.ToBffErrorLogResponseDTO(*dataToInsert)

	return &res, nil

}

func (h *BffErrorLogServiceImpl) UpdateBffErrorLog(id int, input dto.BffErrorLogDTO) (*dto.BffErrorLogResponseDTO, error) {
	dataToInsert := input.ToBffErrorLog()
	dataToInsert.ID = id

	err := data.Upper.Tx(func(tx up.Session) error {
		err := h.repo.Update(tx, *dataToInsert)
		if err != nil {
			return newErrors.Wrap(err, "repo update")
		}
		return nil
	})
	if err != nil {
		return nil, newErrors.Wrap(err, "upper tx")
	}

	dataToInsert, err = h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo get")
	}

	response := dto.ToBffErrorLogResponseDTO(*dataToInsert)

	return &response, nil
}

func (h *BffErrorLogServiceImpl) DeleteBffErrorLog(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		return newErrors.Wrap(err, "repo delete")
	}

	return nil
}

func (h *BffErrorLogServiceImpl) GetBffErrorLog(id int) (*dto.BffErrorLogResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo get")
	}
	response := dto.ToBffErrorLogResponseDTO(*data)

	return &response, nil
}

func (h *BffErrorLogServiceImpl) GetBffErrorLogList(filter dto.BffErrorLogFilterDTO) ([]dto.BffErrorLogResponseDTO, *uint64, error) {
	conditionAndExp := &up.AndExpr{}
	var orders []interface{}

	// example of making conditions
	if filter.Entity != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"entity": *filter.Entity})
	}

	if filter.DateOfStart != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"created_at > ": *filter.DateOfStart})
	}

	if filter.DateOfEnd != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"created_at < ": *filter.DateOfEnd})
	}

	/*if filter.SortByTitle != nil {
		if *filter.SortByTitle == "asc" {
			orders = append(orders, "-title")
		} else {
			orders = append(orders, "title")
		}
	}*/

	orders = append(orders, "-created_at")

	data, total, err := h.repo.GetAll(filter.Page, filter.Size, conditionAndExp, orders)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "repo get all")
	}
	response := dto.ToBffErrorLogListResponseDTO(data)

	return response, total, nil
}
