package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-mongodb/controller"
	"go-mongodb/database"
	"go-mongodb/router"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}

	// Load enviroment variable
	port := os.Getenv("SERVER_PORT")
	databaseURL := os.Getenv("DATABASE_URL")

	database, err := database.NewDatabase(databaseURL)
	if err != nil {
		log.Fatalf("Connect database error: %v", err)
	}
	defer database.Close()

	if err := database.Ping(); err != nil {
		log.Fatalf("Can't connect to database: %v", err)
	} else {
		fmt.Println("Connect to database successfully")
	}

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
