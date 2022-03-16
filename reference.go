package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func referenceRouter(r chi.Router) {
	r.Get("/", referenceMain)
}

func referenceMain(w http.ResponseWriter, r *http.Request) {
	tmp := newTemplate("templates/reference.gohtml")
	tmp.ExecuteTemplate(w, "base", "")
}

func playbookContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// characterID := chi.URLParam(r, "characterID")
		// character, err := characters.GetCharacterByIDString(lionturtledb, characterID)
		// if err != nil {
		// 	http.Error(w, http.StatusText(404), 404)
		// 	return
		// }
		// ctx := context.WithValue(r.Context(), "character", &character)
		// next.ServeHTTP(w, r.WithContext(ctx))
	})
}
