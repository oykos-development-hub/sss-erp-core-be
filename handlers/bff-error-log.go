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

// BffErrorLogHandler is a concrete type that implements BffErrorLogHandler
type bfferrorlogHandlerImpl struct {
	App             *celeritas.Celeritas
	service         services.BffErrorLogService
	errorLogService services.ErrorLogService
}

// NewBffErrorLogHandler initializes a new BffErrorLogHandler with its dependencies
func NewBffErrorLogHandler(app *celeritas.Celeritas, bfferrorlogService services.BffErrorLogService, errorLogService services.ErrorLogService) BffErrorLogHandler {
	return &bfferrorlogHandlerImpl{
		App:             app,
		service:         bfferrorlogService,
		errorLogService: errorLogService,
	}
}

func (h *bfferrorlogHandlerImpl) CreateBffErrorLog(w http.ResponseWriter, r *http.Request) {
	var input dto.BffErrorLogDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateBffErrorLog(input)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "BffErrorLog created successfuly", res)
}

func (h *bfferrorlogHandlerImpl) UpdateBffErrorLog(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.BffErrorLogDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateBffErrorLog(id, input)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "BffErrorLog updated successfuly", res)
}

func (h *bfferrorlogHandlerImpl) DeleteBffErrorLog(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteBffErrorLog(id)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "BffErrorLog deleted successfuly")
}

func (h *bfferrorlogHandlerImpl) GetBffErrorLogById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetBffErrorLog(id)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *bfferrorlogHandlerImpl) GetBffErrorLogList(w http.ResponseWriter, r *http.Request) {
	var filter dto.BffErrorLogFilterDTO

	_ = h.App.ReadJSON(w, r, &filter)

	validator := h.App.Validator().ValidateStruct(&filter)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, total, err := h.service.GetBffErrorLogList(filter)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, int(*total))
}
