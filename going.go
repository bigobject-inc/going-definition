package going

import (
	"os"

	"bitbucket.org/bigobject/going-definitions/flow"
	"bitbucket.org/bigobject/going-definitions/logger"
	"bitbucket.org/bigobject/going-definitions/sourceworker"
	"bitbucket.org/bigobject/going-definitions/worker"
)

// PluginCatalogSourceWorker Plugin catalog source worker
const PluginCatalogSourceWorker = "sourceWorkers"

// PluginCatalogWorker Plugin catalog worker
const PluginCatalogWorker = "workers"

// PluginCatalogTmp Plugin catalog tmp
const PluginCatalogTmp = "tmp"

// Configure going configure
type Configure struct {
	// Web           web.Configure `json:"web"`
	LogLevel      string `json:"logLevel"`
	ChannelBuffer int    `json:"channelBuffer"`
}

// FlowConfigure going flow configure
type FlowConfigure struct {
	Flows         []flow.Setting            `json:"flows"`
	SourceWorkers []sourceworker.Definition `json:"sourceWorkers"`
	Workers       []worker.Definition       `json:"workers"`
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
	ID      string `json:"id"`
	Code    string `json:"code"`
	Catalog string `json:"catalog"`
	Path    string `json:"path"`
}
