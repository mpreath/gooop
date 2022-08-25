package main

import "gooop/workflow"

func main() {

	// create a new workflow
	wf := workflow.New()

	// create two processes
	p1 := workflow.ValidateProcess{}
	p2 := workflow.BaseProcess{}
	p1.SetNext(nil)
	p2.SetNext(nil)

	// add processes to workflow
	// will execute in the order in which they are added
	wf.Add(&p1)
	wf.Add(&p2)

	person := workflow.Person{
		Name:    "Matt Reath",
		Age:     18,
		Address: "N1594 Forest Dr.",
	}
	wf.Run(&person)

	// Example of using an EventFactory to create a domain event
	// and a generic event. Then routing each event accordingly.
	// er := router.NewEventRouter()

	// for i := 0; i < 5; i++ {
	// 	var event_type string
	// 	var event_name string
	// 	if i%2 == 0 {
	// 		event_type = "domain"
	// 		event_name = "user:updated:" + fmt.Sprint(i)
	// 	} else {
	// 		event_type = fmt.Sprint(i)
	// 		event_name = "button:clicked:" + fmt.Sprint(i)
	// 	}
	// 	e := event.GenerateEvent(event_name, event_type)
	// 	log.Printf("Generated Event: %s [%T]\n", e.GetName(), event_type)
	// 	er.RouteEvent(e)
	// }

	// // Wait for concurrect routing to finish
	// er.Wait()

	// Event handling through chain of command

	// err_string := "creating a new handler"

	// domain_handler, err := handler.New("domain")
	// if err != nil {
	// 	fmt.Printf("%s: %v", err_string, err)
	// 	os.Exit(1)
	// }
	// log_handler, err := handler.New("log")
	// if err != nil {
	// 	fmt.Printf("%s: %v", err_string, err)
	// 	os.Exit(1)
	// }

	// err_string = "creating a new event"

	// e, err := event.New("user:updated", "domain")
	// if err != nil {
	// 	fmt.Printf("%s: %v", err_string, err)
	// 	os.Exit(1)
	// }
	// // _, err = event.New("invalid:event", "invalid")
	// // if err != nil {
	// // 	fmt.Printf("%s: %v", err_string, err)
	// // 	os.Exit(1)
	// // }

	// domain_handler.SetNext(log_handler)
	// domain_handler.Execute(e)

}
