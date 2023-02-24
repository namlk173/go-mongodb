package main

import (
	"fmt"
	"go-mongodb/controller"
	"go-mongodb/database"
	"go-mongodb/router"
	"log"
	"net/http"
	"time"
)

const port = 8080

func main() {

	database, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Connect database error: %v", err)
	}
	defer database.Close()

	repository := controller.NewRepository(database)
	handler := controller.NewHandler(repository)
	r := router.NewRouter(handler)

	serv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("127.0.0.1:%v", port),
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
	}

	fmt.Printf("Server are running in port: %v\n", port)
	log.Fatal(serv.ListenAndServe())
}
