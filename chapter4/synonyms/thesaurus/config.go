package thesaurus

import (
	"log"
	"os"
	"sync"

	yml "gopkg.in/yaml.v2"
)

type config struct {
	API api `yaml:"api"`
}

type api struct {
	BigHugeThesaurus bigHugeThesaurus `yaml:"big-huge-thesaurus"`
}

type bigHugeThesaurus struct {
	URL string `yaml:"url"`
	KEY string `yaml:"key"`
}

var instanceConfig *config
var onceYaml sync.Once

func GetConfig() *config {
	onceYaml.Do(func() {
		initializeYaml()
	})
	return instanceConfig
}

func initializeYaml() {
	buf, err := os.ReadFile("config/config.yml")
	if err != nil {
		log.Fatalln(err)
	}
	instanceConfig = &config{}
	err = yml.Unmarshal(buf, instanceConfig)
	if err != nil {
		log.Fatalln(err)
	}
}
