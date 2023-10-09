package handlers

import (
	"net/http"
	"strconv"

	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"
	"gitlab.sudovi.me/erp/core-ms-api/services"

	"github.com/oykos-development-hub/celeritas"
	"github.com/go-chi/chi/v5"
)

// AccountHandler is a concrete type that implements AccountHandler
type accountHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.AccountService
}

// NewAccountHandler initializes a new AccountHandler with its dependencies
func NewAccountHandler(app *celeritas.Celeritas, accountService services.AccountService) AccountHandler {
	return &accountHandlerImpl{
		App:     app,
		service: accountService,
	}
}

func (h *accountHandlerImpl) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var input dto.AccountDTO
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

	res, err := h.service.CreateAccount(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "Account created successfuly", res)
}

func (h *accountHandlerImpl) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.AccountDTO
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

	res, err := h.service.UpdateAccount(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "Account updated successfuly", res)
}

func (h *accountHandlerImpl) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteAccount(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "Account deleted successfuly")
}

func (h *accountHandlerImpl) GetAccountById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetAccount(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *accountHandlerImpl) GetAccountList(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.GetAccountList()
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}
