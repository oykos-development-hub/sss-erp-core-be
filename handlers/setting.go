package handlers

import (
	"net/http"
	"strconv"

	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"
	"gitlab.sudovi.me/erp/core-ms-api/services"

	"github.com/go-chi/chi/v5"
	"github.com/oykos-development-hub/celeritas"
)

// SettingHandler is a concrete type that implements SettingHandler
type settingHandlerImpl struct {
	App             *celeritas.Celeritas
	service         services.SettingService
	errorLogService services.ErrorLogService
}

// NewSettingHandler initializes a new SettingHandler with its dependencies
func NewSettingHandler(app *celeritas.Celeritas, settingService services.SettingService, errorLogService services.ErrorLogService) SettingHandler {
	return &settingHandlerImpl{
		App:             app,
		service:         settingService,
		errorLogService: errorLogService,
	}
}

func (h *settingHandlerImpl) CreateSetting(w http.ResponseWriter, r *http.Request) {
	var input dto.SettingDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateSetting(input)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "Setting created successfuly", res)
}

func (h *settingHandlerImpl) UpdateSetting(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.SettingDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateSetting(id, input)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "Setting updated successfuly", res)
}

func (h *settingHandlerImpl) DeleteSetting(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteSetting(id)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "Setting deleted successfuly")
}

func (h *settingHandlerImpl) GetSettingById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetSetting(id)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *settingHandlerImpl) GetSettingList(w http.ResponseWriter, r *http.Request) {
	var input dto.GetSettingsDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, total, err := h.service.GetSettingList(input)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, int(*total))
}
