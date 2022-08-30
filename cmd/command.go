package cmd

import (
	"gooop/command"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(commandCmd)
}

var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "Demonstrates using commands in Go",
	Long:  `Demonstrates using commands in Go`,
	Run: func(cmd *cobra.Command, args []string) {

		userCmd := command.CreateUserCommand{}
		userCmd.Object = command.User{
			Name: "Matt Reath",
		}
		userCmd.Execute()
		if userCmd.Error != nil {
			log.Fatal(userCmd.Error.Error())
		}
		log.Printf("Command %T Executed on %T", userCmd, userCmd.Object)
	},
}
