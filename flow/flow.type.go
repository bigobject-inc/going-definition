package flow

import "github.com/bigobject-inc/golib/logger"

// Channel flow worker channel
type Channel struct {
	ID string
	C  chan interface{}
}

// NodePosition Going flow node position
type NodePosition struct {
	Px float64 `json:"px"`
	Py float64 `json:"py"`
}

// Setting Going flow setting
type Setting struct {
	ID            string              `json:"id"`
	Name          string              `json:"name"`
	Desc          string              `json:"desc"`
	ChannelBuffer int                 `json:"channelBuffer"`
	SourceWorker  SettingSourceWorker `json:"sourceWorker"`
	Workers       []SettingWorker     `json:"workers"`
}

// SettingSourceWorker Going flow setting source worker
type SettingSourceWorker struct {
	ID       string       `json:"id"`
	Key      string       `json:"key"`
	Name     string       `json:"name"`
	Desc     string       `json:"desc"`
	Version  string       `json:"version"`
	Params   string       `json:"params"`
	Position NodePosition `json:"position"`
}

// SettingWorker Going flow setting worker
type SettingWorker struct {
	ID       string       `json:"id"`
	Parents  []string     `json:"parents"`
	Key      string       `json:"key"`
	Name     string       `json:"name"`
	Desc     string       `json:"desc"`
	Version  string       `json:"version"`
	Params   string       `json:"params"`
	Position NodePosition `json:"position"`
}

// SourceWorkerCreator Going flow source worker creator
type SourceWorkerCreator func(logger logger.Logger) SourceWorker

// WorkerCoreFilter worker core filter
type WorkerCoreFilter func(wc WorkerCore) bool

// WorkerCreator Going flow worker creator
type WorkerCreator func(logger logger.Logger) Worker
