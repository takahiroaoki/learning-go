package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type poll struct {
	ID      string         `json:"id"`
	Title   string         `json:"title"`
	Options []string       `json:"options"`
	Results map[string]int `json:"results,omitempty"`
}

func handlePolls(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlePollsGet(w, r)
		return
	case "POST":
		handlePollsPost(w, r)
		return
	case "DELETE":
		handlePollsDelete(w, r)
		return
	}
	respondHTTPErr(w, r, http.StatusNotFound)
}

var dbMockData = []*poll{
	{
		ID:      "1",
		Title:   "test1",
		Options: []string{"one", "two", "three"},
		Results: map[string]int{
			"one":   100,
			"two":   200,
			"three": 300,
		},
	},
	{
		ID:      "2",
		Title:   "test2",
		Options: []string{"one", "two", "three"},
		Results: map[string]int{
			"one":   500,
			"two":   600,
			"three": 700,
		},
	},
}

func handlePollsGet(w http.ResponseWriter, r *http.Request) {
	p := NewPath(r.URL.Path)
	var result []*poll
	if p.HasID() {
		for _, poll := range dbMockData {
			if poll.ID == p.ID {
				result = append(result, poll)
				break
			}
		}
		if len(result) == 0 {
			respondErr(w, r, http.StatusNotFound, fmt.Errorf("id %s is not found", p.ID))
		}
	} else {
		result = append(result, dbMockData...)
	}
	respond(w, r, http.StatusOK, &result)
}

func handlePollsPost(w http.ResponseWriter, r *http.Request) {
	var p poll
	if err := decodeBody(r, &p); err != nil {
		respondErr(w, r, http.StatusBadRequest, "Failed to decode item", err)
	}
	p.ID = fmt.Sprint(time.Now().UTC().UnixNano())

	dbMockData = append(dbMockData, &p)

	w.Header().Set("Location", "polls/"+p.ID)
	respond(w, r, http.StatusCreated, nil)
}

func handlePollsDelete(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("Not implemented yet"))
}
