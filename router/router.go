package router

import (
	"github.com/gorilla/mux"
	"go-mongodb/controller"
)

func NewRouter(handler *controller.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.WriteMessage).Methods("GET")
	r.HandleFunc("/authors", handler.ListAllAuthor).Methods("GET")
	r.HandleFunc("/author", handler.GetAuthorDetail).Methods("GET")
	r.HandleFunc("/author", handler.InsertAuthor).Methods("POST")
	r.HandleFunc("/author", handler.UpdateAuthor).Methods("PUT", "PATCH")
	r.HandleFunc("/author", handler.DeleteAuthor).Methods("DELETE")
	return r
}
