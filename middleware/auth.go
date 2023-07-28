package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/request"
	"github.com/oykos-development-hub/celeritas/rsa"
	"gitlab.sudovi.me/erp/core-ms-api/errors"
)

// JwtVerifyToken usefull for middleware for verify the jwt token from the Authorization
// this function will serve to middleware and usefull for the idiomatic framework like gorm or chi or just net/http
func (m *Middleware) JwtVerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		JwtToken := strings.Replace(r.Header.Get("Authorization"), fmt.Sprintf("%s ", "Bearer"), "", 1)

		if JwtToken == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			tokenType := token.Claims.(jwt.MapClaims)["token_type"]

			if tokenType != "access_token" {
				return nil, fmt.Errorf("unexpected token type: %v", tokenType)
			}

			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			publicRsa, err := rsa.ReadPublicKeyFromEnv(m.App.JwtToken.RSAPublic)
			if err != nil {
				return nil, err
			}
			return publicRsa, nil
		})

		if err != nil {
			m.App.ErrorLog.Println(err)
		}

		if err != nil || !token.Valid {
			_ = m.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(errors.ErrUnauthorized), errors.ErrUnauthorized, nil)
			return
		}

		rawId := token.Claims.(jwt.MapClaims)["id"].(float64)
		id := fmt.Sprintf("%d", int(rawId))
		if id == "" {
			m.App.ErrorLog.Println("Token not found")
			_ = m.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(errors.ErrUnauthorized), errors.ErrUnauthorized, nil)
			return
		}

		rawExp := token.Claims.(jwt.MapClaims)["exp"].(float64)
		exp := int64(rawExp)
		if exp < time.Now().Unix() {
			_ = m.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(errors.ErrUnauthorized), errors.ErrUnauthorized, nil)
			return
		}

		r.Header.Set("id", id)

		next.ServeHTTP(w, r)
	})
}

// JwtVerifyRefreshToken usefull for middleware for verify the jwt refresh token from the Authorization
// this function will serve to middleware and usefull for the idiomatic framework like gorm or chi or just net/http
func (m *Middleware) JwtVerifyRefreshToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		JwtTokenCookie, _ := r.Cookie("refresh_token")
		if JwtTokenCookie == nil {
			m.App.ErrorLog.Println("Refresh token is not in cookie")
			_ = m.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(errors.ErrUnauthorized), errors.ErrUnauthorized, nil)
			return
		}
		JwtToken := JwtTokenCookie.Value

		if JwtToken == "" {
			m.App.ErrorLog.Printf("Refresh token is empty")
			_ = m.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(errors.ErrUnauthorized), errors.ErrUnauthorized, nil)
			return
		}

		var err error
		token, err := jwt.Parse(JwtToken, func(token *jwt.Token) (interface{}, error) {
			tokenType := token.Claims.(jwt.MapClaims)["token_type"]
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			if tokenType != "refresh_token" {
				return nil, fmt.Errorf("unexpected token type: %v", tokenType)
			}

			privateRsa, err := rsa.ReadPublicKeyFromEnv(m.App.JwtToken.RSAPublic)
			if err != nil {
				m.App.ErrorLog.Printf("Error reading private key RSA from env: %v", err)
				return nil, err
			}
			return privateRsa, nil
		})

		if err != nil {
			m.App.ErrorLog.Println(err)
		}

		if err != nil || !token.Valid {
			_ = m.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(errors.ErrUnauthorized), errors.ErrUnauthorized, nil)
			return
		}

		idFloat := token.Claims.(jwt.MapClaims)["id"].(float64)
		id := fmt.Sprintf("%.0f", idFloat)

		if id == "" {
			_ = m.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(errors.ErrUnauthorized), errors.ErrUnauthorized, nil)
			return
		}

		rawExp := token.Claims.(jwt.MapClaims)["exp"].(float64)
		exp := int64(rawExp)
		if exp < time.Now().Unix() {
			_ = m.App.WriteErrorResponse(w, errors.MapErrorToStatusCode(errors.ErrUnauthorized), errors.ErrUnauthorized, nil)
			return
		}

		r.Header.Set("id", id)
		r.Header.Set("iat", fmt.Sprintf("%.0f", token.Claims.(jwt.MapClaims)["iat"].(float64)))
		r.Header.Set("refresh_token", token.Raw)

		next.ServeHTTP(w, r)
	})
}
