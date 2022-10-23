package handler

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

func (h *Handler) AddURL(w http.ResponseWriter, r *http.Request) {

	//var urlInfo urls.UrlInfo
	//json.NewDecoder(r.Body).Decode(&urlInfo)
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
	fmt.Println(shortURL)

	//w.Write([]byte(fmt.Sprintf("URL added successfully, short version: %s", shortURL)))
	//w.WriteHeader(201)

	//h.ShortURLPage(w, r, shortURL)
	http.Redirect(w, r, shortURL, http.StatusSeeOther)
	//	http.Redirect(w, r, "https://hub.docker.com/_/mongo", 300)
	fmt.Println("f")
}

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

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("template/index.html")
	tmpl.Execute(w, nil)
}

func (h *Handler) ShortURLPage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("template/index_short.html")
	shortURL := chi.URLParam(r, "shortURL")
	tmpl.ExecuteTemplate(w, "shortURL", shortURL)
}
