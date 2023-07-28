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

type userHandlerImpl struct {
	App     *celeritas.Celeritas
	service services.UserService
}

func NewUserHandler(app *celeritas.Celeritas, userService services.UserService) UserHandler {
	return &userHandlerImpl{
		App:     app,
		service: userService,
	}
}

func (h *userHandlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInput dto.UserRegistrationDTO
	_ = h.App.ReadJSON(w, r, &userInput)

	validator := h.App.Validator().ValidateStruct(&userInput)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	user, err := h.service.CreateUser(userInput)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "User created successfuly", user)
}

func (h *userHandlerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var userInput dto.UserUpdateDTO
	_ = h.App.ReadJSON(w, r, &userInput)

	validator := h.App.Validator().ValidateStruct(&userInput)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	response, err := h.service.UpdateUser(id, userInput)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "User updated successfuly", response)
}

func (h *userHandlerImpl) GetLoggedInUser(w http.ResponseWriter, r *http.Request) {
	userIdString := r.Header.Get("id")
	userId, _ := strconv.Atoi(userIdString)

	user, err := h.service.GetUser(userId)

	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "Valid token", user)
}

func (h *userHandlerImpl) GetUserById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	user, err := h.service.GetUser(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponse(w, http.StatusOK, "", user)
}

func (h *userHandlerImpl) GetUserList(w http.ResponseWriter, r *http.Request) {
	var input dto.GetUserListDTO
	_ = h.App.ReadJSON(w, r, &input)

	validator := h.App.Validator().ValidateStruct(&input)
	if !validator.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, validator.Errors)
		return
	}

	users, total, err := h.service.GetUserList(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteDataResponseWithTotal(w, http.StatusOK, "", users, int(*total))
}

func (h *userHandlerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.DeleteUser(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "User deleted successfuly")
}
