package flow

import (
	"context"

	"github.com/bigobject-inc/going-definition/v2/notification"
	"github.com/bigobject-inc/going-definition/v2/sharedmemory"
)

// SourceWorkerCore Going source Worker core interface
type SourceWorkerCore interface {
	GetChannelIn() *Channel
	GetChannelOut() []*Channel
	GetFlow() Flow
	GetID() string
	GetNotification(nID string) (notification.Notification, error)
	GetSetting() SettingSourceWorker
	GetSharedMemory(sID string) (sharedmemory.SharedMemory, error)
	GetSourceWorker() SourceWorker
	GetStatus() string
	SendDataNext(data interface{}) error
	SetChannelIn(chanIn *Channel) error
	SetChannelOut(chanOut *Channel) error
	SetNotification(creatorName, dataSourceName string, params interface{}) (string, error)
	SetSharedMemory(creatorName, dataSourceName string, params interface{}) (string, error)
	Init(sw SourceWorker, f Flow) error
	Defer()
	Start(ctx context.Context) error
	Stop() error
}
