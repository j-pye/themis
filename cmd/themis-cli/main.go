package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func main() {
	answers := make(map[string]interface{})
	uiLayout := make(map[string]interface{})
	data, err := ioutil.ReadFile("cmd/themis-cli/layout.yml")
	if err != nil {
		log.Fatalf("Error loading layout %s", err)
	} else {
		err := yaml.Unmarshal(data, uiLayout)
		if err != nil {
			log.Fatalf("Error parsing layout %s", err)
		}
	}
	prompts := uiLayout["prompts"].([]interface{})
	for _, q := range prompts {
		processQuestion(q)
	}
	print(answers)
}

func processQuestion(q interface{}) {
	fmt.Print(q)
}
