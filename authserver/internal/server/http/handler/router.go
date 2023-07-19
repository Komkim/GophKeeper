package handler

import (
	"authserver/config"
	"authserver/internal/service"
	"authserver/pkg/logging"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

type Handler struct {
	config  *config.Config
	service service.Service
	log     logging.Logger
	//tokenAuth   *jwtauth.JWTAuth
}

func NewHandler(cfg *config.Config, service service.Service, log logging.Logger, secret string) *Handler {
	tokenAuth := jwtauth.New("HS256", []byte(secret), nil)
	return &Handler{
		config:    cfg,
		service:   service,
		log:       log,
		tokenAuth: tokenAuth,
	}
}

func (h *Handler) Init() http.Handler {

	handler := chi.NewRouter()
	handler.Use(middleware.Logger)
	handler.Use(middleware.Recoverer)
	handler.Use(middleware.Compress(5))

	handler.Use(middleware.Timeout(60 * time.Second))

	handler.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	handler.Route("/auth", func(r chi.Router) {
		r.Post("/register", h.SignUpUser)
		r.Post("/login", h.SignInUser)
		r.Get("/refresh", h.RefreshAccessToken)

		handler.Group(func(r chi.Router) {
			//r.Use(jwtauth.Verifier(h.tokenAuth))
			r.Use(h.DeserializeUser)
			r.Get("/logout", h.LogoutUser)
		})
	})

	return handler
}
