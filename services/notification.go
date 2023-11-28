package services

import (
	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
)

type NotificationServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.Notification
}

func NewNotificationServiceImpl(app *celeritas.Celeritas, repo data.Notification) NotificationService {
	return &NotificationServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *NotificationServiceImpl) CreateNotification(input dto.NotificationDTO) (*dto.NotificationResponseDTO, error) {
	data := input.ToNotification()

	id, err := h.repo.Insert(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = data.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	res := dto.ToNotificationResponseDTO(*data)

	return &res, nil
}

func (h *NotificationServiceImpl) UpdateNotification(id int, input dto.NotificationDTO) (*dto.NotificationResponseDTO, error) {
	data := input.ToNotification()
	data.ID = id

	err := h.repo.Update(*data)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	data, err = h.repo.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToNotificationResponseDTO(*data)

	return &response, nil
}

func (h *NotificationServiceImpl) DeleteNotification(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}

func (h *NotificationServiceImpl) GetNotification(id int) (*dto.NotificationResponseDTO, error) {
	data, err := h.repo.Get(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToNotificationResponseDTO(*data)

	return &response, nil
}

func (h *NotificationServiceImpl) GetNotificationList() ([]dto.NotificationResponseDTO, error) {
	data, err := h.repo.GetAll(nil)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}
	response := dto.ToNotificationListResponseDTO(data)

	return response, nil
}
