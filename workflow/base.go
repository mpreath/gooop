package workflow

import "log"

type BaseProcess struct {
	next Process
}

func (p *BaseProcess) SetNext(next Process) {
	p.next = next
}

func (p *BaseProcess) GetNext() Process {
	return p.next
}

func (p *BaseProcess) Execute(person *Person) {
	log.Printf("############# Base processing #############")
	if person.Zip != "" {
		log.Printf("Zip code: %s", person.Zip)
	}
}
