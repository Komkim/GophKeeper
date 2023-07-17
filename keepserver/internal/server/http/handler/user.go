package handler

import (
	"GophKeeper/keepserver/internal/server/http/handler/request"
	"GophKeeper/keepserver/internal/server/http/handler/response"
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) UserLoading(w http.ResponseWriter, r *http.Request) {
	data := &request.User{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.ErrInvalidRequestFormat(err))
		h.log.Error(err)
		return
	}

	_, err := h.service.SetUser(
		data.Login,
		data.Password,
		data.CliCreation,
	)

	if err != nil {
		render.Render(w, r, response.ErrInternalServer(err))
		h.log.Error(err)
		return
	}

	render.Render(w, r, response.UserSuccessfully)
}

func (h *Handler) UserUpdate(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UserGet(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UserGetAll(w http.ResponseWriter, r *http.Request) {

}
