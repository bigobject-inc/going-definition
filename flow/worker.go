package flow

import (
	"github.com/bigobject-inc/going-definition/v2"
	"github.com/bigobject-inc/going-definition/v2/logger"
)

// Worker Going worker interface
type Worker interface {
	BeforeStart() error
	BeforeStop() error
	Defer()
	GetInfo() going.PluginInfo
	GetLogger() logger.Logger
	Init(wc WorkerCore, params string) error
	Notification(data interface{}) error
	Process(data interface{}) (interface{}, error)
	Timeout(seconds int, data interface{}) error
}
