package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"thesaurus"
)

func main() {
	thesaurus := thesaurus.GetInstance()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalf("Failed to search the synonyms of %q: %v", word, err)
		}
		if len(syns) == 0 {
			log.Fatalf("%q has no synonyms", word)
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
