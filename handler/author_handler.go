package handler

import (
	"encoding/json"
	"go-mongodb/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
)

type AuthorHandler struct {
	model.AuthorRepository
}

func NewAuthorHandler(r model.AuthorRepository) *AuthorHandler {
	return &AuthorHandler{
		AuthorRepository: r,
	}
}

func (h *AuthorHandler) ListAllAuthor(w http.ResponseWriter, r *http.Request) {
	skipStr, limitStr := r.URL.Query().Get("skip"), r.URL.Query().Get("limit")
	skip, err := strconv.Atoi(skipStr)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]error{"error": err})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]error{"error": err})
		return
	}

	authors, err := h.AuthorRepository.ListAllAuthor(int64(skip), int64(limit))
	if authors == nil {
		ResponseWithJSON(w, http.StatusNotFound, map[string]string{"error": "no author available"})
		return
	}

	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "server error"})
		return
	}

	ResponseWithJSON(w, http.StatusOK, authors)
}

func (h *AuthorHandler) GetAuthorDetail(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]string{"error": "id not true"})
		return

	}

	author, err := h.AuthorRepository.GetAuthorDetail(objectID)
	if err != nil {
		ResponseWithJSON(w, http.StatusNotFound, map[string]string{"error": "author not found or be deleted"})
		return
	}

	ResponseWithJSON(w, http.StatusOK, author)
}

func (h *AuthorHandler) InsertAuthor(w http.ResponseWriter, r *http.Request) {
	var authorWrite model.AuthorWrite
	err := json.NewDecoder(r.Body).Decode(&authorWrite)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]string{"error": "have some field not valid"})
		return
	}
	authorWrite.IsDeleted = false
	objectID, err := h.AuthorRepository.InsertAuthor(&authorWrite)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]string{"error": "fail to insert author"})
		return
	}

	ResponseWithJSON(w, http.StatusCreated, map[string]interface{}{"_id": objectID})
}

func (h *AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]error{"error": err})
		return
	}

	_, err = h.AuthorRepository.GetAuthorDetail(objectID)
	if err != nil {
		ResponseWithJSON(w, http.StatusNotFound, map[string]string{"error": "author not found"})
		return
	}

	var author model.Author
	err = json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]interface{}{"error": "fail to decode author"})
		return
	}

	author.ID = objectID
	err = h.AuthorRepository.UpdateAuthor(&author)
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError, map[string]error{"error": err})
		return
	}

	ResponseWithJSON(w, http.StatusAccepted, author)
}

func (h *AuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]string{"error": "id not true"})
	}

	err = h.AuthorRepository.DeleteAuthor(objectID)
	if err != nil {
		ResponseWithJSON(w, http.StatusNotFound, map[string]interface{}{"error": "not found author"})
		return
	}

	ResponseWithJSON(w, http.StatusOK, map[string]string{"message": "delete author successfully"})
}

func (h *AuthorHandler) WriteMessage(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Hello world"))
}
