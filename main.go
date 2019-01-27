package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fractalbach/highscores/boardserver"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", boardserver.Handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
