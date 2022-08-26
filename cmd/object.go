package cmd

import (
	"gooop/object"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(objectCmd)
}

var objectCmd = &cobra.Command{
	Use:   "object",
	Short: "Demonstrates using objects in Go",
	Long:  `Demonstrates using objects in Go`,
	Run: func(cmd *cobra.Command, args []string) {

		//create new object, no initialized values
		o1 := object.Object{}

		// create a new object, initialize values
		o2 := object.Object{
			Field1: "value1", // notice we can only access the "public" methods or variables
			//field2: "value2", // field2 is not accessible outside of the object package
		}

		// set values for object 1
		o1.Field1 = "value1"   // Field1 is directly accessible
		o1.SetField2("value2") // field2 is behind acessor methods

		// set values for object 2
		o2.SetField2("value2")

		log.Println("### Object1 (o1) ###")
		log.Printf("o1.Field1: %s", o1.Field1)
		log.Printf("o1.field2: %s", o1.GetField2())

		log.Println("### Object2 (o2) ###")
		log.Printf("o2.Field1: %s", o2.Field1)
		log.Printf("o2.field2: %s", o2.GetField2())

	},
}
