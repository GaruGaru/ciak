package config

type CiakConfig struct {
	MediaPath    string
	ServerConfig CiakServerConfig
	DaemonConfig CiakDaemonConfig
}

type CiakServerConfig struct {
	ServerBinding string
}

type CiakDaemonConfig struct {
	OutputPath       string
	AutoConvertMedia bool
	DeleteOriginal   bool
	Workers          int
	QueueSize        int
}
