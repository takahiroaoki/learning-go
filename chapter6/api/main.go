package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/stretchr/graceful"
	"gopkg.in/mgo.v2"
)

func main() {
	var (
		addr  = flag.String("addr", ":8080", "The address of the endpoint")
		mongo = flag.String("mongo", "localhost", "The address of the mongoDB")
	)
	flag.Parse()
	log.Println("Get access to mongoDB", *mongo)
	var db *mgo.Session
	/*
		mongoInfo := &mgo.DialInfo{
			Addrs:    []string{*mongo},
			Username: "root",
			Password: "password",
		}
		db, err := mgo.DialWithInfo(mongoInfo)
		if err != nil {
			log.Fatalln("Failed to access to mongoDB: ", err)
		}
		defer db.Close()
	*/
	mux := http.NewServeMux()
	mux.HandleFunc("/polls/", withCORS(
		withVars(
			withData(
				db, withAPIKey(handlePolls),
			),
		),
	))
	log.Println("Starting web server...")
	graceful.Run(*addr, 1*time.Second, mux)
	log.Println("Stopping web server...")
}

func withAPIKey(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isValidAPIKey(r.URL.Query().Get("key")) {
			respondErr(w, r, http.StatusUnauthorized, "Invalid API key")
			return
		}
		fn(w, r)
	}
}

func isValidAPIKey(key string) bool {
	return key == "abc123"
}

func withData(d *mgo.Session, fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
			thisDb := d.Copy()
			defer thisDb.Close()
			SetVar(r, "db", thisDb.DB("ballots"))
		*/
		fn(w, r)
	}
}

func withVars(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		OpenVars(r)
		defer CloseVars(r)
		fn(w, r)
	}
}

func withCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Location")
		fn(w, r)
	}
}
