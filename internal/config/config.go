package config

type CiakConfig struct {
	MediaPath    string
	ServerConfig CiakServerConfig
	DaemonConfig CiakDaemonConfig
}

type CiakServerConfig struct {
	ServerBinding         string
	AuthenticationEnabled bool
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
