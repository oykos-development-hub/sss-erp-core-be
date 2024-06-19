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

func (h *rolespermissionHandlerImpl) SyncPermissions(w http.ResponseWriter, r *http.Request) {
	roleID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input []dto.RolesPermissionDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		h.App.ErrorLog.Print(validator.Errors)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, err := h.service.SyncPermissions(roleID, input)
	if err != nil {
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "Permission created successfuly", res)
}
