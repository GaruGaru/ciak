package cmd

import (
	daemonService "github.com/GaruGaru/ciak/daemon"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Start the ciak daemon service",
	Run: func(cmd *cobra.Command, args []string) {
		daemon := daemonService.NewCiakDaemon(conf.DaemonConfig)
		daemon.Start()
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
