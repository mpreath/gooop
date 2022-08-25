package workflow

import (
	"encoding/json"
	"log"
)

type ValidateProcess struct {
	BaseProcess
}

func (p *ValidateProcess) Run(person *Person) {
	log.Println("############ Validating #############")
	personJson, _ := json.MarshalIndent(person, "", "  ")
	log.Println(string(personJson))
	if person.Address == "" {
		log.Println("Address not specified but needs to be.")
	} else {
		log.Println("Address is valid.")
		person.Handled = true
	}
	log.Println("#####################################")
}
