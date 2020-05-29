package flow

import (
	"context"

	"bitbucket.org/bigobject/going-definitions"
	"bitbucket.org/bigobject/going-definitions/logger"
)

// Flow Going Flow interface
type Flow interface {
	GetID() string
	GetLogger() logger.Logger
	GetSetting() Setting
	GetSourceWorkerCore() SourceWorkerCore
	GetStatus() string
	GetWorkerCoreByID(id string) (WorkerCore, error)
	GetWorkerCores() []WorkerCore
	Init(setting Setting, logPath, logLevel string) error
	Defer()
	Start(ctx context.Context) error
	Stop() error
}

// SourceWorker Going Source worker interface
type SourceWorker interface {
	GetInfo() going.PluginInfo
	GetLogger() logger.Logger
	Init(swc SourceWorkerCore, params string) error
	BeforeStart() error
	Start(ctx context.Context) error
	Stop() error
}

// SourceWorkerCore Going source Worker core interface
type SourceWorkerCore interface {
	GetChannelOut() []*Channel
	GetID() string
	GetSetting() SettingSourceWorker
	GetSourceWorker() SourceWorker
	GetStatus() string
	SendDataNext(data interface{}) error
	SetChannelOut(chanOut *Channel) error
	SetSourceWorker(s SourceWorker) error
	Start(ctx context.Context) error
	Stop() error
}

// Worker Going worker interface
type Worker interface {
	GetInfo() going.PluginInfo
	GetLogger() logger.Logger
	Init(wc WorkerCore, params string) error
	Defer()
	BeforeProcess() error
	Process(data interface{}) (interface{}, error)
	Timeout(seconds int, data interface{}) error
}

// WorkerCore Going Worker core interface
type WorkerCore interface {
	GetChannelIn() *Channel
	GetChannelOut() []*Channel
	GetID() string
	GetSetting() SettingWorker
	GetStatus() string
	GetWorker() Worker
	SendDataNext(data interface{}) error
	SetChannelIn(chanIn *Channel) error
	SetChannelOut(chanOut *Channel) error
	SetWorker(w Worker) error
	SetTimeout(seconds int, data interface{})
	Start(ctx context.Context) error
	Stop() error
}

// Channel flow worker channel
type Channel struct {
	ID string
	C  chan interface{}
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
	ID      string `json:"id"`
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

// SourceWorkerCreator Going flow source worker creator
type SourceWorkerCreator func(logger logger.Logger) SourceWorker

// WorkerCreator Going flow worker creator
type WorkerCreator func(logger logger.Logger) Worker
