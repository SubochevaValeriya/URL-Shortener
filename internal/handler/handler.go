package handler

import (
	"URLShortener/internal/service"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"strings"
	//_ "github.com/SubochevaValeriya/URL-Shotener/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func RegisterRoutes(h *Handler) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("template"))
	fmt.Println(fs)
	r.Handle("/templates/*", fs)
	r.Route("/", func(r chi.Router) {
		r.Get("/urls/{shortURL}", h.GetURL)
		r.Get("/", h.MainPage)
		r.Post("/new", h.AddURL) //POST /contacts
		r.Get("/new2", h.ShortURLPage)
	})

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}

	return r
}
