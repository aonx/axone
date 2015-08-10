package main

import (
	"fmt"
	"io"
	"net/http"

	log "github.com/elarasu/handy/logger"
	"github.com/elarasu/handy/version"
	"github.com/spf13/cobra"
	"github.com/surge/surgemq/service"
	"golang.org/x/net/websocket"
)

var _appVersion *version.Version = &version.Version{0, 0, 1}

func startMq() {
	port := "1883"
	log.Debug("starting broker ...", port)
	svr := &service.Server{
		KeepAlive:        300,           // seconds
		ConnectTimeout:   2,             // seconds
		SessionsProvider: "mem",         // keeps sessions in memory
		Authenticator:    "mockSuccess", // always succeed
		TopicsProvider:   "mem",         // keeps topic subscriptions in memory
	}

	// Listen and serve connections at localhost:1883
	err := svr.ListenAndServe("tcp://:" + port)
	fmt.Printf("%v", err)
}

func printBinary(s []byte) {
	fmt.Printf("print b:")
	for n := 0; n < len(s); n++ {
		fmt.Printf("%d,", s[n])
	}
	fmt.Printf("\n")
}

// Echo the data received on the WebSocket.
func echoServer(ws *websocket.Conn) {
	log.Debug("copying data...")
	io.Copy(ws, ws)
}

func startWsServer(httpPort string, path string) {
	log.Debug("starting websocket server on:", path, " port:", httpPort)
	http.Handle(path, websocket.Handler(echoServer))
	err := http.ListenAndServe(httpPort, nil)
	if err != nil {
		panic("WebSocket ListenAndServe: " + err.Error())
	}
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
			// start surge
			startMq()
			//startWsServer(":5080", "/mqtt")
		},
	}
	Cmd.AddCommand(startCmd)
	Cmd.Execute()
}
