/*
Package scoreboard holds a collection of scores that can be stored or loaded.
*/
package scoreboard

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

// Scoreboard keeps track of multiple score entries for a single game type.
type Scoreboard struct {
	Title   string
	Entries []Entry
}

// Entry holds a single entry on the scoreboard.
// When a player beats the other highscores, a new entry is created and
// saved into the scoreboard.
type Entry struct {
	Name  string
	Score int
	Time  time.Time
}

// Store : saves the scoreboard into a file.
func (s *Scoreboard) Store() {
	if s.Title == "" {
		log.Println("Store :: can't store scoreboard with an empty title.")
		return
	}
	filename := s.Title + ".json"
	b, err := json.Marshal(s)
	if err != nil {
		log.Println("Store :: unable to convert scoreboard to json ::", err)
		return
	}
	err = ioutil.WriteFile(filename, b, 0600)
	if err != nil {
		log.Println("Store :: unable to store scoreboard to file ::", err)
		return
	}
}

// Load : retrieves a scoreboard from a file.
func Load(title string) *Scoreboard {
	filename := title + ".json"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Load :: unable to read file ::", err)
	}
	scoreboard := &Scoreboard{}
	err = json.Unmarshal(data, scoreboard)
	if err != nil {
		log.Println("Load :: unable to convert file data to json :: ", err)
	}
	return scoreboard
}
