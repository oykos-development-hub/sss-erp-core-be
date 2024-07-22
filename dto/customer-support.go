package dto

import (
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
)

type CustomerSupportDTO struct {
	UserDocumentationFileID int `json:"user_documentation_file_id"`
}

type CustomerSupportResponseDTO struct {
	ID                      int       `json:"id"`
	UserDocumentationFileID int       `json:"user_documentation_file_id"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

type CustomerSupportFilterDTO struct {
	Page        *int    `json:"page"`
	Size        *int    `json:"size"`
	SortByTitle *string `json:"sort_by_title"`
}

func (dto CustomerSupportDTO) ToCustomerSupport() *data.CustomerSupport {
	return &data.CustomerSupport{
		UserDocumentationFileID: dto.UserDocumentationFileID,
	}
}

func ToCustomerSupportResponseDTO(data data.CustomerSupport) CustomerSupportResponseDTO {
	return CustomerSupportResponseDTO{
		ID:                      data.ID,
		UserDocumentationFileID: data.UserDocumentationFileID,
		CreatedAt:               data.CreatedAt,
		UpdatedAt:               data.UpdatedAt,
	}
}

func ToCustomerSupportListResponseDTO(customer_supports []*data.CustomerSupport) []CustomerSupportResponseDTO {
	dtoList := make([]CustomerSupportResponseDTO, len(customer_supports))
	for i, x := range customer_supports {
		dtoList[i] = ToCustomerSupportResponseDTO(*x)
	}
	return dtoList
}
