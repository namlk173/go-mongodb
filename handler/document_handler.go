package handler

import (
	"encoding/json"
	"go-mongodb/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type DocumentHandler struct {
	model.DocumentRepository
}

func NewDocumentHandler(r model.DocumentRepository) *DocumentHandler {
	return &DocumentHandler{
		DocumentRepository: r,
	}
}

func (h *DocumentHandler) ListAllDocument(w http.ResponseWriter, _ *http.Request) {
	documents, err := h.DocumentRepository.ListAllDocument()
	if err != nil {
		ResponseWithJSON(w, http.StatusNotFound, map[string]string{"error": "server error"})
		return
	}

	if documents == nil {
		ResponseWithJSON(w, http.StatusNotFound, map[string]string{"error": "no document available"})
		return
	}

	ResponseWithJSON(w, http.StatusOK, documents)
}

func (h *DocumentHandler) GetDocumentDetail(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]string{"error": "id not true"})
		return
	}

	document, err := h.DocumentRepository.GetDocumentDetail(objectID)
	if err != nil {
		ResponseWithJSON(w, http.StatusNotFound, map[string]error{"error": err})
		return
	}

	ResponseWithJSON(w, http.StatusOK, document)
}

func (h *DocumentHandler) InsertDocument(w http.ResponseWriter, r *http.Request) {
	var documentWrite model.DocumentWrite
	err := json.NewDecoder(r.Body).Decode(&documentWrite)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]string{"error": "data not validate"})
		return
	}

	document, err := h.DocumentRepository.InsertDocument(&documentWrite)
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "server error"})
		return
	}

	ResponseWithJSON(w, http.StatusCreated, document)
}

func (h *DocumentHandler) UpdateDocument(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]string{"error": "id not true"})
		return
	}

	_, err = h.DocumentRepository.GetDocumentDetail(objectID)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]string{"error": "document not found."})
		return
	}

	var document model.Document
	err = json.NewDecoder(r.Body).Decode(&document)
	if err != nil {
		ResponseWithJSON(w, http.StatusBadRequest, map[string]string{"error": "data not validate"})
		return
	}

	document.ID = objectID
	err = h.DocumentRepository.UpdateDocument(&document)
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError, map[string]error{"error": err})
		return
	}

	ResponseWithJSON(w, http.StatusAccepted, document)
}

func (h *DocumentHandler) DeleteDocument(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ResponseWithJSON(w, http.StatusNotFound, map[string]string{"error": "id not true"})
		return
	}

	err = h.DocumentRepository.DeleteDocument(objectID)
	if err != nil {
		ResponseWithJSON(w, http.StatusInternalServerError, map[string]error{"error": err})
		return
	}

	ResponseWithJSON(w, http.StatusOK, map[string]string{"message": "delete document successful"})
}
