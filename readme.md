# Highscores Server

[![GoDoc](https://godoc.org/github.com/fractalbach/highscores/scoreboard?status.svg)](https://godoc.org/github.com/fractalbach/highscores/scoreboard)
[![Build Status](https://travis-ci.org/fractalbach/highscores.svg?branch=master)](https://travis-ci.org/fractalbach/highscores)

Simple and lightweight Highscores server designed for the
[Balloon Grab](https://github.com/fractalbach/float-up)
Game. Uses POST requests to add new highscores, and
GET requests to view them.

Method |    Path    | Description
-------|------------|--------------
GET | `/` | View scores
POST | `/` | submit a new score entry.







## POST a Score to a board

-   POST an entry in JSON format.

-   If the score is too low, or the board doesn't exist,
    then the request will be ignored.

### POST body : Data Fields

Field Name | Type | Description
-----------|------|--------------
name | string | the player's name
score | integer | the player's score.

### POST body : JSON example

```json
{
    "name": "EXAMPLE",
    "score": 1337
}
```






## GET data from a scoreboard

-   Calling the GET method on a scoreboard will return a JSON page with some
    data describing the scoreboard, and the data of each entry.

-   The **Entries** are almost identical to the submitted entry data,
    but will additionally have a **Time** field that marks when the
    entry was received and added to the scoreboard.

### GET response : Data Fields

Field Name | Type | Description
-----------|------|--------------
Title | string | title of the scoreboard
Description | string | description of the scoreboard.
MaxEntries | int | maximum number of scores that can be saved.
Entries | array | array of entries, highest score is always first.

### GET response : JSON example

```json
{
  "Title": "example",
  "Description": "example",
  "MaxEntries": 20,
  "Entries": [
    {
      "Name": "player1",
      "Score": 200,
      "Time": "2019-01-14T06:19:56.252887329Z"
    },
    {
      "Name": "player2",
      "Score": 100,
      "Time": "2019-01-14T06:19:56.252157208Z"
    }
  ]
}
```





## Future API Paths

___Note___ :
Currently, only 1 board is supported, but once more have been added,
the paths will look like this:

 Method |    Path    | Description
--------|------------|--------------
GET | `/boards` | View a summary of all boards (**Not implemented yet**)
GET | `/boards/{name}` | view the specific board named `{name}` (**Not implemented yet**)
POST | `/boards/{name}` | submit a new score entry to specific board. (**Not implemented yet**)
