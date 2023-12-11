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

// RolesPermissionHandler is a concrete type that implements RolesPermissionHandler
type rolespermissionHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.RolesPermissionService
}

// NewRolesPermissionHandler initializes a new RolesPermissionHandler with its dependencies
func NewRolesPermissionHandler(app *celeritas.Celeritas, rolespermissionService services.RolesPermissionService) RolesPermissionHandler {
	return &rolespermissionHandlerImpl{
		App:     app,
		service: rolespermissionService,
	}
}

func (h *rolespermissionHandlerImpl) CreateRolesPermission(w http.ResponseWriter, r *http.Request) {
	var input dto.RolesPermissionDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.CreateRolesPermission(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "RolesPermission created successfuly", res)
}

func (h *rolespermissionHandlerImpl) UpdateRolesPermission(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.RolesPermissionDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.UpdateRolesPermission(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "RolesPermission updated successfuly", res)
}

func (h *rolespermissionHandlerImpl) DeleteRolesPermission(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteRolesPermission(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "RolesPermission deleted successfuly")
}

func (h *rolespermissionHandlerImpl) GetRolesPermissionById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetRolesPermission(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *rolespermissionHandlerImpl) GetRolesPermissionList(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.GetRolesPermissionList()
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}
