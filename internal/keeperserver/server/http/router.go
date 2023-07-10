package router

import (
	"GophKeeper/config"
	"GophKeeper/internal/keeperserver/auth"
	"GophKeeper/internal/keeperserver/storage"
	"github.com/rs/zerolog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	cfg     *config.KeeperServer
	log     *zerolog.Event
	auth    auth.IAuth
	storage storage.IStorage
}

func NewRouter(cfg *config.KeeperServer, log *zerolog.Event, auth auth.IAuth, storage storage.IStorage) *Router {
	return &Router{
		cfg:     cfg,
		log:     log,
		auth:    auth,
		storage: storage,
	}
}

func (r *Router) Init() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Compress(5))

	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return router
}
