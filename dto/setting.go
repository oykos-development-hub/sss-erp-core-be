package dto

import (
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
)

type SettingDTO struct {
	Title        string  `json:"title" validate:"required,min=2"`
	Abbreviation string  `json:"abbreviation" validate:"required,min=2,max=4"`
	Value        *string `json:"value" validate:"omitempty"`
	Entity       string  `json:"entity" validate:"required"`
	Description  *string `json:"description" validate:"omitempty,min=2"`
	Color        *string `json:"color" validate:"omitempty,min=2"`
	Icon         *string `json:"icon" validate:"omitempty,min=2"`
}

type SettingResponseDTO struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Abbreviation string    `json:"abbreviation"`
	Value        *string   `json:"value"`
	Entity       string    `json:"entity"`
	Description  *string   `json:"description"`
	Color        *string   `json:"color"`
	Icon         *string   `json:"icon"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (dto SettingDTO) ToSetting() *data.Setting {
	return &data.Setting{
		Title:        dto.Title,
		Abbreviation: dto.Abbreviation,
		Entity:       dto.Entity,
		Description:  dto.Description,
		Value:        dto.Value,
		Color:        dto.Color,
		Icon:         dto.Icon,
	}
}

func ToSettingResponseDTO(data data.Setting) SettingResponseDTO {
	return SettingResponseDTO{
		ID:           data.ID,
		Entity:       data.Entity,
		Title:        data.Title,
		Abbreviation: data.Abbreviation,
		Description:  data.Description,
		Value:        data.Value,
		Color:        data.Color,
		Icon:         data.Icon,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func ToSettingListResponseDTO(settings []*data.Setting) []SettingResponseDTO {
	dtoList := make([]SettingResponseDTO, len(settings))
	for i, x := range settings {
		dtoList[i] = ToSettingResponseDTO(*x)
	}
	return dtoList
}

type GetSettingsDTO struct {
	Entity string  `json:"entity" validate:"required"`
	Page   *int    `json:"page" validate:"omitempty"`
	Size   *int    `json:"size" validate:"omitempty"`
	Search *string `json:"search" validate:"omitempty"`
}
