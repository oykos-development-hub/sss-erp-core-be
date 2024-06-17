package dto

import (
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
)

type TemplateDTO struct {
	Title string `json:"title" validate:"required,min=2"`
}

type TemplateResponseDTO struct {
	ID 				int 			`json:"id"`
	Title 		string 		`json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TemplateFilterDTO struct {
	Page *int `json:"page"`
	Size *int `json:"size"`
	SortByTitle    *string `json:"sort_by_title"`
}

func (dto TemplateDTO) ToTemplate() *data.Template {
	return &data.Template{
		Title:     dto.Title,
	}
}

func ToTemplateResponseDTO(data data.Template) TemplateResponseDTO {
	return TemplateResponseDTO{
		ID: data.ID,
		Title: data.Title,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToTemplateListResponseDTO(templates []*data.Template) []TemplateResponseDTO {
	dtoList := make([]TemplateResponseDTO, len(templates))
	for i, x := range templates {
		dtoList[i] = ToTemplateResponseDTO(*x)
	}
	return dtoList
}

