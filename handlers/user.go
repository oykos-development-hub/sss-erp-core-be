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

	userIDString := r.Header.Get("UserID")

	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrUnauthorized), errors.ErrBadRequest, validator.Errors)
		return
	}

	ctx := context.Background()
	ctx = contextutil.SetUserIDInContext(ctx, userID)

	user, err := h.service.CreateUser(ctx, userInput)
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

	userIDString := r.Header.Get("UserID")

	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrUnauthorized), errors.ErrBadRequest, validator.Errors)
		return
	}

	ctx := context.Background()
	ctx = contextutil.SetUserIDInContext(ctx, userID)

	response, err := h.service.UpdateUser(ctx, id, userInput)
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

	userIDString := r.Header.Get("id")

	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(errors.ErrUnauthorized), errors.ErrBadRequest)
		return
	}

	ctx := context.Background()
	ctx = contextutil.SetUserIDInContext(ctx, userID)

	err = h.service.DeleteUser(ctx, id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "User deleted successfuly")
}
