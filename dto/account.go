package dto

import (
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
)

type AccountDTO struct {
	Title        string `json:"title" validate:"required,min=2"`
	ParentID     *int   `json:"parent_id"`
	SerialNumber string `json:"serial_number"`
}

type AccountResponseDTO struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	ParentID     *int      `json:"parent_id"`
	SerialNumber string    `json:"serial_number"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (dto AccountDTO) ToAccount() *data.Account {
	return &data.Account{
		Title:        dto.Title,
		ParentID:     dto.ParentID,
		SerialNumber: dto.SerialNumber,
	}
}

func ToAccountResponseDTO(data data.Account) AccountResponseDTO {
	return AccountResponseDTO{
		ID:           data.ID,
		Title:        data.Title,
		ParentID:     data.ParentID,
		SerialNumber: data.SerialNumber,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func ToAccountListResponseDTO(accounts []*data.Account) []AccountResponseDTO {
	dtoList := make([]AccountResponseDTO, len(accounts))
	for i, x := range accounts {
		dtoList[i] = ToAccountResponseDTO(*x)
	}
	return dtoList
}
