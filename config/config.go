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
	AutoConvertMedia bool
	DeleteOriginal   bool
	Workers          int
	QueueSize        int
}
