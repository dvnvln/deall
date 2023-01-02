package jwtmid

import (
	"context"
	"net/http"

	"github.com/dvnvln/deallscrud/model"
	"github.com/dvnvln/deallscrud/utils"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

func AuthMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			render.Render(w, r, utils.ErrorRenderer(err))
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, model.UserID, claims[model.UserID])
		ctx = context.WithValue(ctx, model.RoleName, claims[model.RoleName])
		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func CheckIsAdmin(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if ctx.Value(model.RoleName) != model.Admin {
			render.Render(w, r, utils.UnauthorizedError())
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
