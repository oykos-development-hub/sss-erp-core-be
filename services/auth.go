package services

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/url"
	"strconv"
	"strings"

	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"
	"golang.org/x/crypto/bcrypt"

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
}

func NewAuthServiceImpl(app *celeritas.Celeritas, userRepo data.User) AuthService {
	return &authServiceImpl{
		App:      app,
		userRepo: userRepo,
		BaseService: BaseServiceImpl{
			App: app,
		},
	}
}

func (s *authServiceImpl) Login(loginInput dto.LoginInput) (*dto.LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(loginInput.Email)
	if err != nil {
		s.App.ErrorLog.Println(err.Error())
		return nil, errors.ErrNotFound
	}

	matches, err := user.PasswordMatches(loginInput.Password)
	if err != nil || !matches {
		return nil, errors.ErrBadRequest
	}

	userToken, err := s.generateAndSaveToken(user.ID)
	if err != nil {
		s.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}

	return &dto.LoginResponse{
		User:  *dto.ToUserResponseDTO(*user),
		Token: *userToken,
	}, nil
}

func (s *authServiceImpl) ValidatePin(id int, pinInput dto.ValidatePinInput) error {
	user, err := s.userRepo.Get(id)
	if err != nil {
		s.App.ErrorLog.Println(err)
		return errors.ErrNotFound
	}

	pinNum, _ := strconv.Atoi(user.Pin)
	pinInputNum, _ := strconv.Atoi(pinInput.Pin)

	if pinNum != pinInputNum {
		return errors.ErrUnauthorized
	}

	return nil
}

func (s *authServiceImpl) RefreshToken(userId int, refreshToken string, iat string) (*jwtdto.Token, error) {

	t, err := s.App.Cache.Get(buildRefreshTokenKey(userId, iat))
	if err != nil || t != refreshToken {
		s.App.ErrorLog.Printf("Refresh token is revoked: %v", err)
		return nil, errors.ErrUnauthorized
	}

	err = s.revokeRefreshToken(userId, iat)
	if err != nil {
		s.App.ErrorLog.Printf("Error rotating refresh tokens: %v", err)
		return nil, errors.ErrInternalServer
	}

	newToken, err := s.generateAndSaveToken(userId)
	if err != nil {
		s.App.ErrorLog.Printf("Error generating new refresh token: %v", err)
		return nil, errors.ErrInternalServer
	}

	return newToken, nil
}

func (s *authServiceImpl) Logout(userId int) error {
	err := s.revokeAllRefreshTokens(userId)
	if err != nil {
		s.App.ErrorLog.Printf("Error revoking refresh token: %v", err)
		return errors.ErrUnauthorized
	}

	return nil
}

// RandomCharset generates a random character from the provided charset.
func RandomCharset(charset string) (string, error) {
	max := big.NewInt(int64(len(charset)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return string(charset[n.Int64()]), nil
}

// GeneratePassword generates a random password with at least one character from each category.
func GeneratePassword(length int) (string, error) {
	if length < 6 {
		return "", fmt.Errorf("password length must be at least 6 characters")
	}

	const (
		lowercase = "abcdefghijklmnopqrstuvwxyz"
		uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		digits    = "0123456789"
		special   = "!@#$%^&*()-_+=[]{}|;:,.<>?"
	)

	password := make([]string, length)
	charsets := []string{lowercase, uppercase, digits, special}

	// Ensure each category is represented
	for i, set := range charsets {
		char, err := RandomCharset(set)
		if err != nil {
			return "", err
		}
		password[i] = char
	}

	// Fill the rest of the password with random characters from all categories
	for i := len(charsets); i < length; i++ {
		char, err := RandomCharset(lowercase + uppercase + digits + special)
		if err != nil {
			return "", err
		}
		password[i] = char
	}

	// Shuffle the password to prevent predictable sequences
	for i := range password {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			return "", err
		}
		password[i], password[j.Int64()] = password[j.Int64()], password[i]
	}

	return strings.Join(password, ""), nil
}

func (s *authServiceImpl) ForgotPasswordV2(input dto.ForgotPassword) error {
	// verify that supplied email exists
	var u *data.User
	u, err := u.GetByEmail(input.Email)
	if err != nil {
		return errors.ErrNotFound
	}

	password, err := GeneratePassword(8)
	if err != nil {
		s.App.ErrorLog.Printf("Error revoking refresh token: %v", err)
		return errors.ErrInternalServer
	}

	newHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		s.App.ErrorLog.Printf("Error generating hash from password: %v", err)
		return errors.ErrInternalServer
	}
	s.App.InfoLog.Println("TEST1")

	u.Password = string(newHash)

	err = u.Update(*u)
	if err != nil {
		s.App.ErrorLog.Printf("Error updating user: %v", err)
		return errors.ErrInternalServer
	}
	s.App.InfoLog.Println("TEST2")

	// email the message
	var data struct {
		Password string
	}
	data.Password = password

	msg := mailer.Message{
		To:       u.Email,
		Subject:  "Predmet: Instrukcije za resetovanje lozinke",
		Template: "password-reset",
		Data:     data,
	}
	s.App.InfoLog.Println("TEST3")

	s.App.Mail.Jobs <- msg
	res := <-s.App.Mail.Results
	if res.Error != nil {
		return err
	}

	s.App.InfoLog.Println(u.Email, res.Error, res.Success)

	return nil
}

func (s *authServiceImpl) ForgotPassword(input dto.ForgotPassword) error {
	// verify that supplied email exists
	var u *data.User
	u, err := u.GetByEmail(input.Email)
	if err != nil {
		return errors.ErrNotFound
	}

	// create and sign the link to password reset form
	link := s.buildPasswordResetLink(input.Email, "")
	sign := urlsigner.Signer{
		Secret: []byte(s.App.EncryptionKey),
	}
	signedLink := sign.GenerateTokenFromString(link)

	// email the message
	var data struct {
		Link string
	}
	data.Link = signedLink

	msg := mailer.Message{
		To:       u.Email,
		Subject:  "Password reset",
		Template: "password-reset",
		Data:     data,
	}

	s.App.Mail.Jobs <- msg
	res := <-s.App.Mail.Results
	if res.Error != nil {
		return err
	}

	return nil
}

func (s *authServiceImpl) ResetPasswordVerify(email, token string) (bool, error) {
	link := s.buildPasswordResetLink(email, token)

	// validate the url
	signer := urlsigner.Signer{
		Secret: []byte(s.App.EncryptionKey),
	}
	valid := signer.VerifyToken(link)
	if !valid {
		return false, errors.ErrUnauthorized
	}

	// make sure it's not expired
	expired := signer.Expired(link, 60)
	if expired {
		return false, errors.ErrExpired
	}

	return true, nil
}

func (s *authServiceImpl) ResetPassword(input dto.ResetPassword) error {
	// get the user
	var u data.User
	user, err := u.GetByEmail(input.Email)
	if err != nil {
		s.App.ErrorLog.Printf("Failed to retrieve user: %v", err)
		return errors.ErrInternalServer
	}

	// reset the password
	err = user.ResetPassword(user.ID, input.Password)
	if err != nil {
		s.App.ErrorLog.Printf("Failed to reset password: %v", err)
		return errors.ErrInternalServer
	}

	return nil
}

func (s *authServiceImpl) buildPasswordResetLink(email string, hash string) string {
	email = url.QueryEscape(email)
	hash = url.QueryEscape(hash)

	if hash != "" {
		return fmt.Sprintf("%s/reset-password?email=%s&token=%s", s.App.Frontend.URL, email, hash)
	}
	return fmt.Sprintf("%s/reset-password?email=%s", s.App.Frontend.URL, email)
}

func (s *authServiceImpl) generateAndSaveToken(userID int) (*jwtdto.Token, error) {
	userToken, err := s.App.JwtToken.Sign(jwt.MapClaims{
		"id": userID,
	})

	if err != nil {
		return nil, err
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
	return err
}

func (s *authServiceImpl) revokeAllRefreshTokens(userID int) error {
	err := s.App.Cache.EmptyByMatch(
		buildRefreshTokenKey(userID, ""),
	)
	return err
}
