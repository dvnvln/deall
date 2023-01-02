package handler

import (
	"log"
	"net/http"

	"github.com/dvnvln/deallscrud/model"
	"github.com/dvnvln/deallscrud/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// Request Handler - POST /posts - Create a new user.
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	data := &model.CreateUserReqBody{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrorRenderer(err))
		return
	}
	if err := data.Validate(); err != nil {
		render.Render(w, r, utils.ErrorRenderer(err))
		return
	}
	res, err := h.controller.CreateUser(r.Context(), *data)
	if err != nil {
		render.Render(w, r, utils.ServerErrorRenderer(err))
		return
	}
	render.JSON(w, r, res)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	res, err := h.controller.ListUser(r.Context())
	if err != nil {
		render.Render(w, r, utils.ErrorRenderer(err))
		return
	}
	render.JSON(w, r, res)
	log.Println("handler-handler-")
}

func (h *Handler) Read(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	res, err := h.controller.GetUser(r.Context(), userID)
	if err != nil {
		render.Render(w, r, utils.ErrorRenderer(err))
		return
	}
	render.JSON(w, r, res)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	data := &model.CreateUserReqBody{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrorRenderer(err))
		return
	}
	if err := data.Validate(); err != nil {
		render.Render(w, r, utils.ErrorRenderer(err))
		return
	}
	res, err := h.controller.UpdateUser(r.Context(), userID, *data)
	if err != nil {
		render.Render(w, r, utils.ErrorRenderer(err))
		return
	}
	render.JSON(w, r, res)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	res, err := h.controller.DeleteUser(r.Context(), userID)
	if err != nil {
		render.Render(w, r, utils.ErrorRenderer(err))
		return
	}
	render.JSON(w, r, res)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	log.Print("handler here..")
	data := &model.UserReqBody{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrorRenderer(err))
		return
	}
	if err := data.Validate(); err != nil {
		render.Render(w, r, utils.ErrorRenderer(err))
		return
	}
	log.Print("Calling controller..")
	res, err := h.controller.UserLogin(r.Context(), *data)
	if err != nil {
		render.Render(w, r, utils.ServerErrorRenderer(err))
		return
	}
	render.JSON(w, r, res)

	log.Print("handler done..")
}
