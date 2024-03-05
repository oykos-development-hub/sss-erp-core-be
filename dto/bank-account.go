package dto

import (
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
)

type BankAccountDTO struct {
	Title string `json:"title" validate:"required,min=2"`
	SupplierID int `json:"supplier_id" validate:"required"`
}

type BankAccountResponseDTO struct {
	ID 				int 			`json:"id"`
	Title 		string 		`json:"title"`
	SupplierID int `json:"supplier_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BankAccountFilterDTO struct {
	Page *int `json:"page"`
	Size *int `json:"size"`
	SupplierID *int `json:"supplier_id"`
	SortByTitle    *string `json:"sort_by_title"`
}

func (dto BankAccountDTO) ToBankAccount() *data.BankAccount {
	return &data.BankAccount{
		Title:     dto.Title,
		SupplierID: dto.SupplierID,
	}
}

func ToBankAccountResponseDTO(data data.BankAccount) BankAccountResponseDTO {
	return BankAccountResponseDTO{
		ID: data.ID,
		Title: data.Title,
		SupplierID: data.SupplierID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToBankAccountListResponseDTO(bank_accounts []*data.BankAccount) []BankAccountResponseDTO {
	dtoList := make([]BankAccountResponseDTO, len(bank_accounts))
	for i, x := range bank_accounts {
		dtoList[i] = ToBankAccountResponseDTO(*x)
	}
	return dtoList
}
