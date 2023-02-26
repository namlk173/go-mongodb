package router

import (
	"github.com/gorilla/mux"
	"go-mongodb/handler"
)

func NewRouter(authorHandler *handler.AuthorHandler, documentHandler *handler.DocumentHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", authorHandler.WriteMessage).Methods("GET")
	r.HandleFunc("/authors", authorHandler.ListAllAuthor).Methods("GET")
	r.HandleFunc("/author", authorHandler.GetAuthorDetail).Methods("GET")
	r.HandleFunc("/author", authorHandler.InsertAuthor).Methods("POST")
	r.HandleFunc("/author", authorHandler.UpdateAuthor).Methods("PUT", "PATCH")
	r.HandleFunc("/author", authorHandler.DeleteAuthor).Methods("DELETE")

	r.HandleFunc("/document", documentHandler.WriteMessage).Methods("GET")
	return r
}
