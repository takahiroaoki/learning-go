package thesaurus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BigHuge struct {
	url    string
	apiKey string
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string
	response, err := http.Get(b.url + b.apiKey + "/" + term + "/json")
	if err != nil {
		return syns, fmt.Errorf("bighuge: failed to search synonyms of %q: %v", term, err)
	}
	var data synonyms
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}
	syns = append(syns, data.Noun.Syn...)
	syns = append(syns, data.Verb.Syn...)
	return syns, err
}

func GetInstance() *BigHuge {
	config := GetConfig()
	apiKey := config.API.BigHugeThesaurus.KEY
	url := config.API.BigHugeThesaurus.URL
	return &BigHuge{apiKey: apiKey, url: url}
}
