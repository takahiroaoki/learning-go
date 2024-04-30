package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

type dbData struct {
	Polls []*poll `json:"polls"`
}

func handlePollsGet(w http.ResponseWriter, r *http.Request) {
	/*
		db := GetVar(r, "db").(*mgo.Database)

			c := db.C("polls")
			var q *mgo.Query
	*/
	jsonFile, err := os.Open("db.json")
	if err != nil {
		fmt.Println("Failed to open json file", err)
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Failed to read json file", err)
		return
	}
	var dbData dbData
	json.Unmarshal(jsonData, &dbData)
	var polls []*poll = dbData.Polls

	p := NewPath(r.URL.Path)
	var result []*poll
	if p.HasID() {
		for _, poll := range polls {
			if poll.ID == p.ID {
				result = append(result, poll)
				break
			}
		}
		if len(result) == 0 {
			respondErr(w, r, http.StatusNotFound, fmt.Errorf("id %s is not found", p.ID))
		}
	} else {
		result = append(result, polls...)
	}
	respond(w, r, http.StatusOK, &result)
}

func handlePollsPost(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("Not implemented yet"))
}

func handlePollsDelete(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("Not implemented yet"))
}
