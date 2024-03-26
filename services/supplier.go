package services

import (
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type SupplierServiceImpl struct {
	App             *celeritas.Celeritas
	repo            data.Supplier
	bankAccountRepo data.BankAccount
}

func NewSupplierServiceImpl(app *celeritas.Celeritas, repo data.Supplier, bankAccountRepo data.BankAccount) SupplierService {
	return &SupplierServiceImpl{
		App:             app,
		repo:            repo,
		bankAccountRepo: bankAccountRepo,
	}
}

func (h *SupplierServiceImpl) CreateSupplier(input dto.SupplierDTO) (*dto.SupplierResponseDTO, error) {
	supplier := input.ToSupplier()
	var id int

	err := data.Upper.Tx(func(tx up.Session) error {
		if supplier.Entity == "" {
			supplier.Entity = "supplier"
		}

		var err error
		id, err = h.repo.Insert(tx, *supplier)
		if err != nil {
			return errors.ErrInternalServer
		}

		for _, account := range input.BankAccounts {
			newAccount := data.BankAccount{
				Title:      account,
				SupplierID: id,
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			}

			if _, err = h.bankAccountRepo.Insert(tx, newAccount); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	supplier, err = supplier.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	supplier.BankAccounts = input.BankAccounts

	res := dto.ToSupplierResponseDTO(*supplier)

	return &res, nil
}

func (h *SupplierServiceImpl) UpdateSupplier(id int, input dto.SupplierDTO) (*dto.SupplierResponseDTO, error) {
	updatedData := input.ToSupplier()

	err := data.Upper.Tx(func(tx up.Session) error {
		// Get existing bank accounts for the supplier
		existingBankAccounts, err := h.bankAccountRepo.GetSupplierBankAccounts(id)
		if err != nil {
			return err
		}

		existingBankAccountSet := make(map[string]struct{})
		for _, account := range existingBankAccounts {
			existingBankAccountSet[account] = struct{}{}
		}

		receivedAccountSet := make(map[string]struct{})
		for _, account := range input.BankAccounts {
			receivedAccountSet[account] = struct{}{}
		}

		// Insert new bank accounts which exist in the received list but not in the existing list
		for account := range receivedAccountSet {
			if _, ok := existingBankAccountSet[account]; !ok {
				newAccount := data.BankAccount{
					Title:      account,
					SupplierID: id,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}

				if _, err = h.bankAccountRepo.Insert(tx, newAccount); err != nil {
					return err
				}
			}
		}

		// Delete bank accounts which exist in the existing list but not in the received list
		for account := range existingBankAccountSet {
			if _, ok := receivedAccountSet[account]; !ok {
				if err = h.bankAccountRepo.Delete(tx, account); err != nil {
					return err
				}
			}
		}

		updatedData.ID = id

		if updatedData.Entity == "" {
			updatedData.Entity = "supplier"
		}

		err = h.repo.Update(*updatedData)
		if err != nil {
			return errors.ErrInternalServer
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	updatedData, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	updatedData.BankAccounts = input.BankAccounts

	response := dto.ToSupplierResponseDTO(*updatedData)

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

	// Get bank account for the supplier
	accounts, err := h.bankAccountRepo.GetSupplierBankAccounts(id)
	if err != nil {
		return nil, err
	}

	data.BankAccounts = accounts

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

	// Get bank accounts for each supplier
	for _, supplier := range data {
		accounts, err := h.bankAccountRepo.GetSupplierBankAccounts(supplier.ID)
		if err != nil {
			return nil, nil, err
		}

		supplier.BankAccounts = accounts
	}

	response := dto.ToSupplierListResponseDTO(data)

	return response, total, nil
}
