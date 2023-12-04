package dto

import (
	"encoding/json"
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
)

type GetNotificationListInput struct {
	ToUserID *int `json:"to_user_id"`
	Page     *int `json:"page"`
	Size     *int `json:"size"`
}

type NotificationDTO struct {
	FromContent string          `json:"from_content"`
	FromUserID  int             `json:"from_user_id"`
	ToUserID    int             `json:"to_user_id"`
	Module      string          `json:"module"`
	Content     string          `json:"content"`
	IsRead      bool            `json:"is_read"`
	Path        string          `json:"path"`
	Data        json.RawMessage `json:"data"`
}

type NotificationResponseDTO struct {
	ID          int             `json:"id"`
	FromContent string          `json:"from_content"`
	FromUserID  int             `json:"from_user_id"`
	ToUserID    int             `json:"to_user_id"`
	Module      string          `json:"module"`
	Content     string          `json:"content"`
	IsRead      bool            `json:"is_read"`
	Path        string          `json:"path"`
	Data        json.RawMessage `json:"data"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

func (dto NotificationDTO) ToNotification() *data.Notification {
	return &data.Notification{
		FromContent: dto.FromContent,
		FromUserID:  dto.FromUserID,
		ToUserID:    dto.ToUserID,
		Module:      dto.Module,
		Content:     dto.Content,
		IsRead:      dto.IsRead,
		Path:        dto.Path,
		Data:        dto.Data,
	}
}

func ToNotificationResponseDTO(data data.Notification) NotificationResponseDTO {
	return NotificationResponseDTO{
		ID:          data.ID,
		FromContent: data.FromContent,
		FromUserID:  data.FromUserID,
		ToUserID:    data.ToUserID,
		Module:      data.Module,
		Content:     data.Content,
		IsRead:      data.IsRead,
		Path:        data.Path,
		Data:        data.Data,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func ToNotificationListResponseDTO(notifications []*data.Notification) []NotificationResponseDTO {
	dtoList := make([]NotificationResponseDTO, len(notifications))
	for i, x := range notifications {
		dtoList[i] = ToNotificationResponseDTO(*x)
	}
	return dtoList
}
