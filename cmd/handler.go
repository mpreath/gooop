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

		var eventQueue []handler.Event

		eventQueue = append(eventQueue, handler.NewEvent("userCreated", map[string]string{
			"object":   "user",
			"objectId": "1",
			"action":   "created",
		}))

		eventQueue = append(eventQueue, handler.NewEvent("propertyCreated", map[string]string{
			"object":   "property",
			"objectId": "1",
			"action":   "created",
		}))

		eventQueue = append(eventQueue, handler.NewEvent("sessionDeleted", map[string]string{
			"object":   "session",
			"objectId": "1",
			"action":   "deleted",
		}))

		for _, event := range eventQueue {
			handler.HandleEvent(event)
		}

	},
}
