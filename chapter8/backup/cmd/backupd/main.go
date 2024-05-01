package main

import (
	"backup"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/matryer/filedb"
)

type path struct {
	Path string
	Hash string
}

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			log.Fatalln(fatalErr)
		}
	}()

	var (
		interval = flag.Int("interval", 10, "The interval of check")
		archive  = flag.String("archive", "archive", "The destination of archive")
		dbpath   = flag.String("db", "db", "The path to filedb")
	)
	flag.Parse()

	m := &backup.Monitor{
		Destination: *archive,
		Archiver:    backup.ZIP,
		Paths:       make(map[string]string),
	}
	db, err := filedb.Dial(*dbpath)
	if err != nil {
		fatalErr = err
		return
	}
	defer db.Close()
	col, err := db.C("paths")
	if err != nil {
		fatalErr = err
		return
	}

	var path path
	col.ForEach(func(_ int, data []byte) bool {
		if err := json.Unmarshal(data, &path); err != nil {
			fatalErr = err
			return true
		}
		m.Paths[path.Path] = path.Hash
		return false
	})
	if fatalErr != nil {
		return
	}

	if len(m.Paths) < 1 {
		fatalErr = errors.New("no paths. add some via backup tool")
		return
	}

	check(m, col)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
Loop:
	for {
		select {
		case <-time.After(time.Duration(*interval) * time.Second):
			check(m, col)
		case <-signalChan:
			fmt.Println()
			log.Println("Stopping...")
			break Loop
		}
	}
}

func check(m *backup.Monitor, col *filedb.C) {
	log.Println("Checking...")
	counter, err := m.Now()
	if err != nil {
		log.Panicln("Failed to back up: ", err)
	}
	if counter > 0 {
		log.Printf("Archived %d directories\n", counter)

		var path path
		col.SelectEach(func(_ int, data []byte) (bool, []byte, bool) {
			if err := json.Unmarshal(data, &path); err != nil {
				log.Println("Failed to read JSON. Ahead to the next item: ", err)
				return true, data, false
			}
			path.Hash, _ = m.Paths[path.Path]
			newdata, err := json.Marshal(&path)
			if err != nil {
				log.Println("Failed to output JSON. Ahead to the next item: ", err)
				return true, data, false
			}
			return true, newdata, false
		})
	} else {
		log.Println("No modifications found")
	}
}
