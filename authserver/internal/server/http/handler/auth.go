package handler

import (
	"authserver/internal/server/http/handler/request"
	"authserver/internal/server/http/handler/response"
	"authserver/pkg/token"
	"errors"
	"github.com/go-chi/render"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

//func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
//	http.SetCookie(w, &http.Cookie{
//		HttpOnly: true,
//		MaxAge: -1, // Delete the cookie.
//		SameSite: http.SameSiteLaxMode,
//		Name:  "jwt",
//		Value: "",
//	})
//
//	render.Render(w, r, response.UserLogout)
//}
//
//func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
//	err := r.ParseForm()
//	if err != nil{
//		render.Render(w, r, response.ErrInternalServer(err))
//		h.log.Error(err)
//		return
//	}
//	userName := r.PostForm.Get("username")
//	userPassword := r.PostForm.Get("password")
//
//	if userName == "" || userPassword == "" {
//		render.Render(w, r, response.ErrInvalidRequestFormat(err))
//		h.log.Errorf("Missing username or password. Err: %s", err)
//		return
//	}
//
//	_, token, err := h.tokenAuth.Encode(map[string]interface{}{"username": userName})
//	if err != nil{
//		render.Render(w, r, response.ErrInternalServer(err))
//		h.log.Errorf("Token auth encode err: %s", err)
//		return
//	}
//
//	http.SetCookie(w, &http.Cookie{
//		HttpOnly: true,
//		Expires: time.Now().Add(7 * 24 * time.Hour),
//		SameSite: http.SameSiteLaxMode,
//		Name:  "jwt",
//		Value: token,
//	})
//	render.Render(w, r, response.UserAuthenticated)
//}
//
//func (h *Handler) Autentification(w http.ResponseWriter, r *http.Request) {
//
//}

func (h *Handler) SignUpUser(w http.ResponseWriter, r *http.Request) {
	data := &request.User{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.ErrInvalidRequestFormat(err))
		h.log.Error(err)
		return
	}

	//if data.Password != data.PasswordConfirm {
	//	render.Render(w, r, response.ErrStatusBadRequest)
	//	h.log.Error(errors.New("Status bad request"))
	//	return
	//}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		render.Render(w, r, response.ErrStatusBadGateway)
		h.log.Error(errors.New(response.ErrStatusBadGateway.StatusText))
		return
	}

	resp := h.service.CreateUser(data.Login, string(hashedPassword), data.CliCreation)

	render.Render(w, resp.Request, response.StatusCreated)
}

func (h *Handler) SignInUser(w http.ResponseWriter, r *http.Request) {
	data := &request.User{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.ErrInvalidRequestFormat(err))
		h.log.Error(err)
		return
	}

	user, err := h.service.GetUserByLogin(data.Login)
	if err != nil {
		render.Render(w, r, response.ErrStatusForbidden)
		h.log.Error(errors.New(response.ErrStatusForbidden.StatusText))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(hashedPassword))
	if err != nil {
		render.Render(w, r, response.ErrStatusBadGateway)
		h.log.Error(errors.New(response.ErrStatusBadGateway.StatusText))
		return
	}

	accessTokenDetails, err := token.CreateToken(user.ID.String(), h.config.Token.AccessTokenExpiresIn, h.config.Token.AccessTokenPrivateKey)
	if err != nil {
		render.Render(w, r, response.ErrStatusUnprocessableEntity)
		h.log.Error(errors.New(response.ErrStatusUnprocessableEntity.StatusText))
		return
	}

	refreshTokenDetails, err := token.CreateToken(user.ID.String(), h.config.Token.RefreshTokenExpiresIn, h.config.Token.RefreshTokenPrivateKey)
	if err != nil {
		render.Render(w, r, response.ErrStatusUnprocessableEntity)
		h.log.Error(errors.New(response.ErrStatusUnprocessableEntity.StatusText))
		return
	}

	err = h.service.SetToken(accessTokenDetails.TokenUuid, user.ID.String(), accessTokenDetails.ExpiresIn)
	if err != nil {
		render.Render(w, r, response.ErrStatusUnprocessableEntity)
		h.log.Error(errors.New(response.ErrStatusUnprocessableEntity.StatusText))
		return
	}

	err = h.service.SetToken(refreshTokenDetails.TokenUuid, user.ID.String(), accessTokenDetails.ExpiresIn)
	if err != nil {
		render.Render(w, r, response.ErrStatusUnprocessableEntity)
		h.log.Error(errors.New(response.ErrStatusUnprocessableEntity.StatusText))
		return
	}

	r.AddCookie(&http.Cookie{
		Name:     "access_token",
		Value:    *accessTokenDetails.Token,
		Path:     "/",
		MaxAge:   h.config.Token.AccessTokenMaxAge * 60,
		Secure:   false,
		HttpOnly: true,
		Domain:   "localhost",
	})

	r.AddCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    *refreshTokenDetails.Token,
		Path:     "/",
		MaxAge:   h.config.Token.RefreshTokenMaxAge * 60,
		Secure:   false,
		HttpOnly: true,
		Domain:   "localhost",
	})

	r.AddCookie(&http.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   h.config.Token.AccessTokenMaxAge * 60,
		Secure:   false,
		HttpOnly: false,
		Domain:   "localhost",
	})

	render.Render(w, r, response.UserSignInOk)
}

func (h *Handler) RefreshAccessToken(w http.ResponseWriter, r *http.Request) {
	refresh_token, err := r.Cookie("refresh_token")
	if err != nil || refresh_token == nil {
		render.Render(w, r, response.ErrStatusForbidden)
		h.log.Error(errors.New(response.ErrStatusForbidden.StatusText))
		return
	}

	tokenClaims, err := token.ValidateToken(refresh_token.Value, h.config.Token.RefreshTokenPublicKey)
	if err != nil {
		render.Render(w, r, response.ErrStatusForbidden)
		h.log.Error(errors.New(response.ErrStatusForbidden.StatusText))
		return
	}

	userID, err := h.service.GetToken(tokenClaims.TokenUuid)
	if err != redis.Nil {
		render.Render(w, r, response.ErrStatusForbidden)
		h.log.Error(errors.New(response.ErrStatusForbidden.StatusText))
		return
	}

	user, err := h.service.GetUserById(userID)
	if err != nil {
		render.Render(w, r, response.ErrStatusForbidden)
		h.log.Error(errors.New(response.ErrStatusForbidden.StatusText))
		return
	}

	accessTokenDetails, err := token.CreateToken(user.ID.String(), h.config.Token.AccessTokenExpiresIn, h.config.Token.AccessTokenPrivateKey)
	if err != nil {
		if err != nil {
			render.Render(w, r, response.ErrStatusUnprocessableEntity)
			h.log.Error(errors.New(response.ErrStatusUnprocessableEntity.StatusText))
			return
		}
	}

	err = h.service.SetToken(accessTokenDetails.TokenUuid, user.ID.String(), accessTokenDetails.ExpiresIn)
	if err != nil {
		render.Render(w, r, response.ErrStatusUnprocessableEntity)
		h.log.Error(errors.New(response.ErrStatusUnprocessableEntity.StatusText))
		return
	}

	r.AddCookie(&http.Cookie{
		Name:     "access_token",
		Value:    *accessTokenDetails.Token,
		Path:     "/",
		MaxAge:   h.config.Token.AccessTokenMaxAge * 60,
		Secure:   false,
		HttpOnly: true,
		Domain:   "localhost",
	})

	r.AddCookie(&http.Cookie{
		Name:     "logged_in",
		Value:    "true",
		Path:     "/",
		MaxAge:   h.config.Token.AccessTokenMaxAge * 60,
		Secure:   false,
		HttpOnly: false,
		Domain:   "localhost",
	})

	render.Render(w, r, response.RefreshAccessToken)
}

func (h *Handler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	refresh_token, err := r.Cookie("refresh_token")
	if err != nil || refresh_token == nil {
		render.Render(w, r, response.ErrStatusForbidden)
		h.log.Error(errors.New(response.ErrStatusForbidden.StatusText))
		return
	}

	access_token, err := r.Cookie("access_token")
	if err != nil || refresh_token == nil {
		render.Render(w, r, response.ErrStatusForbidden)
		h.log.Error(errors.New(response.ErrStatusForbidden.StatusText))
		return
	}

	tokenClaims, err := token.ValidateToken(refresh_token.Value, h.config.Token.RefreshTokenPublicKey)
	if err != nil {
		if err != nil {
			render.Render(w, r, response.ErrStatusForbidden)
			h.log.Error(errors.New(response.ErrStatusForbidden.StatusText))
			return
		}
	}

	err = h.service.DelToken(tokenClaims.TokenUuid)
	if err != nil {
		render.Render(w, r, response.ErrStatusBadGateway)
		h.log.Error(errors.New(response.ErrStatusBadGateway.StatusText))
		return
	}

	err = h.service.DelToken(access_token.Value)
	if err != nil {
		render.Render(w, r, response.ErrStatusBadGateway)
		h.log.Error(errors.New(response.ErrStatusBadGateway.StatusText))
		return
	}

	expired := time.Now().Add(-time.Hour * 24)
	r.AddCookie(&http.Cookie{
		Name:    "access_token",
		Value:   "",
		Expires: expired,
	})
	r.AddCookie(&http.Cookie{
		Name:    "refresh_token",
		Value:   "",
		Expires: expired,
	})
	r.AddCookie(&http.Cookie{
		Name:    "logged_in",
		Value:   "",
		Expires: expired,
	})
	render.Render(w, r, response.Logout)
}
