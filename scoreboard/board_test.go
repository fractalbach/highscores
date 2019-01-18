package scoreboard

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const fancystring = `
_________________________________________________
%s
=================================================
%v
`

const (
	examplePrefix   = "test_data_output"
	exampleFilename = examplePrefix + ".json"
)

var board = Scoreboard{
	Title:       "Highscores",
	Description: "Example of a highscores board.",
	MaxEntries:  5,
}

func display() {
	fmt.Println(board)
	fmt.Printf(fancystring, board.Title, board.Entries)
}

func show(t *testing.T) {
	s := ""
	for _, e := range board.Entries {
		s += fmt.Sprintf("(%v) ", e.Score)
	}
	t.Log(s)
}

func scoresString(board *Scoreboard) string {
	return fmt.Sprintf("%#v", board)
}

func TestMain(t *testing.T) {
	examples := []struct {
		name  string
		score int
	}{
		{"noob1", 1},
		{"W", 40},
		{"Y", 20},
		{"noob2", 2},
		{"V", 50},
		{"Z", 10},
		{"X", 30},
		{"waddup", 111},
		{"asdfasdf", 122},
		{"1337", 1337},
		{"proaf", 188},
		{"noob5", 5},
		{"omgz", 999},
		{"A", 4242},
		{"B", 1001},
		{"C", 1000},
		{"noob3", 3},
		{"noob4", 4},
	}
	for _, v := range examples {
		board.Post(NewEntry(v.name, v.score))
		show(t)
	}
}

func TestSave(t *testing.T) {
	board.SaveAs(examplePrefix)
	data, err := ioutil.ReadFile(exampleFilename)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
	f, err := os.Stat(exampleFilename)
	if err != nil {
		t.Error(err)
	}
	t.Log(f.Name(), f.Size(), f.Mode(), f.ModTime(), f.IsDir())
}

func TestFileExists(t *testing.T) {
	exists := BoardExists(examplePrefix)
	if !exists {
		t.Error("BoardExists returned false, but the file should be there.")
		t.Fail()
	}
}

func TestLoad(t *testing.T) {
	loadedBoard := Load(examplePrefix)
	s := scoresString(loadedBoard)
	s2 := scoresString(&board)
	t.Log(s)
	t.Log(s2)
	if s != s2 {
		t.Error("scores are not the same after saving & reloading them.")
	}
}

func TestOverwrite(t *testing.T) {
	b := NewScoreboard("overwrite board.", "the new one that overwrote the last one.")
	b.Post(NewEntry("the only player", 100))
	b.SaveAs(examplePrefix)
	bLoaded := Load(examplePrefix)
	s1 := fmt.Sprintf("%#v", b)
	s2 := fmt.Sprintf("%#v", bLoaded)
	t.Log(s1)
	t.Log(s2)
	if s1 != s2 {
		t.Error("a new board doesn't save correctly after overwriting an existing file.")
	}
}

func TestBadStructure(t *testing.T) {
	b := Scoreboard{
		Title:       "asdf",
		Description: "qwertyuiop",
	}
	b.Post(NewEntry("asdfasdf", 100))
	b.Post(NewEntry("asdfasdf33", 300))
	b.Post(NewEntry("asdfasdf33", 200))
}
