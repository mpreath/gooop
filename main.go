package main

import (
	"fmt"
	"gooop/event"
	"gooop/router"
	"log"
)

func main() {

	// Example of using an EventFactory to create a domain event
	// and a generic event. Then routing each event accordingly.
	er := router.NewEventRouter()

	for i := 0; i < 5; i++ {
		var event_type string
		var event_name string
		if i%2 == 0 {
			event_type = "domain"
			event_name = "user:updated:" + fmt.Sprint(i)
		} else {
			event_type = fmt.Sprint(i)
			event_name = "button:clicked:" + fmt.Sprint(i)
		}
		e := event.GenerateEvent(event_name, event_type)
		log.Printf("Generated Event: %s [%T]\n", e.GetName(), event_type)
		er.RouteEvent(e)
	}

	// Wait for concurrect routing to finish
	er.Wait()
}
