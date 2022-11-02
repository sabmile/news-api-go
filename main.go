package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sabmile/news-api-go/news"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	s := "welcom"
	w.Write([]byte(s))
}

func searchHandler(newsapi *news.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		params := u.Query()
		searchQuery := params.Get("q")

		fmt.Println(searchQuery)
	}
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

	apiKey := os.Getenv("API_KEY")
	if err != nil {
		log.Fatal("Env: apiKey must be set")
	}

	client := &http.Client{Timeout: 10 * time.Second}
	newsapi := news.NewClient(client, apiKey)

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/search", searchHandler(newsapi))
	http.ListenAndServe(":"+port, mux)
}
