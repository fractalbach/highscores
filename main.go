package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fractalbach/highscores/scoreboard"
)

type Message struct {
	Name  string
	Score int
}

var myboard = scoreboard.NewScoreboard(
	"Balloon Game Highscores",
	"An example of a very simplistic",
)

func myInternalError(err error) string {
	return fmt.Sprintf("Internal Server Error =(\nError:%s", err)
}

// GET board info
// displays information about the board.
func getBoardHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.MarshalIndent(myboard, "", "  ")
	if err != nil {
		http.Error(w, myInternalError(err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(b))
}

// POST new score
// attempts to add a new score to the scoreboard.
func postBoardHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	msg := Message{}
	if err := decoder.Decode(&msg); err != nil {
		http.Error(w, "Bad Request. Invalid JSON Message.", http.StatusBadRequest)
	}
	myboard.Post(scoreboard.NewEntry(msg.Name, msg.Score))
}

// Board handler
// takes GET and POST requests and calls the correct handler.
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getBoardHandler(w, r)
	case "POST":
		postBoardHandler(w, r)
	default:
		http.Error(w, "Method Not Allowed.", http.StatusMethodNotAllowed)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
