package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var Cmd = &cobra.Command{
		Use:   "axonectl",
		Short: "axone command line",
		Long:  `axone messaging`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of axone",
		Long:  `All software has versions. This is axone's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("axone messaging v0.0.1")
		},
	}
	Cmd.AddCommand(versionCmd)
	Cmd.Execute()
}
