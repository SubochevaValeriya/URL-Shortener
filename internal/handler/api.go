package handler

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

// AddURL makes short URL for original URL
func (h *Handler) AddURL(w http.ResponseWriter, r *http.Request) {

	originalURL := r.FormValue("longURL")
	if originalURL == "" {
		http.Error(w, "empty request", 400)
		return
	}

	shortURL, err := h.services.AddURL(originalURL)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		return
	}

	//w.Write([]byte(fmt.Sprintf("URL added successfully, short version: %s", shortURL)))
	//w.WriteHeader(201)

	//h.ShortURLPage(w, r, shortURL)
	http.Redirect(w, r, shortURL, http.StatusSeeOther)

}

// GetURL redirects to original URL
func (h *Handler) GetURL(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")
	if shortURL == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	originalURL, err := h.services.GetURL(shortURL)

	if err != nil {
		http.Error(w, fmt.Sprintf("Original version of URL: %s not found", shortURL), 404)
		return
	}

	if string(originalURL[:4]) != "http" {
		originalURL = "https://" + originalURL
	}

	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}

// MainPage is main page of the site
func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("template/index.html")
	tmpl.Execute(w, nil)
}

// ShortURLPage is the page when user can copy short URL
func (h *Handler) ShortURLPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("template/index_short.html")
	shortURL := chi.URLParam(r, "shortURL")
	tmpl.ExecuteTemplate(w, "shortURL", shortURL)
}
