package cmd

import (
	"fmt"
	"github.com/GaruGaru/ciak/config"
	daemonService "github.com/GaruGaru/ciak/daemon"
	"github.com/GaruGaru/ciak/discovery"
	webServer "github.com/GaruGaru/ciak/server"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ciak",
	Short: "Ciak is a lightweight media server.",
	Run: func(cmd *cobra.Command, args []string) {
		daemon := daemonService.NewCiakDaemon(conf.DaemonConfig)
		go daemon.Start()

		mediaDiscovery := discovery.FileSystemMediaDiscovery{
			BasePath: conf.MediaPath,
		}

		server := webServer.CiakServer{
			MediaDiscovery: mediaDiscovery,
			Config:         conf.ServerConfig,
		}

		server.Run()
	},
}

var conf = config.CiakConfig{
	DaemonConfig: config.CiakDaemonConfig{},
	ServerConfig: config.CiakServerConfig{},
}

func init() {
	pFlags := rootCmd.PersistentFlags()
	pFlags.StringVarP(&conf.MediaPath, "media", "m", "/data", "Path containing media files")
	pFlags.StringVarP(&conf.ServerConfig.ServerBinding, "bind", "b", "0.0.0.0:8082", "interface and port binding for the web server")
	pFlags.BoolVarP(&conf.DaemonConfig.AutoConvertMedia, "auto-convert-media", "a", false, "if active the server automatically converts media files to a streamable format")
	pFlags.BoolVarP(&conf.DaemonConfig.DeleteOriginal, "delete-original-media", "d", false, "if active delete the original media after conversion")
	pFlags.IntVarP(&conf.DaemonConfig.QueueSize, "queue-size", "q", 1000, "daemon tasks queue size")
	pFlags.IntVarP(&conf.DaemonConfig.Workers, "workers", "c", 2, "daemon number of workers")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
