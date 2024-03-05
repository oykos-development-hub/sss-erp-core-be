package services

import (
	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type BankAccountServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.BankAccount
}

func NewBankAccountServiceImpl(app *celeritas.Celeritas, repo data.BankAccount) BankAccountService {
	return &BankAccountServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *BankAccountServiceImpl) CreateBankAccount(input dto.BankAccountDTO) (*dto.BankAccountResponseDTO, error) {
	createData := input.ToBankAccount()

	id, err := h.repo.Insert(data.Upper, *createData)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	createData, err = createData.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToBankAccountResponseDTO(*createData)

	return &res, nil
}

func (h *BankAccountServiceImpl) UpdateBankAccount(id int, input dto.BankAccountDTO) (*dto.BankAccountResponseDTO, error) {
	data := input.ToBankAccount()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToBankAccountResponseDTO(*data)

	return &response, nil
}

func (h *BankAccountServiceImpl) DeleteBankAccount(title string) error {
	err := h.repo.Delete(data.Upper, title)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *BankAccountServiceImpl) GetBankAccount(id int) (*dto.BankAccountResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToBankAccountResponseDTO(*data)

	return &response, nil
}

func (h *BankAccountServiceImpl) GetBankAccountList(filter dto.BankAccountFilterDTO) ([]dto.BankAccountResponseDTO, *uint64, error) {
	conditionAndExp := &up.AndExpr{}
	var orders []interface{}

	if filter.SupplierID != nil {
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"supplier_id": *filter.SupplierID})
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
	response := dto.ToBankAccountListResponseDTO(data)

	return response, total, nil
}
