package dto

import (
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
)

type GetNotificationListInput struct {
	ToUserID *string `json:"to_user_id"`
	Page     *int    `json:"page"`
	Size     *int    `json:"size"`
}

type NotificationDTO struct {
	From       string `json:"from"`
	FromUserID int    `json:"from_user_id"`
	ToUserID   int    `json:"to_user_id"`
	Module     string `json:"module"`
	Content    string `json:"content"`
	IsRead     bool   `json:"is_read"`
}

type NotificationResponseDTO struct {
	ID         int       `json:"id"`
	From       string    `json:"from"`
	FromUserID int       `json:"from_user_id"`
	ToUserID   int       `json:"to_user_id"`
	Module     string    `json:"module"`
	Content    string    `json:"content"`
	IsRead     bool      `json:"is_read"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (dto NotificationDTO) ToNotification() *data.Notification {
	return &data.Notification{
		From:       dto.From,
		FromUserID: dto.FromUserID,
		ToUserID:   dto.ToUserID,
		Module:     dto.Module,
		Content:    dto.Content,
		IsRead:     dto.IsRead,
	}
}

func ToNotificationResponseDTO(data data.Notification) NotificationResponseDTO {
	return NotificationResponseDTO{
		ID:         data.ID,
		From:       data.From,
		FromUserID: data.FromUserID,
		ToUserID:   data.ToUserID,
		Module:     data.Module,
		Content:    data.Content,
		IsRead:     data.IsRead,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}
}

func ToNotificationListResponseDTO(notifications []*data.Notification) []NotificationResponseDTO {
	dtoList := make([]NotificationResponseDTO, len(notifications))
	for i, x := range notifications {
		dtoList[i] = ToNotificationResponseDTO(*x)
	}
	return dtoList
}
