package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-mongodb/database"
	"go-mongodb/handler"
	"go-mongodb/repository"
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
	var (
		port    = os.Getenv("SERVER_PORT")
		DB_HOST = os.Getenv("DB_HOST")
		DB_PORT = os.Getenv("DB_PORT")
		DB_TYPE = os.Getenv("DB")
		DB_USER = os.Getenv("USER_DB")
		DB_PASS = os.Getenv("PASS_DB")
	)

	var certificate string
	if DB_USER != "" {
		certificate = fmt.Sprintf("%v:%v@", DB_USER, DB_PASS)
	} else {
		certificate = ""
	}

	databaseURL := fmt.Sprintf("%v://%v%v:%v", DB_TYPE, certificate, DB_HOST, DB_PORT)

	db, err := database.NewDatabase(databaseURL)
	if err != nil {
		log.Fatalf("Connect database error: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Can't connect to database: %v", err)
	} else {
		fmt.Println("Connect to database successfully")
	}

	authorRepository := repository.NewAuthorRepository(db)
	documentRepository := repository.NewDocumentRepository(db)

	authorHandler := handler.NewAuthorHandler(authorRepository)
	documentHandler := handler.NewDocumentHandler(documentRepository)

	r := router.NewRouter(authorHandler, documentHandler)

	serv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("127.0.0.1:%v", port),
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
	}

	fmt.Printf("Server are running in port: %v\n", port)
	log.Fatal(serv.ListenAndServe())
}
