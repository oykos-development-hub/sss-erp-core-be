package dto

import (
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
)

type BffErrorLogDTO struct {
	Error  string                `json:"error"`
	Code   int                   `json:"code"`
	Entity data.HandlersEntities `json:"entity"`
}

type BffErrorLogResponseDTO struct {
	ID        int                   `json:"id"`
	Error     string                `json:"error"`
	Code      int                   `json:"code"`
	Entity    data.HandlersEntities `json:"entity"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}

type BffErrorLogFilterDTO struct {
	Page        *int       `json:"page"`
	Size        *int       `json:"size"`
	DateOfStart *time.Time `json:"date_of_start"`
	DateOfEnd   *time.Time `json:"date_of_end"`
	Entity      *string    `json:"entity"`
}

func (dto BffErrorLogDTO) ToBffErrorLog() *data.BffErrorLog {
	return &data.BffErrorLog{
		Error:  dto.Error,
		Code:   dto.Code,
		Entity: dto.Entity,
	}
}

func ToBffErrorLogResponseDTO(data data.BffErrorLog) BffErrorLogResponseDTO {
	return BffErrorLogResponseDTO{
		ID:        data.ID,
		Error:     data.Error,
		Code:      data.Code,
		Entity:    data.Entity,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToBffErrorLogListResponseDTO(error_logs []*data.BffErrorLog) []BffErrorLogResponseDTO {
	dtoList := make([]BffErrorLogResponseDTO, len(error_logs))
	for i, x := range error_logs {
		dtoList[i] = ToBffErrorLogResponseDTO(*x)
	}
	return dtoList
}
