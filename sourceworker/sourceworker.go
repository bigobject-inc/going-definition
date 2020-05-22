package sourceworker

import (
	"context"

	"bitbucket.org/bigobject/going-definitions/logger"
	"github.com/kataras/golog"
)

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

// Definition Going source worker definition
type Definition struct {
	Key          string
	Name         string
	Desc         string
	Version      string
	Creator      Creator
	ParamsSchema string
	Path         string
}

// DefinitionPlugin Going source worker definition plugin
type DefinitionPlugin struct {
	Name         string
	Desc         string
	Version      string
	ParamsSchema string
}

// DefinitionSetting Going source worker definition setting
type DefinitionSetting struct {
	Key          string `json:"key"`
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	Version      string `json:"version"`
	Path         string `json:"path"`
	ParamsSchema string `json:"paramsSchema"`
}

// Creator Going source worker creator
type Creator func(logger *golog.Logger) SourceWorker

// CreatorsMap Going source worker creator map
type CreatorsMap map[string]Creator
