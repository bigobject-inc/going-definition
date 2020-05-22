package worker

import (
	"bitbucket.org/bigobject/going-definitions/logger"
	"github.com/kataras/golog"
)

// Worker Going worker interface
type Worker interface {
	GetLogger() *logger.Logger
	Init(wc Core, params string) error
	Defer()
	BeforeProcess() error
	Process(data interface{}) (interface{}, error)
	Timeout(seconds int, data interface{}) error
}

// Definition Going worker definition
type Definition struct {
	Key          string
	Name         string
	Desc         string
	Version      string
	Creator      Creator
	ParamsSchema string
	Path         string
}

// DefinitionPlugin Going worker definition plugin
type DefinitionPlugin struct {
	Name         string
	Desc         string
	Version      string
	ParamsSchema string
}

// DefinitionSetting Going worker definition setting
type DefinitionSetting struct {
	Key          string `json:"key"`
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	Version      string `json:"version"`
	Path         string `json:"path"`
	ParamsSchema string `json:"paramsSchema"`
}

// Timeout Going worker timeout
type Timeout struct {
	Seconds int
	Data    interface{}
}

// Creator Going worker creator
type Creator func(logger *golog.Logger) Worker

// CreatorsMap Going worker creator map
type CreatorsMap map[string]Creator
