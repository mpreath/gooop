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

		e := handler.UserEvent{
			BaseEvent: handler.BaseEvent{
				Name: "userUpdated",
				Attributes: map[string]string{
					"object":   "user",
					"objectId": "1",
					"action":   "updated",
				},
			}}

		handler.HandleEvent(&e)

	},
}
