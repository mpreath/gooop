package workflow

import "log"

type ValidateProcess struct {
	BaseProcess
}

func (p *ValidateProcess) Execute(person *Person) {
	log.Println("############ Validating #############")
	if person.Address == "" {
		log.Println("Address not specified but needs to be.")
	} else {
		log.Println("Person is valid.")
		person.Zip = "55555"
	}
	log.Println("#####################################")
}
