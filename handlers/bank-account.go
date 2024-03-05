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

// BankAccountHandler is a concrete type that implements BankAccountHandler
type bankaccountHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.BankAccountService
}

// NewBankAccountHandler initializes a new BankAccountHandler with its dependencies
func NewBankAccountHandler(app *celeritas.Celeritas, bankAccountService services.BankAccountService) BankAccountHandler {
	return &bankaccountHandlerImpl{
		App:     app,
		service: bankAccountService,
	}
}

func (h *bankaccountHandlerImpl) CreateBankAccount(w http.ResponseWriter, r *http.Request) {
	var input dto.BankAccountDTO
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

	res, err := h.service.CreateBankAccount(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "BankAccount created successfuly", res)
}

func (h *bankaccountHandlerImpl) UpdateBankAccount(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var input dto.BankAccountDTO
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

	res, err := h.service.UpdateBankAccount(id, input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "BankAccount updated successfuly", res)
}

func (h *bankaccountHandlerImpl) DeleteBankAccount(w http.ResponseWriter, r *http.Request) {
	title := chi.URLParam(r, "title")

	err := h.service.DeleteBankAccount(title)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "BankAccount deleted successfuly")
}

func (h *bankaccountHandlerImpl) GetBankAccountById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := h.service.GetBankAccount(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", res)
}

func (h *bankaccountHandlerImpl) GetBankAccountList(w http.ResponseWriter, r *http.Request) {
	var filter dto.BankAccountFilterDTO

	_ = h.App.ReadJSON(w, r, &filter)

	validator := h.App.Validator().ValidateStruct(&filter)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	res, total, err := h.service.GetBankAccountList(filter)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", res, int(*total))
}
