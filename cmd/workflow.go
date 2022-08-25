package cmd

import (
	"gooop/workflow"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(workflowCmd)
}

var workflowCmd = &cobra.Command{
	Use:   "workflow",
	Short: "Demonstrates using a workflow design pattern",
	Long:  `Demonstrates using a workflow design pattern`,
	Run: func(cmd *cobra.Command, args []string) {
		// create a new workflow
		wf := workflow.New()

		// create two processes
		p1 := workflow.ValidateProcess{}
		p2 := workflow.BaseProcess{}
		p1.SetNext(nil)
		p2.SetNext(nil)

		// add processes to workflow
		// will execute in the order in which they are added
		wf.Add(&p1)
		wf.Add(&p2)

		person1 := workflow.Person{
			Name:    "Matt Reath",
			Age:     42,
			Address: "N1594 Forest Dr.",
			City:    "Norway",
			State:   "MI",
			Zip:     "49870",
		}
		wf.Run(&person1)

		person2 := workflow.Person{
			Name:    "Test User",
			Age:     18,
			Address: "",
			City:    "Norway",
			State:   "MI",
			Zip:     "49870",
		}
		wf.Run(&person2)
	},
}
