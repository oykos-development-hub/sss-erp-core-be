package services

import (
	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	"github.com/upper/db/v4"
	up "github.com/upper/db/v4"
)

type AccountServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Account
}

func NewAccountServiceImpl(app *celeritas.Celeritas, repo data.Account) AccountService {
	return &AccountServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *AccountServiceImpl) CreateAccount(input dto.AccountDTO) (*dto.AccountResponseDTO, error) {
	data := input.ToAccount()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToAccountResponseDTO(*data)

	return &res, nil
}

func (h *AccountServiceImpl) UpdateAccount(id int, input dto.AccountDTO) (*dto.AccountResponseDTO, error) {
	data := input.ToAccount()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToAccountResponseDTO(*data)

	return &response, nil
}

func (h *AccountServiceImpl) DeleteAccount(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *AccountServiceImpl) GetAccount(id int) (*dto.AccountResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToAccountResponseDTO(*data)

	return &response, nil
}

func (h *AccountServiceImpl) GetAccountList(input dto.GetAccountsFilter) ([]dto.AccountResponseDTO, int, error) {
	var cond []up.LogicalExpr
	var combinedCond *up.AndExpr
	if input.Search != nil && *input.Search != "" {
		search := "%" + *input.Search + "%"
		searchCond := up.Or(
			db.Cond{"title ILIKE": search},
			db.Cond{"serial_number ILIKE": search},
		)
		cond = append(cond, searchCond)
	}
	if input.Search != nil && *input.Search != "" {
		cond = append(cond, up.And(db.Cond{"id": input.ID}))
	}
	if len(cond) > 0 {
		combinedCond = up.And(cond...)
	}

	data, total, err := h.repo.GetAll(input.Page, input.Size, combinedCond)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, -1, errors.ErrInternalServer
	}
	response := dto.ToAccountListResponseDTO(data)

	return response, total, nil
}
