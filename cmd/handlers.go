package cmd

import (
	"fmt"
	"gooop/event"
	"gooop/handler"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(handlersCmd)
}

var handlersCmd = &cobra.Command{
	Use:   "handlers",
	Short: "Demonstrates using the event handler design pattern",
	Long:  `Demonstrates using the event handler design pattern`,
	Run: func(cmd *cobra.Command, args []string) {
		// Event handling through chain of command

		err_string := "creating a new handler"

		domain_handler, err := handler.New("domain")
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
		domain_handler.Execute(e)
	},
}
