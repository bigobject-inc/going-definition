package going

import (
	"os"

	"bitbucket.org/bigobject/going-definitions/logger"
)

const (
	// PluginSourceWorker Plugin source worker
	PluginSourceWorker = "sourceWorkers"
	// PluginWorker Plugin worker
	PluginWorker = "workers"
)

const (
	// StatusCreated Status created
	StatusCreated = "created"
	// StatusReady Status ready
	StatusReady = "ready"
	// StatusRunning Status running
	StatusRunning = "running"
	// StatusStop Status stop
	StatusStop = "stop"
	// StatusReceived Status received
	StatusReceived = "received"
	// StatusProcessed Status processed
	StatusProcessed = "processed"
)

// Configure going configure
type Configure struct {
	// Web           web.Configure `json:"web"`
	LogLevel string `json:"logLevel"`
}

// Logger going logger
type Logger struct {
	Logger  logger.Logger
	LoggerF *os.File
}

// Paths going paths
type Paths struct {
	Root     string
	Logs     string
	Settings string
	Plugin   string
}

// Plugin going plugin
type Plugin struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Path string `json:"path"`
}

// PluginInfo going plugin info
type PluginInfo struct {
	Name         string
	Desc         string
	Version      string
	ParamsSchema string
}
