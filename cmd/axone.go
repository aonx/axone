package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/aonx/momonga/server"
	log "github.com/elarasu/basis/logger"
	"github.com/elarasu/basis/version"
	"github.com/spf13/cobra"
)

var _appVersion *version.Version = &version.Version{0, 0, 2}

func startMq(configFile string) {
	pid := os.Getpid()
	log.Info("Server pid: ", pid)

	confpath, _ := filepath.Abs(configFile)
	app := server.NewApplication(confpath)
	app.Start()
	app.Loop()
}

func printBinary(s []byte) {
	fmt.Printf("print b:")
	for n := 0; n < len(s); n++ {
		fmt.Printf("%d,", s[n])
	}
	fmt.Printf("\n")
}

func main() {
	var configFile string
	runtime.GOMAXPROCS(runtime.NumCPU())
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
			startMq(configFile)
		},
	}
	startCmd.Flags().StringVarP(&configFile, "config", "c", "", "config file")
	Cmd.AddCommand(startCmd)
	Cmd.Execute()
}
