/*
Package scoreboard holds a collection of scores that can be stored or loaded.

The scoreboard data are saved as JSON files.
*/
package scoreboard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
	"time"
)

const (
	DefaultMaxEntries = 20
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
	mutex       sync.Mutex
}

// NewScoreboard creates a scoreboard object with default settings.
func NewScoreboard(title, description string) *Scoreboard {
	return &Scoreboard{
		Title:       title,
		Description: description,
		MaxEntries:  DefaultMaxEntries,
	}
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
		Time:  time.Now().UTC(),
	}
}

type entries []*Entry

func (e entries) Len() int      { return len(e) }
func (e entries) Swap(i, j int) { e[i], e[j] = e[j], e[i] }

type byScore struct{ entries }

// Less is actually More, because we want the highest numbers to be first.
func (s byScore) Less(i, j int) bool {
	a := s.entries[i].Score
	b := s.entries[j].Score
	return a.Compare(b) > 0
}

/*
Post :
compares the given entry's score to others in the scoreboard.
Returns false if the score is too low and the board is already full.
Returns true if the entry is successfull added.
*/
func (s *Scoreboard) Post(entry *Entry) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.MaxEntries == 0 {
		return false
	}
	if len(s.Entries) < s.MaxEntries {
		s.Entries = append(s.Entries, entry)
		sort.Stable(byScore{s.Entries})
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
	s.mutex.Lock()
	defer s.mutex.Unlock()
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

// BoardExists looks to see if the board has already been saved in this current
// directory.  Returns true if it can find the board, returns false otherwise.
func BoardExists(filename string) bool {
	if _, err := os.Stat(filename + ".json"); !os.IsNotExist(err) {
		return true
	}
	return false
}

// LastModifed returns the last-modified time of the file of the given
// scoreboard.  The second return value is an "ok" check.  If there are any
// problems accessing the file (or if it doesn't exist), then that boolean
// will return false.
func LastModified(filename string) (time.Time, bool) {
	info, err := os.Stat(filename + ".json")
	if err != nil {
		return time.Time{}, false
	}
	return info.ModTime(), true
}

func (s *Scoreboard) lowestEntry() *Entry {
	index := len(s.Entries) - 1
	return s.Entries[index]
}

// insert entry by removing the lowest score and adding a new one.
func (s *Scoreboard) insert(e *Entry) {
	index := len(s.Entries) - 1
	s.Entries[index] = e
	sort.Stable(byScore{s.Entries})
}

func (s Scoreboard) GoString() string {
	out := ""
	for _, e := range s.Entries {
		out += fmt.Sprintf("(%v) ", e.Score)
	}
	return out
}
