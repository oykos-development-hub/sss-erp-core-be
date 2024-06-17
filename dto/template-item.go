package dto

import (
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
)

type TemplateItemDTO struct {
	TemplateID         int `json:"template_id"`
	FileID             int `json:"file_id"`
	OrganizationUnitID int `json:"organization_unit_id"`
}

type TemplateItemResponseDTO struct {
	ID                 int       `json:"id"`
	TemplateID         int       `json:"template_id"`
	FileID             int       `json:"file_id"`
	OrganizationUnitID int       `json:"organization_unit_id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type TemplateItemFilterDTO struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	SortByTitle        *string `json:"sort_by_title"`
	TemplateID         *int    `json:"template_id"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
}

func (dto TemplateItemDTO) ToTemplateItem() *data.TemplateItem {
	return &data.TemplateItem{
		TemplateID:         dto.TemplateID,
		FileID:             dto.FileID,
		OrganizationUnitID: dto.OrganizationUnitID,
	}
}

func ToTemplateItemResponseDTO(data data.TemplateItem) TemplateItemResponseDTO {
	return TemplateItemResponseDTO{
		ID:                 data.ID,
		TemplateID:         data.TemplateID,
		FileID:             data.FileID,
		OrganizationUnitID: data.OrganizationUnitID,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
	}
}

func ToTemplateItemListResponseDTO(template_items []*data.TemplateItem) []TemplateItemResponseDTO {
	dtoList := make([]TemplateItemResponseDTO, len(template_items))
	for i, x := range template_items {
		dtoList[i] = ToTemplateItemResponseDTO(*x)
	}
	return dtoList
}
