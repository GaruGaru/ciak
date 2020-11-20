package config

type CiakConfig struct {
	MediaPath    string
	ServerConfig CiakServerConfig
	DaemonConfig CiakDaemonConfig
}

type CiakServerConfig struct {
	ServerBinding         string `json:"bind"`
	AuthenticationEnabled bool   `json:"enable_auth"`
	OmdbApiKey            string `json:"omdb_api_key"`
}

type CiakDaemonConfig struct {
	OutputPath          string
	AutoConvertMedia    bool
	DeleteOriginal      bool
	Workers             int
	QueueSize           int
	Database            string
	TransferDestination string
}
