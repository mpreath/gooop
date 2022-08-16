package main

import (
	"gooop/event"
	"gooop/router"
	"log"
)

func main() {

	// Example of using an EventFactory to create a domain event
	// and a generic event. Then routing each event accordingly.
	er := router.NewEventRouter()

	e1 := event.GenerateEvent("user:updated", "domain")
	log.Printf("Generated Event: %s [%T]\n", e1.GetName(), e1)
	er.RouteEvent(e1)

	e2 := event.GenerateEvent("button:clicked", "other")
	log.Printf("Generated Event: %s [%T]\n", e2.GetName(), e2)
	er.RouteEvent(e2)

	// Wait for concurrect routing to finish
	er.Wait()
}
