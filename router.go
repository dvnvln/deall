package main

import (
	"net/http"
	"os"

	"github.com/dvnvln/deallscrud/handler"
	jwtmid "github.com/dvnvln/deallscrud/middleware"
	"github.com/dvnvln/deallscrud/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func initRoute(handler *handler.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello World! sangchi"))
	})
	r.Route("/users", func(r chi.Router) {
		r.Use(jwtauth.Verifier(utils.GetToken(os.Getenv("SECRET")).Auth))
		r.Use(jwtauth.Authenticator)
		r.Use(jwtmid.AuthMiddleware)
		r.Get("/", handler.List)
		r.Get("/{userID}", handler.Read)

		r.Group(func(r chi.Router) {
			r.Use(jwtmid.CheckIsAdmin)
			r.Post("/", handler.Create)
			r.Route("/{userID}", func(r chi.Router) {
				r.Put("/", handler.Update)
				r.Delete("/", handler.Delete)
			})
		})
	})
	r.Route("/login", func(r chi.Router) {
		r.Post("/", handler.Login)
	})

	return r
}
