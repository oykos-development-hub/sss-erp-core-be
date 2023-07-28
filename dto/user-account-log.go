package dto

import (
	"encoding/json"
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
)

type UserAccountLogDTO struct {
	TargetUserAccountID int             `json:"target_user_account_id"`
	SourceUserAccountID int             `json:"source_user_account_id"`
	ChangeType          int             `json:"change_type"`
	PreviousValue       json.RawMessage `json:"previous_value"`
	NewValue            json.RawMessage `json:"new_value"`
}

type UserAccountLogResponseDTO struct {
	ID                  int             `json:"id"`
	CreatedAt           time.Time       `json:"created_at"`
	TargetUserAccountID int             `json:"target_user_account_id"`
	SourceUserAccountID int             `json:"source_user_account_id"`
	ChangeType          int             `json:"change_type"`
	PreviousValue       json.RawMessage `json:"previous_value"`
	NewValue            json.RawMessage `json:"new_value"`
}

func (dto UserAccountLogDTO) ToUserAccountLog() *data.UserAccountLog {
	return &data.UserAccountLog{
		TargetUserAccountID: dto.TargetUserAccountID,
		SourceUserAccountID: dto.SourceUserAccountID,
		ChangeType:          dto.ChangeType,
		PreviousValue:       dto.PreviousValue,
		NewValue:            dto.NewValue,
		CreatedAt:           time.Now(),
	}
}

func ToUserAccountLogResponseDTO(data data.UserAccountLog) UserAccountLogResponseDTO {
	return UserAccountLogResponseDTO{
		ID:                  data.ID,
		TargetUserAccountID: data.TargetUserAccountID,
		SourceUserAccountID: data.SourceUserAccountID,
		ChangeType:          data.ChangeType,
		PreviousValue:       data.PreviousValue,
		NewValue:            data.NewValue,
		CreatedAt:           data.CreatedAt,
	}
}

func ToUserAccountLogListResponseDTO(user_account_logs []*data.UserAccountLog) []UserAccountLogResponseDTO {
	dtoList := make([]UserAccountLogResponseDTO, len(user_account_logs))
	for i, x := range user_account_logs {
		dtoList[i] = ToUserAccountLogResponseDTO(*x)
	}
	return dtoList
}
