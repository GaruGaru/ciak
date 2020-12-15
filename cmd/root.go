package cmd

import (
	"fmt"
	"github.com/GaruGaru/ciak/pkg/cache"
	"github.com/GaruGaru/ciak/pkg/config"
	"github.com/GaruGaru/ciak/pkg/media/details"
	"github.com/GaruGaru/ciak/pkg/media/discovery"
	"github.com/GaruGaru/ciak/pkg/server"
	"github.com/GaruGaru/ciak/pkg/server/auth"
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

		authenticator := auth.NewEnvAuthenticator()

		detailsRetrievers := make([]details.Retriever, 0)
		if conf.ServerConfig.OmdbApiKey != "" {
			detailsRetrievers = append(detailsRetrievers, details.Omdb(conf.ServerConfig.OmdbApiKey))
		}

		detailsController := details.NewController(cache.Memory(), detailsRetrievers...)

		server := server.NewCiakServer(conf.ServerConfig, mediaDiscovery, authenticator, detailsController)

		err := server.Run()

		if err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&conf.MediaPath, "media", "/data", "Path containing media files")
	rootCmd.PersistentFlags().StringVar(&conf.ServerConfig.ServerBinding, "bind", "0.0.0.0:8082", "interface and port binding for the web server")
	rootCmd.PersistentFlags().BoolVar(&conf.ServerConfig.AuthenticationEnabled, "auth", false, "if active enable user authentication for the web server")
	rootCmd.PersistentFlags().StringVar(&conf.DaemonConfig.Database, "db", "ciak_daemon.db", "database file used for persistence")
	rootCmd.PersistentFlags().StringVar(&conf.ServerConfig.OmdbApiKey, "omdb-api-key", "", "omdb movie metadata api key")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
