package controller

import (
	"encoding/json"
	"go-mongodb/model"
	"net/http"
)

type Handler struct {
	model.Repository
}

func NewHandler(r model.Repository) *Handler {
	return &Handler{
		Repository: r,
	}
}

func (h *Handler) ListAllAuthor(w http.ResponseWriter, r *http.Request) {
	authors, err := h.Repository.ListAllAuthor()
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "no author available"})
		return
	}
	ResponseWithJSON(w, http.StatusOK, authors)
}

func (h *Handler) GetAuthorDetail(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	author, err := h.Repository.GetAuthorDetail(id)
	if err != nil {
		ResponseWithJSON(w, http.StatusNotFound, map[string]string{"error": "not found author"})
		return
	}
	ResponseWithJSON(w, http.StatusOK, author)
}

func (h *Handler) InsertAuthor(w http.ResponseWriter, r *http.Request) {
	var authorWrite model.AuthorWrite
	err := json.NewDecoder(r.Body).Decode(&authorWrite)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]string{"error": "have some field not valid"})
		return
	}
	author, err := h.Repository.InsertAuthor(&authorWrite)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]string{"error": "fail to insert author"})
		return
	}
	ResponseWithJSON(w, http.StatusCreated, &author)
}

func (h *Handler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var authorWrite model.AuthorWrite
	err := json.NewDecoder(r.Body).Decode(&authorWrite)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]interface{}{"error": "fail to decode author"})
		return
	}
	id := r.URL.Query().Get("id")
	author, err := h.Repository.UpdateAuthor(id, authorWrite)
	if err != nil {
		ResponseWithJSON(w, http.StatusNotFound, map[string]interface{}{"error": "not found author"})
		return
	}
	ResponseWithJSON(w, http.StatusOK, author)
}

func (h *Handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := h.Repository.Delete(id)
	if err != nil {
		ResponseWithJSON(w, http.StatusNotFound, map[string]interface{}{"error": "not found author"})
		return
	}
	ResponseWithJSON(w, http.StatusOK, map[string]string{"message": "delete author successfully"})
}

func ResponseWithJSON(w http.ResponseWriter, code int, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (h *Handler) WriteMessage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
