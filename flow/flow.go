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
