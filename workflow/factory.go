package workflow

import "log"

type Workflow struct {
	process Process
}

type Process interface {
	Run(*Person)
	SetNext(Process)
	GetNext() Process
}

func New() Workflow {
	return Workflow{process: nil}
}

func (w *Workflow) Add(process Process) {
	var i Process
	i = w.process
	if i == nil {
		// initial case, w.process == nil
		w.process = process
	} else {
		// other cases, need to find the end i.next == nil
		for ; i.GetNext() != nil; i = w.process.GetNext() {
		}
		i.SetNext(process)
	}

}

func (w *Workflow) Run(person *Person) {
	for i := w.process; i != nil; i = i.GetNext() {
		if !person.Handled {
			i.Run(person)
		} else {
			log.Println("****** Already handled *******")
		}
	}
}
