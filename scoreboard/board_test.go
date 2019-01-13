package scoreboard

import (
	"fmt"
	"io/ioutil"
	"testing"
)

const fancystring = `
_________________________________________________
%s
=================================================
%v
`

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

func TestSaveLoad(t *testing.T) {
	board.SaveAs("testing_example")
	data, err := ioutil.ReadFile("testing_example.json")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))

}
