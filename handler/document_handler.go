package handler

import (
	"go-mongodb/model"
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

func (h *DocumentHandler) WriteMessage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world."))
}
