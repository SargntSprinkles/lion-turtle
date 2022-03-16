package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/SargntSprinkles/lion-turtle/characters"
	"github.com/SargntSprinkles/lion-turtle/playbooks"
	"github.com/SargntSprinkles/lion-turtle/util"
	"github.com/go-chi/chi/v5"
)

func characterRouter(r chi.Router) {
	r.Get("/create", createCharacter)
	r.Route("/{characterID}", func(r chi.Router) {
		r.Use(characterContext)
		r.Get("/", viewCharacter)
		r.Get("/edit", editCharacter)
		r.Post("/", updateCharacter)
		r.Get("/delete", deleteCharacter)
		r.Post("/delete/confirm", confirmDeleteCharacter)
	})
}

func characterContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		characterID := chi.URLParam(r, "characterID")
		character, err := characters.GetCharacterByIDString(characterID)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "character", &character)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createCharacter(w http.ResponseWriter, r *http.Request) {
	log.Print("Create character")
	newCharacter := characters.Character{Name: characters.RandomName()}
	log.Printf("Creating new character: %s", newCharacter.Name)
	newCharacter.Save()
	stringID := strconv.FormatUint(uint64(newCharacter.ID), 10)
	newURL := fmt.Sprintf("/character/%s/edit", stringID)
	http.Redirect(w, r, newURL, http.StatusSeeOther)
}

func viewCharacter(w http.ResponseWriter, r *http.Request) {
	viewData := map[string]interface{}{}
	theCharacter, ok := r.Context().Value("character").(*characters.Character)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	viewData["character"] = theCharacter
	viewData["highlight"] = "character"
	characterTemplate := newTemplate("templates/character.gohtml")
	characterTemplate.ExecuteTemplate(w, "base", viewData)
}

func editCharacter(w http.ResponseWriter, r *http.Request) {
	viewData := map[string]interface{}{}
	theCharacter, ok := r.Context().Value("character").(*characters.Character)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	viewData["character"] = theCharacter
	viewData["playbooks"] = playbooks.PlaybookNames
	viewData["highlight"] = "character"
	editTemplate := newTemplate("templates/character_edit.gohtml")
	editTemplate.ExecuteTemplate(w, "base", viewData)
}

func updateCharacter(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	character, ok := r.Context().Value("character").(*characters.Character)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	formErr := character.GetFormData(r)
	util.CheckError(formErr, "")
	if formErr != nil {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}
	character.Save()
	stringID := strconv.FormatUint(uint64(character.ID), 10)
	newURL := fmt.Sprintf("/character/%s", stringID)
	http.Redirect(w, r, newURL, http.StatusSeeOther)
}

func deleteCharacter(w http.ResponseWriter, r *http.Request) {
	theCharacter, ok := r.Context().Value("character").(*characters.Character)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	characterTemplate := newTemplate("templates/delete_confirmation.gohtml")
	characterTemplate.ExecuteTemplate(w, "base", theCharacter)
}

func confirmDeleteCharacter(w http.ResponseWriter, r *http.Request) {
	character, ok := r.Context().Value("character").(*characters.Character)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
	character.Delete()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
