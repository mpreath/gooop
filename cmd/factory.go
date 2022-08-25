package cmd

import (
	"fmt"
	"gooop/event"
	"gooop/router"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(factoryCmd)
}

var factoryCmd = &cobra.Command{
	Use:   "factory",
	Short: "Demonstrates using the factory design pattern",
	Long:  `Demonstrates using the factory design pattern`,
	Run: func(cmd *cobra.Command, args []string) {
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
			e, err := event.New(event_name, event_type)
			if err != nil {
				log.Println(err.Error())
			} else {
				log.Printf("Generated Event: %s [%T]\n", e.GetName(), e)
				er.RouteEvent(e)
			}
		}

		// Wait for concurrect routing to finish
		er.Wait()
	},
}
