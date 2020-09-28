package flow

import (
	"context"

	"github.com/bigobject-inc/going-definition/v2"
	"github.com/bigobject-inc/going-definition/v2/logger"
)

// SourceWorker Going Source worker interface
type SourceWorker interface {
	GetInfo() going.PluginInfo
	GetLogger() logger.Logger
	Init(swc SourceWorkerCore, params string) error
	BeforeStart() error
	Start(ctx context.Context) error
	Stop() error
}
