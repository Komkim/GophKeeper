package handler

import (
	"GophKeeper/keepserver/internal/service"
	"GophKeeper/keepserver/pkg/logging"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

type Handler struct {
	service service.Service
	log     logging.Logger
}

func NewHandler(service service.Service, log logging.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
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

	handler.Route("/api", func(r chi.Router) {

		handler.Route("/file", func(r chi.Router) {
			r.Post("/", h.FileLoading)
			r.Get("/", h.FileGetAll)
			r.Get("/:id", h.FileGet)
		})

		handler.Route("/card", func(r chi.Router) {
			r.Post("/", h.CardLoading)
			r.Get("/", h.CardGetAll)
			r.Get("/:id", h.CardGet)
		})

		handler.Route("/note", func(r chi.Router) {
			r.Post("/", h.NoteLoading)
			r.Get("/", h.NoteGetAll)
			r.Get("/:id", h.NoteGet)
		})

		handler.Route("/user", func(r chi.Router) {
			r.Post("/", h.UserLoading)
			//r.Get("/", h.NoteGetAll)
			r.Get("/", h.UserGet)
		})
	})

	return handler
}
