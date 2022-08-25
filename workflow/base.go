package workflow

import (
	"encoding/json"
	"log"
)

type BaseProcess struct {
	next Process
}

func (p *BaseProcess) SetNext(next Process) {
	p.next = next
}

func (p *BaseProcess) GetNext() Process {
	return p.next
}

func (p *BaseProcess) Run(person *Person) {
	log.Println("########## Base processing ##########")
	personJson, _ := json.MarshalIndent(person, "", "  ")
	log.Println(string(personJson))
	log.Println("#####################################")
}
