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

		eventQueue := make([]handler.Event, 3)

		userEvent := handler.UserEvent{
			BaseEvent: handler.BaseEvent{
				Name: "userUpdated",
				Attributes: map[string]string{
					"object":   "user",
					"objectId": "1",
					"action":   "updated",
				},
			}}

		eventQueue[0] = &userEvent

		propertyEvent := handler.PropertyEvent{
			BaseEvent: handler.BaseEvent{
				Name: "propertyCreated",
				Attributes: map[string]string{
					"object":   "property",
					"objectId": "1",
					"action":   "created",
				},
			}}

		eventQueue[1] = &propertyEvent

		otherEvent := handler.OtherEvent{
			BaseEvent: handler.BaseEvent{
				Name: "sessionDeleted",
				Attributes: map[string]string{
					"object":   "session",
					"objectId": "1",
					"action":   "deleted",
				},
			}}

		eventQueue[2] = &otherEvent

		for _, event := range eventQueue {
			handler.HandleEvent(event)
		}

	},
}
