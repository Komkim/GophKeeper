package handler

import (
	"authserver/internal/server/http/handler/response"
	"authserver/pkg/token"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strings"
)

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func LoggedInRedirector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		е, _, _ := jwtauth.FromContext(r.Context())

		if е != nil && jwt.Validate(е) == nil {
			http.Redirect(w, r, "/login", 302)
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) DeserializeUser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var access_token string
		authorization := r.Header.Get("Authorization")

		if strings.HasPrefix(authorization, "Bearer ") {
			access_token = strings.TrimPrefix(authorization, "Bearer ")
		} else {
			cookie, err := r.Cookie("access_token")
			if err != nil || cookie == nil {
				render.Render(w, r, response.ErrNotAuthenticated)
				h.log.Error(errors.New(response.ErrNotAuthenticated.StatusText))
				return
			}
			access_token = cookie.Value
		}

		tokenClaims, err := token.ValidateToken(access_token, h.config.Token.AccessTokenPublicKey)
		if err != nil {
			render.Render(w, r, response.ErrStatusForbidden)
			h.log.Error(errors.New(response.ErrStatusForbidden.StatusText))
			return
		}

		_, err = h.service.GetToken(tokenClaims.TokenUuid)
		if err != redis.Nil {
			render.Render(w, r, response.ErrStatusForbidden)
			h.log.Error(errors.New(response.ErrStatusForbidden.StatusText))
			return
		}

		r.Header.Set("access_token_uuid", tokenClaims.TokenUuid)

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
