package cmd

import (
	"github.com/GaruGaru/ciak/discovery"
	server2 "github.com/GaruGaru/ciak/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the ciak web server",
	Run: func(cmd *cobra.Command, args []string) {

		mediaDiscovery := discovery.FileSystemMediaDiscovery{
			BasePath: conf.MediaPath,
		}

		server := server2.CiakServer{
			MediaDiscovery: mediaDiscovery,
			Config:         conf.ServerConfig,
		}

		server.Run()

	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
