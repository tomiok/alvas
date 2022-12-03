package user

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
	"github.com/tomiok/alvas/pkg/web"
	"net/http"
)

type KeyUserID string

const KeyUID KeyUserID = "userID"

type JWTAuthService struct {
	Secret string
}

type CreateJWTPayload struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func NewJWTAuthService(secret string) *JWTAuthService {
	return &JWTAuthService{
		Secret: secret,
	}
}

// AdminMiddleware middleware to protect routes from users without the "admin" role
func (j *JWTAuthService) AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			log.Warn().Msgf("%s - no token given", web.Trace())
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		valid, err := j.ValidateJWT(token)

		if err != nil {
			log.Warn().Msgf("%s - %s", web.Trace(), err.Error())
			web.ResponseUnauthorized(w, "unauthorized - invalid token")
			return
		}

		if !valid {
			log.Warn().Msgf("%s - token invalid", web.Trace())
			web.ResponseUnauthorized(w, "unauthorized - invalid token")
			return
		}

		claims, err := j.DecodeJWT(token)

		if err != nil {
			log.Warn().Msgf("%s - %s", web.Trace(), err.Error())
			web.ResponseUnauthorized(w, "unauthorized - cannot decode token")
			return
		}

		if claims.Role != "admin" {
			log.Warn().Msgf("%s - %s is not an admin", web.Trace(), claims.Email)
			web.ResponseUnauthorized(w, "unauthorized - not an admin token")
			return
		}

		ctx := context.WithValue(r.Context(), KeyUID, claims.Id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (j *JWTAuthService) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			log.Warn().Msgf("%s - no token given", web.Trace())
			web.ResponseUnauthorized(w, "unauthorized - no token given")
			return
		}

		valid, err := j.ValidateJWT(token)

		if err != nil {
			log.Warn().Msgf("%s - %s", web.Trace(), err.Error())
			web.ResponseUnauthorized(w, "unauthorized - invalid token")
			return
		}

		if !valid {
			log.Warn().Msgf("%s - token invalid", web.Trace())
			web.ResponseUnauthorized(w, "unauthorized - invalid token")
			return
		}

		claims, err := j.DecodeJWT(token)
		if err != nil {
			log.Warn().Msgf("%s - token invalid", web.Trace())
			web.ResponseUnauthorized(w, "unauthorized - invalid token")
			return
		}

		userID := chi.URLParam(r, string(KeyUID))

		if userID != "" {
			if claims.Id != userID {
				log.Warn().Msgf("%s - token invalid", web.Trace())
				web.ResponseUnauthorized(w, "unauthorized - mismatch in user in URL and token")
				return
			}
		}

		ctx := context.WithValue(r.Context(), KeyUID, claims.Id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (j *JWTAuthService) DecodeJWT(token string) (*CustomClaims, error) {
	claims := &CustomClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (j *JWTAuthService) ValidateJWT(token string) (bool, error) {
	tokenData, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(j.Secret), nil
	})

	if err != nil {
		return false, err
	}
	if _, ok := tokenData.Claims.(*CustomClaims); ok && tokenData.Valid {
		return true, nil
	} else {
		return false, nil
	}
}

type CustomClaims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
