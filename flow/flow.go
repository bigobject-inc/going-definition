package flow

import (
	"context"

	"bitbucket.org/bigobject/going-definitions/logger"
)

// Flow Going Flow interface
type Flow interface {
	GetID() string
	GetStatus() string
	GetSetting() Setting
	GetLogger() *logger.Logger
	Start(ctx context.Context) error
	Stop() error
}

// WorkerCore Going Worker core interface
type WorkerCore interface {
	GetID() string
	GetStatus() string
	// GetSetting() Setting
	GetChannelIn() *Channel
	GetChannelOut() []*Channel
	GetFromWorkerID() []string
	GetToWorkerID() []string
	GetWorker() Worker
	GetSharedMemoryValue(sID, key string) (string, error)
	SetWorker(worker Worker)
	SetChannelIn(channel *Channel)
	SetChannelOut(channel *Channel)
	SetFromWorkerID(id string)
	SetToWorkerID(id string)
	SetTimeout(seconds int, data interface{})
	SetSharedMemory(creatorName string, dataSourceName string, params interface{}) (string, error)
	SetSharedMemoryValue(sID, key, value string, expired int) error
	SendDataNext(data interface{}) error
	Start(ctx context.Context) error
	Stop() error
}

// Worker Going worker interface
type Worker interface {
	GetLogger() *logger.Logger
	Init(wc WorkerCore, params string) error
	Defer()
	BeforeProcess() error
	Process(data interface{}) (interface{}, error)
	Timeout(seconds int, data interface{}) error
}

// SourceWorker Going Source worker interface
type SourceWorker interface {
	GetID() string
	GetStatus() string
	// GetSetting() Setting
	GetLogger() *logger.Logger
	Init(params string) error
	BeforeRun() error
	Start(ctx context.Context) error
	Stop() error
}

// Setting Going flow setting
type Setting struct {
	ID           string              `json:"id"`
	Name         string              `json:"name"`
	Desc         string              `json:"desc"`
	SourceWorker SettingSourceWorker `json:"sourceWorker"`
	Workers      []SettingWorker     `json:"workers"`
}

// SettingSourceWorker Going flow setting source worker
type SettingSourceWorker struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Version string `json:"version"`
	Params  string `json:"params"`
}

// SettingWorker Going flow setting worker
type SettingWorker struct {
	ID      string   `json:"id"`
	Parents []string `json:"parents"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Desc    string   `json:"desc"`
	Version string   `json:"version"`
	Params  string   `json:"params"`
}

// Channel Going worker channel
type Channel struct {
	ID      string
	From    string
	To      string
	Channel chan interface{}
}
