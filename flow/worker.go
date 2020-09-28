package flow

import (
	"github.com/bigobject-inc/going-definition/v2"
	"github.com/bigobject-inc/going-definition/v2/logger"
)

// Worker Going worker interface
type Worker interface {
	GetInfo() going.PluginInfo
	GetLogger() logger.Logger
	Init(wc WorkerCore, params string) error
	Defer()
	BeforeProcess() error
	Process(data interface{}) (interface{}, error)
	Notification(data interface{}) error
	Timeout(seconds int, data interface{}) error
}
