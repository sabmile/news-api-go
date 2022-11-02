package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	s := "welcom"
	w.Write([]byte(s))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
