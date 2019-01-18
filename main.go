package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fractalbach/highscores/scoreboard"
)

const (
	filename           = "first"
	exampleTitle       = "Balloon Game Highscores"
	exampleDescription = "An example of a very simplistic"
)

var (
	myboard *scoreboard.Scoreboard
	lastmod time.Time
)

func init() {
	if scoreboard.BoardExists(filename) {
		log.Println("board exists, loading file.")
		reload()
		log.Println("last modified:", lastmod)
	} else {
		log.Println("board does not yet exist, creating file.")
		myboard = scoreboard.NewScoreboard(exampleTitle, exampleDescription)
		myboard.SaveAs(filename)
		updateLastMod()
		log.Println("last modified:", lastmod)
	}
}

func hasBeenModified() bool {
	lmod, ok := scoreboard.LastModified(filename)
	if !ok {
		log.Println("hasBeenModified: cannot retrieve last-modified time.")
	}
	return lmod == lastmod
}

func updateLastMod() {
	lmod, ok := scoreboard.LastModified(filename)
	if !ok {
		log.Println("updateLastMod: cannot retrieve last-modified time.")
	}
	lastmod = lmod
}

func reload() {
	myboard = scoreboard.Load(filename)
	updateLastMod()
}

type Message struct {
	Name  string
	Score float64
}

func myInternalError(err error) string {
	return fmt.Sprintf("Internal Server Error =(\nError:%s", err)
}

// GET board info
// displays all information about the scoreboard.
func getBoardHandler(w http.ResponseWriter, r *http.Request) {
	if hasBeenModified() {
		reload()
	}
	b, err := json.Marshal(myboard)
	if err != nil {
		http.Error(w, myInternalError(err), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

// POST new score
// attempts to add a new score to the scoreboard.
func postBoardHandler(w http.ResponseWriter, r *http.Request) {
	if hasBeenModified() {
		reload()
	}
	decoder := json.NewDecoder(r.Body)
	msg := Message{}
	if err := decoder.Decode(&msg); err != nil {
		http.Error(w, "Bad Request. Invalid JSON Message: "+err.Error(), http.StatusBadRequest)
		return
	}
	myboard.Post(scoreboard.NewEntry(msg.Name, int(msg.Score)))
	myboard.SaveAs(filename)
	updateLastMod()
}

// Board handler
// takes GET and POST requests and calls the correct handler.
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	switch r.Method {
	case "GET":
		getBoardHandler(w, r)
	case "POST":
		postBoardHandler(w, r)
	case "OPTIONS":
		return
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
