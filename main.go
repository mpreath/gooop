package main

import (
	"fmt"
	handler "gooop/chain"
	"gooop/event"
	"os"
)

func main() {

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

	err_string := "creating a new handler"

	domain_handler, err := handler.New("domain")
	if err != nil {
		fmt.Printf("%s: %v", err_string, err)
		os.Exit(1)
	}
	generic_handler, err := handler.New("generic")
	if err != nil {
		fmt.Printf("%s: %v", err_string, err)
		os.Exit(1)
	}
	log_handler, err := handler.New("log")
	if err != nil {
		fmt.Printf("%s: %v", err_string, err)
		os.Exit(1)
	}

	err_string = "creating a new event"

	e, err := event.New("user:updated", "domain")
	if err != nil {
		fmt.Printf("%s: %v", err_string, err)
		os.Exit(1)
	}
	// _, err = event.New("invalid:event", "invalid")
	// if err != nil {
	// 	fmt.Printf("%s: %v", err_string, err)
	// 	os.Exit(1)
	// }

	domain_handler.SetNext(log_handler)
	log_handler.SetNext(generic_handler)
	domain_handler.Execute(e)

}
