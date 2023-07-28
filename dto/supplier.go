package dto

import (
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
)

type GetSupplierListInput struct {
	Search *string `json:"search"`
	Page   *int    `json:"page"`
	Size   *int    `json:"size"`
}

type SupplierDTO struct {
	Title        string `json:"title" validate:"required"`
	Abbreviation string `json:"abbreviation"`
	OfficialID   string `json:"official_id"`
	Address      string `json:"address"`
	Description  string `json:"description"`
	FolderID     int    `json:"folder_id"`
}

type SupplierResponseDTO struct {
	ID           int       `json:"id"`
	Title        string    `json:"title" validate:"required"`
	Abbreviation string    `json:"abbreviation"`
	OfficialID   string    `json:"official_id"`
	Address      string    `json:"address"`
	Description  string    `json:"description"`
	FolderID     int       `json:"folder_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (dto SupplierDTO) ToSupplier() *data.Supplier {
	return &data.Supplier{
		Title:        dto.Title,
		Abbreviation: dto.Abbreviation,
		OfficialID:   dto.OfficialID,
		Address:      dto.Address,
		Description:  dto.Description,
		FolderID:     dto.FolderID,
	}
}

func ToSupplierResponseDTO(data data.Supplier) SupplierResponseDTO {
	return SupplierResponseDTO{
		ID:           data.ID,
		Title:        data.Title,
		Abbreviation: data.Abbreviation,
		OfficialID:   data.OfficialID,
		Address:      data.Address,
		Description:  data.Description,
		FolderID:     data.FolderID,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func ToSupplierListResponseDTO(suppliers []*data.Supplier) []SupplierResponseDTO {
	dtoList := make([]SupplierResponseDTO, len(suppliers))
	for i, x := range suppliers {
		dtoList[i] = ToSupplierResponseDTO(*x)
	}
	return dtoList
}
