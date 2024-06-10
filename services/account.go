package services

import (
	"context"
	"fmt"
	"sort"

	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
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

type BySerialNumberLength []dto.AccountDTO

func (a BySerialNumberLength) Len() int {
	return len(a)
}

func (a BySerialNumberLength) Less(i, j int) bool {
	return len(a[i].SerialNumber) < len(a[j].SerialNumber)
}

func (a BySerialNumberLength) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (h *AccountServiceImpl) CreateAccountList(ctx context.Context, input []dto.AccountDTO) ([]dto.AccountResponseDTO, error) {
	var latestCountVersion int
	counts, total, err := h.GetAccountList(dto.GetAccountsFilter{})
	if err != nil {
		return nil, errors.ErrInternalServer
	}
	if total > 0 {
		latestCountVersion = counts[0].Version
	}

	sort.Sort(BySerialNumberLength(input))

	var accountData []*data.Account
	for _, account := range input {
		data := account.ToAccount()
		data.Version = latestCountVersion + 1

		id, err := h.repo.Insert(ctx, *data)
		if err != nil {
			return nil, errors.ErrInternalServer
		}

		data, err = data.Get(id)
		if err != nil {
			return nil, errors.ErrInternalServer
		}
		accountData = append(accountData, data)
	}

	res := dto.ToAccountListResponseDTO(accountData)

	return res, nil
}

func (h *AccountServiceImpl) DeleteAccount(ctx context.Context, id int) error {
	err := h.repo.Delete(ctx, id)
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
	conditionAndExp := &up.AndExpr{}

	if input.Search != nil && *input.Search != "" {
		search := "%" + *input.Search + "%"
		searchCond := up.Or(
			up.Cond{"title ILIKE": search},
			up.Cond{"serial_number ILIKE": search},
		)
		conditionAndExp = up.And(conditionAndExp, searchCond)
	}
	if input.ID != nil && *input.ID != 0 {
		conditionAndExp = up.And(conditionAndExp, up.Cond{"id": input.ID})
	}
	if input.SerialNumber != nil && *input.SerialNumber != "" {
		conditionAndExp = up.And(conditionAndExp, up.Cond{"serial_number": up.Like(*input.SerialNumber)})
	}
	if input.Title != nil && *input.Title != "" {
		conditionAndExp = up.And(conditionAndExp, up.Cond{"title": up.Like(*input.Title)})
	}
	if input.Version != nil && *input.Version != 0 {
		conditionAndExp = up.And(conditionAndExp, up.Cond{"version": input.Version})
	} else {
		latestVersionSubquery := `SELECT id FROM accounts WHERE version = (SELECT MAX(version) FROM accounts)`
		conditionAndExp = up.And(conditionAndExp, up.Raw(fmt.Sprintf("id IN (%s)", latestVersionSubquery)))
		//deb := fmt.Sprintf("id IN (%s)", latestVersionSubquery)
	}
	if input.Leaf {
		leafSubquery := `SELECT a.id FROM accounts a LEFT JOIN accounts b ON a.id = b.parent_id WHERE b.id IS NULL`
		conditionAndExp = up.And(conditionAndExp, up.Raw(fmt.Sprintf("id IN (%s)", leafSubquery)))
	}
	if input.TopMost {
		conditionAndExp = up.And(conditionAndExp, up.Cond{"parent_id": nil})
	}

	data, total, err := h.repo.GetAll(input.Page, input.Size, conditionAndExp)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, -1, errors.ErrInternalServer
	}
	response := dto.ToAccountListResponseDTO(data)

	return response, total, nil
}
