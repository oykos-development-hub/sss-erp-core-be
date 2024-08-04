package dto

import (
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
)

type ListOfParameterDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ListOfParameterResponseDTO struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ListOfParameterFilterDTO struct {
	Page        *int    `json:"page"`
	Size        *int    `json:"size"`
	SortByTitle *string `json:"sort_by_title"`
}

func (dto ListOfParameterDTO) ToListOfParameter() *data.ListOfParameter {
	return &data.ListOfParameter{
		Title:       dto.Title,
		Description: dto.Description,
	}
}

func ToListOfParameterResponseDTO(data data.ListOfParameter) ListOfParameterResponseDTO {
	return ListOfParameterResponseDTO{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func ToListOfParameterListResponseDTO(list_of_parameters []*data.ListOfParameter) []ListOfParameterResponseDTO {
	dtoList := make([]ListOfParameterResponseDTO, len(list_of_parameters))
	for i, x := range list_of_parameters {
		dtoList[i] = ToListOfParameterResponseDTO(*x)
	}
	return dtoList
}
