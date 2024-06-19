package services

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"net/url"
	"strconv"
	"time"

	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	newErrors "gitlab.sudovi.me/erp/core-ms-api/pkg/errors"

	"github.com/golang-jwt/jwt"
	"github.com/oykos-development-hub/celeritas"
	jwtdto "github.com/oykos-development-hub/celeritas/jwt/dto"
	"github.com/oykos-development-hub/celeritas/mailer"
	"github.com/oykos-development-hub/celeritas/urlsigner"
)

type authServiceImpl struct {
	App      *celeritas.Celeritas
	userRepo data.User
	BaseService
	logRepo data.Log
}

func NewAuthServiceImpl(app *celeritas.Celeritas, userRepo data.User, logRepo data.Log) AuthService {
	return &authServiceImpl{
		App:      app,
		userRepo: userRepo,
		logRepo:  logRepo,
		BaseService: BaseServiceImpl{
			App: app,
		},
	}
}

func (s *authServiceImpl) Login(loginInput dto.LoginInput) (*dto.LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(loginInput.Email)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo auth get by email")
	}

	matches, err := user.PasswordMatches(loginInput.Password)
	if err != nil || !matches {
		return nil, newErrors.Wrap(err, "repo auth password matches")
	}

	userToken, err := s.generateAndSaveToken(user.ID)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo auth generate token")
	}

	_, err = s.logRepo.Insert(data.Log{
		ChangedAt: time.Now(),
		UserID:    user.ID,
		Entity:    data.EntityLogin,
	})

	if err != nil {
		return nil, newErrors.Wrap(err, "repo auth insert log")
	}

	return &dto.LoginResponse{
		User:  *dto.ToUserResponseDTO(*user),
		Token: *userToken,
	}, nil
}

func (s *authServiceImpl) ValidatePin(id int, pinInput dto.ValidatePinInput) error {
	user, err := s.userRepo.Get(id)
	if err != nil {
		return newErrors.Wrap(err, "repo auth get")
	}

	pinNum, _ := strconv.Atoi(user.Pin)
	pinInputNum, _ := strconv.Atoi(pinInput.Pin)

	if pinNum != pinInputNum {
		return newErrors.Wrap(nil, "repo auth pin match")
	}

	return nil
}

func (s *authServiceImpl) RefreshToken(userId int, refreshToken string, iat string) (*jwtdto.Token, error) {

	t, err := s.App.Cache.Get(buildRefreshTokenKey(userId, iat))
	if err != nil || t != refreshToken {
		return nil, newErrors.Wrap(err, "repo auth build token")
	}

	err = s.revokeRefreshToken(userId, iat)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo auth revoke token")
	}

	newToken, err := s.generateAndSaveToken(userId)
	if err != nil {
		return nil, newErrors.Wrap(err, "repo auth generate token")
	}

	return newToken, nil
}

func (s *authServiceImpl) Logout(userId int) error {
	err := s.revokeAllRefreshTokens(userId)
	if err != nil {
		return newErrors.Wrap(err, "repo auth revoke token")
	}

	_, err = s.logRepo.Insert(data.Log{
		ChangedAt: time.Now(),
		UserID:    userId,
		Entity:    data.EntityLogout,
	})

	if err != nil {
		return newErrors.Wrap(err, "repo auth log insert")
	}

	return nil
}

// RandomCharset generates a random character from the provided charset.
func RandomCharset(charset string) (string, error) {
	max := big.NewInt(int64(len(charset)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", newErrors.Wrap(err, "repo auth random")
	}
	return string(charset[n.Int64()]), nil
}

func (s *authServiceImpl) ForgotPassword(input dto.ForgotPassword) error {
	// verify that supplied email exists
	var u *data.User
	u, err := u.GetByEmail(input.Email)
	if err != nil {
		return newErrors.Wrap(err, "repo auth get by email")
	}

	// create and sign the link to password reset form
	link := s.buildPasswordResetLink(input.Email, "")
	sign := urlsigner.Signer{
		Secret: []byte(s.App.EncryptionKey),
	}
	signedLink := sign.GenerateTokenFromString(link)

	var data struct {
		Link string
	}
	data.Link = signedLink

	msg := mailer.Message{
		To:       u.Email,
		Subject:  "=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte("Zahtjev za izmjenu lozinke")) + "?=",
		Template: "password-reset",
		Data:     data,
	}

	s.App.Mail.Jobs <- msg
	res := <-s.App.Mail.Results
	if res.Error != nil {
		return newErrors.Wrap(err, "repo auth insert mail result")
	}

	return nil
}

func (s *authServiceImpl) ResetPasswordVerify(email, token string) (*dto.ResetPasswordVerifyResponse, error) {
	link := s.buildPasswordResetLink(email, token)

	// validate the url
	signer := urlsigner.Signer{
		Secret: []byte(s.App.EncryptionKey),
	}
	valid := signer.VerifyToken(link)
	if !valid {
		return nil, newErrors.Wrap(nil, "repo auth verify token")
	}

	// make sure it's not expired
	expired := signer.Expired(link, 60)
	if expired {
		return nil, newErrors.Wrap(nil, "repo auth expired")
	}

	encryptedEmail, _ := s.Encrypt(email)
	var response dto.ResetPasswordVerifyResponse
	response.EncryptedEmail = encryptedEmail

	return &response, nil
}

func (s *authServiceImpl) ResetPassword(input dto.ResetPassword) error {
	email, err := s.Decrypt(input.EncryptedEmail)
	if err != nil {
		return newErrors.Wrap(err, "repo auth decrypt")
	}

	// get the user
	var u data.User
	user, err := u.GetByEmail(email)
	if err != nil {
		return newErrors.Wrap(err, "repo auth get by email")
	}

	// reset the password
	err = user.ResetPassword(user.ID, input.Password)
	if err != nil {
		return newErrors.Wrap(err, "repo auth reset password")
	}

	return nil
}

func (s *authServiceImpl) buildPasswordResetLink(email string, hash string) string {
	email = url.QueryEscape(email)
	hash = url.QueryEscape(hash)

	if hash != "" {
		return fmt.Sprintf("%s/reset-password?email=%s&hash=%s", s.App.Frontend.URL, email, hash)
	}
	return fmt.Sprintf("%s/reset-password?email=%s", s.App.Frontend.URL, email)
}

func (s *authServiceImpl) generateAndSaveToken(userID int) (*jwtdto.Token, error) {
	userToken, err := s.App.JwtToken.Sign(jwt.MapClaims{
		"id": userID,
	})

	if err != nil {
		return nil, newErrors.Wrap(err, "repo auth jwt token sign")
	}

	_ = s.App.Cache.Set(
		buildRefreshTokenKey(userID, userToken.RefreshToken.Iat),
		userToken.RefreshToken.Value,
		int(s.App.JwtToken.JwtRefreshTokenTimeExp.Nanoseconds()),
	)

	return userToken, nil
}

func buildRefreshTokenKey(userID int, issuedAt string) string {
	return fmt.Sprintf("refresh_token_%d_%s", userID, issuedAt)
}

func (s *authServiceImpl) revokeRefreshToken(userID int, iat string) error {
	err := s.App.Cache.Forget(
		buildRefreshTokenKey(userID, iat),
	)
	return newErrors.Wrap(err, "repo auth cache forget")
}

func (s *authServiceImpl) revokeAllRefreshTokens(userID int) error {
	err := s.App.Cache.EmptyByMatch(
		buildRefreshTokenKey(userID, ""),
	)
	return newErrors.Wrap(err, "repo auth empty cache")
}
