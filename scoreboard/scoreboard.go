/*
Package scoreboard holds a collection of scores that can be stored or loaded.

The scoreboard data are saved as JSON files.
*/
package scoreboard

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
	"time"
)

// ScoreType specifies the golang type that represents a score.
// It's specified here in case it needs to be changed later.
type ScoreType int

/*
Compare returns an integer that is...
	negative if a < b
	zero     if a = b
	positive if a > b
This allows ScoreType to be "Comparable" to itself.
In the future, there might be an interface for "Comparable",
and ScoreType would represent it's concrete type.
*/
func (a ScoreType) Compare(b ScoreType) int {
	return int(a - b)
}

// Scoreboard keeps track of multiple score entries for a single game type.
type Scoreboard struct {
	Title       string
	Description string
	MaxEntries  int
	Entries     []*Entry
}

// Entry holds a single entry on the scoreboard.
// When a player beats the other highscores, a new entry is created and
// saved into the scoreboard.
type Entry struct {
	Name  string
	Score ScoreType
	Time  time.Time
}

// NewEntry creates a standard integer highscore entry using the current time.
// This is most likely what will be used.
func NewEntry(name string, score int) *Entry {
	return &Entry{
		Name:  name,
		Score: ScoreType(score),
		Time:  time.Now(),
	}
}

type Entries []*Entry

func (e Entries) Len() int      { return len(e) }
func (e Entries) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

type ByScore struct{ Entries }

// Less is actually More, because we want the highest numbers to be first.
func (s ByScore) Less(i, j int) bool {
	a := s.Entries[i].Score
	b := s.Entries[j].Score
	return a.Compare(b) > 0
}

/*
Post :
compares the given entry's score to others in the scoreboard.
Returns false if the score is too low and the board is already full.
Returns true if the entry is successfull added.
*/
func (s *Scoreboard) Post(entry *Entry) bool {
	if len(s.Entries) < s.MaxEntries {
		s.Entries = append(s.Entries, entry)
		sort.Stable(ByScore{s.Entries})
		return true
	}
	incomingScore := entry.Score
	lowestScore := s.lowestEntry().Score
	if incomingScore.Compare(lowestScore) > 0 {
		s.insert(entry)
		return true
	}
	return false
}

// SaveAs : saves the scoreboard into a json file.
func (s *Scoreboard) SaveAs(filename string) {
	filename = filename + ".json"
	b, err := json.Marshal(s)
	if err != nil {
		log.Println("SaveAs :: unable to convert scoreboard to json ::", err)
		return
	}
	err = ioutil.WriteFile(filename, b, 0600)
	if err != nil {
		log.Println("SaveAs :: unable to save scoreboard to file ::", err)
		return
	}
}

// Load : retrieves a scoreboard from a file.
func Load(filename string) *Scoreboard {
	filename = filename + ".json"
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

func (s *Scoreboard) lowestEntry() *Entry {
	index := len(s.Entries) - 1
	return s.Entries[index]
}

// insert entry by removing the lowest score and adding a new one.
func (s *Scoreboard) insert(e *Entry) {
	index := len(s.Entries) - 1
	s.Entries[index] = e
	sort.Stable(ByScore{s.Entries})
}
