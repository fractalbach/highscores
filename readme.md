Highscores Server
==============================

[![GoDoc](https://godoc.org/github.com/fractalbach/highscores/scoreboard?status.svg)](https://godoc.org/github.com/fractalbach/highscores/scoreboard)
[![Build Status](https://travis-ci.org/fractalbach/highscores.svg?branch=master)](https://travis-ci.org/fractalbach/highscores)

Simple and lightweight Highscores server designed for the
[Balloon Grab](https://github.com/fractalbach/float-up)
Game. Uses POST requests to add new highscores, and
GET requests to view them.

## API Documentation

___Note___ : These haven't been implemented yet.
It's a general description of what this project is about
Also acts as road map for it's Development.


 Method |    Path    | Description
--------|------------|--------------
GET | `/boards` | View a summary of all boards
GET | `/boards/{name}` | view the specific board named `{name}`
POST | `/boards/{name}` | submit a new score entry to specific board.

## Posting a Score to a board

-   POST an entry in JSON format.

-   If the score is too low, or the board doesn't exist,
    then the request will be ignored.

-   ___Note___ : more data fields will probably be added in the future.
    Probably an identifier or something that prevents random requests
    from being sent.

### Data Fields

Field Name | Type | Description
-----------|------|--------------
name | string | the player's name
score | integer | the player's score.

### JSON example
```json
{
    "name": "EXAMPLE",
    "score": 1337,
}
```
