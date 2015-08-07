package main

import (
	"fmt"

	"github.com/elarasu/handy/version"
	"github.com/spf13/cobra"
	"github.com/surge/surgemq/service"
)

var _appVersion *version.Version = &version.Version{0, 0, 1}

func startMq() {
	svr := &service.Server{
		KeepAlive:        300,           // seconds
		ConnectTimeout:   2,             // seconds
		SessionsProvider: "mem",         // keeps sessions in memory
		Authenticator:    "mockSuccess", // always succeed
		TopicsProvider:   "mem",         // keeps topic subscriptions in memory
	}

	// Listen and serve connections at localhost:1883
	err := svr.ListenAndServe("tcp://:1883")
	fmt.Printf("%v", err)
}

func main() {
	var Cmd = &cobra.Command{
		Use:   "axone",
		Short: "axone command line",
		Long:  `axone messaging`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "show version of oxone",
		Long:  `All software has versions. This is axone's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(_appVersion.AppString("axone"))
		},
	}
	Cmd.AddCommand(versionCmd)
	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "start the broker",
		Long:  `Start the message broker from command line`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("starting broker ...")
			// start surge
			startMq()
		},
	}
	Cmd.AddCommand(startCmd)
	Cmd.Execute()
}
