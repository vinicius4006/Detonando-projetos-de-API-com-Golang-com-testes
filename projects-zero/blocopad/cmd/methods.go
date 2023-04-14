package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"com.blocopad/blocopad/internal/backend"
	"com.blocopad/blocopad/internal/db"
	"github.com/gorilla/mux"
)

func WriteResponse(status int, body interface{}, w http.ResponseWriter) {
	w.WriteHeader(status)                              // header status code
	w.Header().Set("Content-Type", "application/json") // type content application
	payload, _ := json.Marshal(body)                   // json my content
	w.Write(payload)                                   // body content in json
}

// HTTP Handlers

func ReadNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if data, err := backend.GetKey(id); err == nil {
		WriteResponse(http.StatusOK, data, w)
	} else {
		if err.Error() == "not found" {
			WriteResponse(http.StatusNotFound, "Note not found", w)
		} else {
			WriteResponse(http.StatusInternalServerError, "Error", w)
		}
	}
}

func WriteNote(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var note db.Note
	if err := decoder.Decode(&note); err != nil {
		WriteResponse(http.StatusBadRequest, map[string]string{"error": err.Error()}, w)
		return
	}

	uuidString, err := backend.SaveKey(note.Text, note.OneTime)
	fmt.Println(uuidString, err)
	if err != nil {
		WriteResponse(http.StatusBadRequest, map[string]string{"error": "invalid request"}, w)
	} else {
		WriteResponse(http.StatusOK, map[string]string{"code": uuidString}, w)
	}
}
