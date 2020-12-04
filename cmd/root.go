package cmd

import (
	"fmt"
	"github.com/GaruGaru/ciak/internal/cache"
	"github.com/GaruGaru/ciak/internal/config"
	"github.com/GaruGaru/ciak/internal/daemon"
	"github.com/GaruGaru/ciak/internal/media/details"
	"github.com/GaruGaru/ciak/internal/media/discovery"
	"github.com/GaruGaru/ciak/internal/media/encoding"
	"github.com/GaruGaru/ciak/internal/server"
	"github.com/GaruGaru/ciak/internal/server/auth"
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

		mediaDiscovery := discovery.NewFileSystemDiscovery(conf.MediaPath)

		encoder := encoding.FFMpeg()

		authenticator := auth.NewEnvAuthenticator()

		detailsRetrievers := make([]details.Retriever, 0)
		if conf.ServerConfig.OmdbApiKey != "" {
			detailsRetrievers = append(detailsRetrievers, details.Omdb(conf.ServerConfig.OmdbApiKey))
		}

		detailsController := details.NewController(cache.Memory(), detailsRetrievers...)

		daemon, err := daemon.New(conf.DaemonConfig, mediaDiscovery, encoder)

		if err != nil {
			panic(err)
		}

		server := server.NewCiakServer(conf.ServerConfig, mediaDiscovery, authenticator, daemon, detailsController)

		err = daemon.Start()

		if err != nil {
			panic(err)
		}

		err = server.Run()

		if err != nil {
			panic(err)
		}

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
	rootCmd.PersistentFlags().StringVar(&conf.DaemonConfig.Database, "db", "ciak_daemon.db", "database file used for persistence")
	rootCmd.PersistentFlags().StringVar(&conf.DaemonConfig.TransferDestination, "transfer-path", "", "path where to transfer media")
	rootCmd.PersistentFlags().StringVar(&conf.ServerConfig.OmdbApiKey, "omdb-api-key", "", "omdb movie metadata api key")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
