package main

import (
	"log"
	"net/http"
	"os"
	"user/database"
	"user/handler"
	"user/server"
)

var (
	port         = "8082"
	databaseFile = "./sqliteDB/userprofile.db"
)

func init() {
	if env := os.Getenv("PORT"); env != "" {
		port = env
	}
	if env := os.Getenv("DATABASE_FILE_NAME"); env != "" {
		databaseFile = env
	}
}

func main() {
	db, err := database.InitDB(databaseFile)
	if err != nil {
		panic(err)
	}
	server := server.Server{
		Router: http.NewServeMux(),
	}
	h := handler.Handler{
		DB: db,
	}
	server.InitRoute(&h)
	log.Println("started user-profile microservice...")
	log.Fatal(http.ListenAndServe(`:`+port, server.Router))

}
