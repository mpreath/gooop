package main

import (
	"fmt"
	"gooop/event"
	"gooop/router"
)

func main() {

	// Example of using an EventFactory to create a domain event
	// and a generic event. Then routing each event accordingly.

	fmt.Println("------------------------------------------")

	e1 := event.GenerateEvent("domain:event", "domain")
	fmt.Printf("Generated Event: %s [%T]\n", e1.GetName(), e1)
	router.RouteEvent(e1)

	fmt.Println("------------------------------------------")

	e2 := event.GenerateEvent("generic:event", "other")
	fmt.Printf("Generated Event: %s [%T]\n", e2.GetName(), e2)
	router.RouteEvent(e2)
}
