package cmd

import (
	"fmt"
	"github.com/GaruGaru/ciak/config"
	"github.com/GaruGaru/ciak/core"
	"github.com/GaruGaru/ciak/daemon"
	"github.com/GaruGaru/ciak/discovery"
	"github.com/GaruGaru/ciak/encoding"
	"github.com/GaruGaru/ciak/server"
	"github.com/spf13/cobra"
	"os"
)

var (
	conf config.CiakConfig
)

var rootCmd = &cobra.Command{
	Use:   "ciak",
	Short: "Ciak is a lightweight media server.",
	Run: func(cmd *cobra.Command, args []string) {

		conf.DaemonConfig.OutputPath = conf.MediaPath

		mediaDiscovery := discovery.FileSystemMediaDiscovery{
			BasePath: conf.MediaPath,
		}

		encoder := encoding.FFMpegEncoder{}

		ciak := core.Ciak{
			Config: conf,
			Daemon: daemon.NewCiakDaemon(conf.DaemonConfig, mediaDiscovery, encoder),
			Server: server.NewCiakServer(conf.ServerConfig, mediaDiscovery),
		}

		go ciak.Daemon.Start()
		ciak.Server.Run()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&conf.MediaPath, "media", "/data", "Path containing media files")
	rootCmd.PersistentFlags().StringVar(&conf.ServerConfig.ServerBinding, "bind", "0.0.0.0:8082", "interface and port binding for the web server")
	rootCmd.PersistentFlags().BoolVar(&conf.DaemonConfig.AutoConvertMedia, "auto-convert-media", false, "if active the server automatically converts media files to a streamable format")
	rootCmd.PersistentFlags().BoolVar(&conf.DaemonConfig.DeleteOriginal, "delete-original-media", false, "if active delete the original media after conversion")
	rootCmd.PersistentFlags().IntVar(&conf.DaemonConfig.QueueSize, "queue-size", 1000, "daemon tasks queue size")
	rootCmd.PersistentFlags().IntVar(&conf.DaemonConfig.Workers, "workers", 2, "daemon number of workers")
	rootCmd.PersistentFlags().BoolVar(&conf.ServerConfig.AuthenticationEnabled, "auth", false, "if active enable user authentication for the web server")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
