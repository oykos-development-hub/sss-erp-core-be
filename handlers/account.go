package handlers

import (
	"context"
	"net/http"
	"strconv"

	"gitlab.sudovi.me/erp/core-ms-api/contextutil"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"
	"gitlab.sudovi.me/erp/core-ms-api/services"

	"github.com/go-chi/chi/v5"
	"github.com/oykos-development-hub/celeritas"
)

// AccountHandler is a concrete type that implements AccountHandler
type accountHandlerImpl struct {
	App             *celeritas.Celeritas
	service         services.AccountService
	errorLogService services.ErrorLogService
}

// NewAccountHandler initializes a new AccountHandler with its dependencies
func NewAccountHandler(app *celeritas.Celeritas, accountService services.AccountService, errorLogService services.ErrorLogService) AccountHandler {
	return &accountHandlerImpl{
		App:             app,
		service:         accountService,
		errorLogService: errorLogService,
	}
}

func (h *accountHandlerImpl) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var input []dto.AccountDTO
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
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

	userIDString := r.Header.Get("UserID")

	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrUnauthorized), errors.ErrBadRequest, validator.Errors)
		return
	}

	ctx := context.Background()
	ctx = contextutil.SetUserIDInContext(ctx, userID)

	res, err := h.service.CreateAccountList(ctx, input)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "Account created successfuly", res)
}

func (h *accountHandlerImpl) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	userIDString := r.Header.Get("UserID")

	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(errors.ErrUnauthorized), errors.ErrUnauthorized)
		return
	}

	ctx := context.Background()
	ctx = contextutil.SetUserIDInContext(ctx, userID)

	err = h.service.DeleteAccount(ctx, id)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "Account deleted successfuly")
}

func (h *accountHandlerImpl) GetAccountById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetAccount(id)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *accountHandlerImpl) GetAccountList(w http.ResponseWriter, r *http.Request) {
	var input dto.GetAccountsFilter
	err := h.App.ReadJSON(w, r, &input)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	res, total, err := h.service.GetAccountList(input)
	if err != nil {
		h.errorLogService.CreateErrorLog(err)
		h.App.ErrorLog.Print(err)
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, total)
}
