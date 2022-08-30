package cmd

import (
	"gooop/handler"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(handlerCmd)
}

var handlerCmd = &cobra.Command{
	Use:   "handler",
	Short: "Demonstrates using event handlers in Go",
	Long:  `Demonstrates using event handlers in Go`,
	Run: func(cmd *cobra.Command, args []string) {

		// empty queue of events
		var eventQueue []handler.Event

		// create a user event
		userEvent := handler.UserEvent{}
		userEvent.SetName("userCreated")
		userEvent.SetAttributes(map[string]string{
			"object":   "user",
			"objectId": "1",
			"action":   "created",
		})

		// add the event to the queue
		eventQueue = append(eventQueue, &userEvent)

		propertyEvent := handler.PropertyEvent{}
		propertyEvent.SetName("propertyCreated")
		propertyEvent.SetAttributes(map[string]string{
			"object":   "property",
			"objectId": "1",
			"action":   "created",
		})

		eventQueue = append(eventQueue, &propertyEvent)

		otherEvent := handler.OtherEvent{}
		otherEvent.SetName("sessionDeleted")
		otherEvent.SetAttributes(map[string]string{
			"object":   "session",
			"objectId": "1",
			"action":   "deleted",
		})

		eventQueue = append(eventQueue, &otherEvent)

		// loop through the event queue and pass each event to the queue
		for _, event := range eventQueue {
			handler.HandleEvent(event)
		}

	},
}
