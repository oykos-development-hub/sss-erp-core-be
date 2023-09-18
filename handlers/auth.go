package handlers

import (
	"net/http"
	"strconv"
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"
	"gitlab.sudovi.me/erp/core-ms-api/services"

	"github.com/go-chi/chi/v5"
	"github.com/oykos-development-hub/celeritas"
)

type authHandlerImpl struct {
	App        *celeritas.Celeritas
	service    services.AuthService
	logService services.UserAccountLogService
}

func NewAuthHandler(app *celeritas.Celeritas, authService services.AuthService, logService services.UserAccountLogService) AuthHandler {
	return &authHandlerImpl{
		App:        app,
		service:    authService,
		logService: logService,
	}
}

func (h *authHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var loginInput dto.LoginInput

	_ = h.App.ReadJSON(w, r, &loginInput)

	loginRes, err := h.service.Login(loginInput)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	cookieExpireDuration, _ := time.ParseDuration(h.App.JwtToken.JwtRefreshTokenTimeExp.String())
	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    loginRes.Token.RefreshToken.Value,
		MaxAge:   0,
		HttpOnly: true,
		Expires:  time.Now().Add(cookieExpireDuration),
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}

	http.SetCookie(w, cookie)
	_ = h.App.WriteDataResponse(w, http.StatusOK, "", loginRes)
}

func (h *authHandlerImpl) ValidatePin(w http.ResponseWriter, r *http.Request) {
	userIdString := r.Header.Get("id")
	userId, _ := strconv.Atoi(userIdString)

	var input dto.ValidatePinInput
	_ = h.App.ReadJSON(w, r, &input)

	v := h.App.Validator().ValidateStruct(&input)
	if !v.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, v.Errors)
		return
	}

	err := h.service.ValidatePin(userId, input)

	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}
	/*
		userLog := dto.UserAccountLogDTO{
			TargetUserAccountID: userId,
			SourceUserAccountID: userId,
			ChangeType:          1,
			PreviousValue:       json.RawMessage(`{"test": "1"}`),
			NewValue:            json.RawMessage(`{"test": "2"}`),
		}

		_, err = h.logService.CreateUserAccountLog(userLog)

		if err != nil {
			_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
			return
		}
	*/
	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "Valid pin")
}

func (h *authHandlerImpl) RefreshToken(w http.ResponseWriter, r *http.Request) {
	userIdString := r.Header.Get("id")
	iat := r.Header.Get("iat")
	refreshToken := r.Header.Get("refresh_token")
	userId, _ := strconv.Atoi(userIdString)

	token, err := h.service.RefreshToken(userId, refreshToken, iat)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	cookieExpireDuration, _ := time.ParseDuration(h.App.JwtToken.JwtRefreshTokenTimeExp.String())
	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    token.RefreshToken.Value,
		MaxAge:   0,
		HttpOnly: true,
		Expires:  time.Now().Add(cookieExpireDuration),
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}

	http.SetCookie(w, cookie)
	_ = h.App.WriteDataResponse(w, http.StatusOK, "", token)
}

func (h *authHandlerImpl) Logout(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	err := h.service.Logout(id)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "Successfully revoked tokens", nil)
}

func (h *authHandlerImpl) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var forgotPasswordInput dto.ForgotPassword
	_ = h.App.ReadJSON(w, r, &forgotPasswordInput)

	err := h.service.ForgotPassword(forgotPasswordInput)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "Email sent")
}

func (h *authHandlerImpl) ResetPasswordVerify(w http.ResponseWriter, r *http.Request) {
	var input dto.ResetPasswordVerify
	_ = h.App.ReadJSON(w, r, &input)

	v := h.App.Validator().ValidateStruct(&input)
	if !v.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, v.Errors)
		return
	}

	ok, err := h.service.ResetPasswordVerify(input.Email, input.Token)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	if ok {
		_ = h.App.WriteSuccessResponse(w, http.StatusOK, "Reset password link verified")
	} else {
		_ = h.App.WriteErrorResponse(w, http.StatusOK, errors.ErrBadRequest)
	}
}

func (h *authHandlerImpl) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var input dto.ResetPassword
	_ = h.App.ReadJSON(w, r, &input)

	v := h.App.Validator().ValidateStruct(&input)
	if !v.Valid() {
		_ = h.App.WriteErrorResponseWithData(w, errors.MapErrorToStatusCode(errors.ErrBadRequest), errors.ErrBadRequest, v.Errors)
		return
	}

	ok, err := h.service.ResetPasswordVerify(input.Email, input.Token)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	if !ok {
		_ = h.App.WriteErrorResponse(w, http.StatusOK, errors.ErrBadRequest)
	}

	err = h.service.ResetPassword(input)
	if err != nil {
		_ = h.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(err), err)
		return
	}

	_ = h.App.WriteSuccessResponse(w, http.StatusOK, "Password reset successful")
}
